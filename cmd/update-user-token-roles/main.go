package main

import (
	"time"

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
	entity := entities.Get()
	defer entity.GetSvc().Sentry.Flush(2 * time.Second)

	log.Info("start job updateUserRoles ...")
	if err := job.NewUpdateUserTokenRolesJob(entity, nil).Run(); err != nil {
		log.Fatal(err, "failed to run job")
		return
	}

	log.Info("done")
}
