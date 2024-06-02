package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/milvius"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/parquet"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/postgres"
	"github.com/klimenkoOleg/migrate-stock-data/internal/app"
	"github.com/klimenkoOleg/migrate-stock-data/internal/common"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/usercases/uploadticks"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"time"
)

func main() {
	log := common.NewDefaultLogger()
	ctx, cancel := context.WithCancel(context.Background())

	config := app.MustLoadConfig(log, "conf/uploader.yml")

	app := app.New(config, log)

	worker := NewUploadWorker(config, log)
	app.Run(ctx, cancel, worker)
}

type Worker interface {
	Do(ctx context.Context, appCloser *common.Closer) error
}

type UploadWorker struct {
	config *app.Config
	log    *zap.Logger
}

func NewUploadWorker(config *app.Config, log *zap.Logger) *UploadWorker {
	return &UploadWorker{config: config, log: log}
}

type myQueryTracer struct {
	log *zap.Logger
}

func (t *myQueryTracer) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {

	for k, v := range data {
		t.log.Debug("Executing command", zap.Any(k, v))
	}

}

func (w *UploadWorker) getPgConnection(ctx context.Context) (uploadticks.Repo, error) {
	dbConfig, err := pgxpool.ParseConfig(w.config.PostgresDSN)
	dbConfig.ConnConfig.Logger = &myQueryTracer{log: w.log}
	pool, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to Postgres, %w", err)
	}

	//defer func() {
	//	pool.Close()
	//}()
	pgRepo := postgres.NewRepo(pool, w.log)

	return pgRepo, nil
}

func (w *UploadWorker) Do(ctx context.Context, appCloser *common.Closer) error { // SOLID broken

	//defer func() {
	//	if err != nil {
	//		appCloser.Close()
	//		a.log.Fatal("init error", zap.Error(err))
	//	}
	//}()
	var repo uploadticks.Repo

	w.config.TargetType = "milvius"

	switch w.config.TargetType {
	case "postgres":
		pgRepo, err := w.getPgConnection(ctx)
		if err != nil {
			return fmt.Errorf("connect to Postgres, %w", err)
		}
		repo = pgRepo
	case "milvius":
		mRepo, err := w.getMilviusConnection(ctx)
		if err != nil {
			return fmt.Errorf("connect to Milvius, %w", err)
		}
		repo = mRepo
	default:
		log.Fatal("Unknown repo target type")
	}
	appCloser.AddCloseFunc(repo.Close)

	endDate := w.config.DateStarted.Add(time.Hour * 24 * time.Duration(w.config.WindowDays))
	reader, err := parquet.NewTickParquetReader(
		w.config.ParquetPath,
		w.config.DateStarted,
		&endDate,
	)
	if err != nil {
		//log.Error("open reader error", zap.Error(err))
		return fmt.Errorf("open reader error, %w", err)
	}
	appCloser.AddCloser(reader)
	//defer func() {
	//	errClose := reader.Close()
	//	log.Error("closing reader", zap.Error(errClose))
	//}()

	go func() {
		tickService := uploadticks.NewTickDumperService(repo, reader, w.config.LimitBatchSize, w.config.LimitMaxRecords)

		if err := tickService.ImportFromStorageToDatabase(ctx); err != nil {
			log.Error("importing service", zap.Error(err))
		}

		if err := repo.Flush(ctx); err != nil {
			log.Error("repo flush failed", zap.Error(err))
		}
		appCloser.Close()
	}()

	return nil
}

func (w *UploadWorker) getMilviusConnection(ctx context.Context) (uploadticks.Repo, error) {
	client := milvius.NewMilvius(w.config, w.log) //, w.config.WindowDays)
	//client.Prepare(ctx)

	return client, nil
}
