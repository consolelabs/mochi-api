package main

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/job"
	"github.com/defipod/mochi/pkg/logger"
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
	svc, err := service.NewService(cfg, log)
	if err != nil {
		log.Fatal(err, "service.NewService failed")
		return
	}
	if err := job.NewCommonwealthProposalData(entities.Get(), svc, log).Run(); err != nil {
		log.Fatal(err, "failed to run job")
	}
	log.Info("done")
}
