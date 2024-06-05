package main

import (
	"context"
	"github.com/klimenkoOleg/migrate-stock-data/internal/app"
	"github.com/klimenkoOleg/migrate-stock-data/internal/common"
	"github.com/klimenkoOleg/migrate-stock-data/internal/worker/uploader"
)

func main() {
	log := common.NewDefaultLogger()
	ctx, cancel := context.WithCancel(context.Background())

	config := app.MustLoadConfig(log, "conf/uploader.yml")
	app := app.New(config, log)

	worker := uploader.NewWorker(config, log)
	app.Run(ctx, cancel, worker)
}
