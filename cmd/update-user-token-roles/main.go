package main

import (
	"time"

	"github.com/getsentry/sentry-go"

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

	err = sentry.Init(sentry.ClientOptions{
		Dsn:        "https://b632003ad0874c5182ee572bb7ad3c6c@sentry.daf.ug/4",
		Debug:      true,
		ServerName: "mochi-update-user-token-roles",
	})
	if err != nil {
		log.Fatal(err, "can't init sentry service")
	}
	defer sentry.Flush(2 * time.Second)

	log.Info("start job updateUserRoles ...")
	if err := job.NewUpdateUserTokenRolesJob(entities.Get(), sentry.CurrentHub().Client(), nil).Run(); err != nil {
		log.Fatal(err, "failed to run job")
		return
	}

	log.Info("done")
}
