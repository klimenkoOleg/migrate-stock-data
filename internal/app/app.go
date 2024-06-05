package app

import (
	"context"
	"github.com/klimenkoOleg/migrate-stock-data/internal/common"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"time"
)

type Worker interface {
	Do(ctx context.Context, appCloser *common.Closer) error
}

type App struct {
	log    *zap.Logger
	config *Config
}

func New(config *Config, log *zap.Logger) *App {
	return &App{config: config, log: log}
}

func (a *App) Run(ctx context.Context, cancel context.CancelFunc, worker Worker) {
	appCloser := common.New(5*time.Second, a.log)
	appCloser.AddCloseFunc(func() error {
		cancel()
		return nil
	})

	err := worker.Do(ctx, appCloser) // doesn't block execution
	if err != nil {
		appCloser.Close()
		log.Fatal("init error", zap.Error(err))
	}

	appCloser.Wait()
	log.Info("Application successfully stopped")
}
