package main

import (
	"context"
	"github.com/klimenkoOleg/migrate-stock-data/internal/app"
	"github.com/klimenkoOleg/migrate-stock-data/internal/common"
	"github.com/klimenkoOleg/migrate-stock-data/internal/worker/fetcher"
)

func main() {
	log := common.NewDefaultLogger()
	ctx, cancel := context.WithCancel(context.Background())

	config := app.MustLoadConfig(log, "conf/fetcher.yml")
	if config.DateStarted == nil || config.DateEnded == nil {
		log.Fatal("date.start and date.end should be specified in config")
	}

	app := app.New(config, log)
	app.Run(ctx, cancel, fetcher.NewWorker(config, log))
}
