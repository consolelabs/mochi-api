package main

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/job"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
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

	err = entities.Init(cfg, log)
	if err != nil {
		log.Fatal(err, "failed to init entities")
		return
	}

	svc, err := service.NewService(cfg, log)
	if err != nil {
		log.Fatal(err, "service.NewService failed")
		return
	}

	log.Info("start job clean up expired assign contract ...")
	if err := job.NewCleanUpExpiredAssignContractJob(entities.Get(), svc, log).Run(); err != nil {
		log.Fatal(err, "failed to run job")
		return
	}

	log.Info("done")
}
