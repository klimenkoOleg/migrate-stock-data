package uploader

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/milvus"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/parquet"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/postgres"
	"github.com/klimenkoOleg/migrate-stock-data/internal/app"
	"github.com/klimenkoOleg/migrate-stock-data/internal/common"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/usercases/uploadticks"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"time"
)

type UploadWorker struct {
	config *app.Config
	log    *zap.Logger
}

func NewWorker(config *app.Config, log *zap.Logger) *UploadWorker {
	return &UploadWorker{config: config, log: log}
}

func (w *UploadWorker) Do(ctx context.Context, appCloser *common.Closer) error {
	w.config.TargetType = "milvus" // TODO remove this test.
	var repo uploadticks.Repo
	switch w.config.TargetType {
	case "postgres":
		pgRepo, err := w.getPgConnection(ctx)
		if err != nil {
			return fmt.Errorf("connect to Postgres, %w", err)
		}
		repo = pgRepo
	case "milvus":
		mRepo, err := w.getMilvusConnection(ctx)
		if err != nil {
			return fmt.Errorf("connect to Milvus, %w", err)
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

	go func() {
		var tickService Service = uploadticks.NewTickUploaderService(repo, reader, w.config.LimitBatchSize)
		if err := tickService.ImportFromStorageToDatabase(ctx); err != nil {
			log.Error("importing Service", zap.Error(err))
		}

		//if err := repo.Flush(ctx); err != nil {
		//	log.Error("repo flush failed", zap.Error(err))
		//}
		appCloser.Close()
	}()

	return nil
}

func (w *UploadWorker) getMilvusConnection(_ context.Context) (uploadticks.Repo, error) {
	client := milvus.NewMilvus(w.config, w.log)
	//client.Prepare(ctx)

	return client, nil
}

func (w *UploadWorker) getPgConnection(ctx context.Context) (uploadticks.Repo, error) {
	dbConfig, err := pgxpool.ParseConfig(w.config.PostgresDSN)
	dbConfig.ConnConfig.Logger = &myQueryTracer{log: w.log}
	pool, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to Postgres, %w", err)
	}
	pgRepo := postgres.New(pool, w.log)

	return pgRepo, nil
}

type myQueryTracer struct {
	log *zap.Logger
}

func (t *myQueryTracer) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	for k, v := range data {
		t.log.Debug("Executing command", zap.Any(k, v))
	}
}
