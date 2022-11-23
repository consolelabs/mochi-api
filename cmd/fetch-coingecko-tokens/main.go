package main

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/job"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/vault"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	vault, err := vault.New(&cfg)
	if err != nil {
		log.Error(err, "failted to init vault")
	}

	if vault != nil {
		cfg = *vault.LoadConfig()
	}

	// *** entities ***
	err = entities.Init(cfg, log)
	if err != nil {
		log.Fatal(err, "failed to init entities")
	}

	if err := job.NewFetchCoingeckoTokensJob(entities.Get(), log).Run(); err != nil {
		log.Fatal(err, "failed to run job")
	}

	log.Info("done")
}
