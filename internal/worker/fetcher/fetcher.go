package fetcher

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

type UploadWorker struct {
	config *app.Config
	log    *zap.Logger
}

func NewWorker(config *app.Config, log *zap.Logger) *UploadWorker {
	return &UploadWorker{config: config, log: log}
}

func (w *UploadWorker) Do(ctx context.Context, appCloser *common.Closer) error {
	pool, err := pgxpool.Connect(ctx, w.config.PostgresDSN)
	if err != nil {
		return fmt.Errorf("connect to Postgres, %w", err)
	}
	appCloser.AddCloseFunc(func() error {
		pool.Close()
		return nil
	})
	pgRepo := postgres.New(pool, w.log)

	writer, err := parquet.NewTickParquetWriter(w.config.ParquetPath)
	if err != nil {
		//log.Error("open writer error", zap.Error(err))
		return fmt.Errorf("open writer error, %w", err)
	}
	appCloser.AddCloser(writer)

	go func() {
		var tickService service = fetchticks.NewTickFetcherService(pgRepo, writer)
		err = tickService.FetchFromDBtoParquet(ctx, *w.config.DateStarted, *w.config.DateEnded)
		if err != nil {
			log.Error("importing service", zap.Error(err))
		}
		// Finish all the app work.
		appCloser.Close()
	}()

	return nil
}
