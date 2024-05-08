package main

import (
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/scheduler"
	"github.com/defipod/mochi/pkg/service"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()
	// *** entities ***
	err := entities.Init(cfg, log)
	if err != nil {
		log.Fatal(err, "failed to init entities")
	}

	service, err := service.NewService(cfg, log)
	if err != nil {
		log.Fatal(err, "failed to init service")
	}
	defer service.Sentry.Flush(2 * time.Second)

	if err := scheduler.NewUpdateBinanceSpotHistory(entities.Get(), log, service, cfg).Run(); err != nil {
		log.Fatal(err, "failed to run job")
	}

	log.Info("done")
}
