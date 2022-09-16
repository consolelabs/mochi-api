package main

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/job"
	"github.com/defipod/mochi/pkg/logger"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()
	err := entities.Init(cfg, log)
	if err != nil {
		log.Fatal(err, "failed to init entities")
	}

	log.Info("start job updateGMStreak ...")
	if err := job.NewUpdateGMStreakJob(entities.Get(), log).Run(); err != nil {
		log.Fatal(err, "failed to run job")
	}

	log.Info("done")
}
