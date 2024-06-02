package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/parquet"
	"github.com/klimenkoOleg/migrate-stock-data/internal/adapter/postgres"
	"github.com/klimenkoOleg/migrate-stock-data/internal/app"
	"github.com/klimenkoOleg/migrate-stock-data/internal/common"
	"github.com/klimenkoOleg/migrate-stock-data/internal/core/usercases/fetchticks"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func main() {
	log := common.NewDefaultLogger()
	ctx, cancel := context.WithCancel(context.Background())
	config := app.MustLoadConfig(log, "conf/fetcher.yml")
	if config.DateStarted == nil || config.DateEnded == nil {
		log.Fatal("date.start and date.end should be specified in config")
	}

	app := app.New(config, log)
	app.Run(ctx, cancel, NewUploadWorker(config, log))
}

type UploadWorker struct {
	config *app.Config
	log    *zap.Logger
}

func NewUploadWorker(config *app.Config, log *zap.Logger) *UploadWorker {
	return &UploadWorker{config: config, log: log}
}

func (w *UploadWorker) Do(ctx context.Context, appCloser *common.Closer) error {

	pool, err := pgxpool.Connect(ctx, w.config.PostgresDSN)
	if err != nil {
		return fmt.Errorf("connect to Postgres, %w", err)
	}
	appCloser.AddCloseFunc(pool.Close)
	//defer func() {
	//	pool.Close()
	//}()
	//pgRepo := clickhouse.NewRepo(w.log, pool)
	pgRepo := postgres.NewRepo(pool, w.log)

	writer, err := parquet.NewTickParquetWriter(w.config.ParquetPath)
	if err != nil {
		//log.Error("open writer error", zap.Error(err))
		return fmt.Errorf("open writer error, %w", err)
	}
	appCloser.AddCloser(writer)
	//defer func() {
	//	errClose := writer.Close()
	//	log.Error("closing writer", zap.Error(errClose))
	//}()

	go func() {
		tickService := fetchticks.NewTickFetcherService(pgRepo, writer)
		err = tickService.Do(ctx, *w.config.DateStarted, *w.config.DateEnded)
		if err != nil {
			log.Error("importing service", zap.Error(err))
		}
		// Finish all the app work.
		appCloser.Close()
	}()

	return nil
}
