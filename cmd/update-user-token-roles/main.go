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
		return
	}

	log.Info("start job updateUserRoles ...")
	if err := job.NewUpdateUserTokenRolesJob(entities.Get(), nil).Run(); err != nil {
		log.Fatal(err, "failed to run job")
		return
	}

	log.Info("done")
}
