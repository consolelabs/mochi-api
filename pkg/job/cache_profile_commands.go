package job

import (
	"log"
	"sync"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type cacheProfileCommands struct {
	logger logger.Logger
	entity *entities.Entity
}

func NewCacheProfileCommands(l logger.Logger, cfg config.Config) Job {
	err := entities.Init(cfg, l)
	if err != nil {
		log.Fatal(err, "failed to init entities")
	}

	return &cacheProfileCommands{
		logger: l,
		entity: entities.Get(),
	}
}

func (j *cacheProfileCommands) Run() error {
	j.logger.Infof("Start caching profile commands ...")

	list, err := j.entity.GetRepo().ProfileCommandUsage.GetTopProfileUsage(5)
	if err != nil {
		j.logger.Error(err, "repo.ProfileCommandUsage.GetTopProfileUsage() failed")
		return err
	}

	var wg sync.WaitGroup
	wg.Add(6 * len(list))

	for _, u := range list {
		l := j.logger.Fields(logger.Fields{"profile_id": u.ProfileId, "platform_id": u.UserPlatformId, "platform": u.Platform})
		l.Info("start caching /profile data")

		if u.Platform == "discord" {
			go func(u model.CommandUsageCounter) {
				_, err = j.entity.GetSvc().MochiProfile.GetByDiscordID(u.UserPlatformId, false)
				if err != nil {
					l.Error(err, "svc.MochiProfile.GetByDiscordID() failed")
				}
				wg.Done()
			}(u)
		} else if u.Platform == "telegram" {
			go func(u model.CommandUsageCounter) {
				_, err = j.entity.GetSvc().MochiProfile.GetByTelegramID(u.UserPlatformId, false)
				if err != nil {
					l.Error(err, "svc.MochiProfile.GetByTelegramID() failed")
				}
				wg.Done()
			}(u)
		} else {
			wg.Done()
		}

		go func(u model.CommandUsageCounter) {
			_, err = j.entity.GetSvc().MochiProfile.GetProfileActivities(u.ProfileId)
			if err != nil {
				l.Error(err, "svc.MochiProfile.GetProfileActivities() failed")
			}
			wg.Done()
		}(u)

		go func(u model.CommandUsageCounter) {
			_, err = j.entity.GetSvc().MochiPay.GetBalance(u.ProfileId, "", "")
			if err != nil {
				l.Error(err, "svc.MochiPay.GetBalance() failed")
			}
			wg.Done()
		}(u)

		go func(u model.CommandUsageCounter) {
			_, err = j.entity.GetSvc().MochiPay.GetProfileCustodialWallets(u.ProfileId)
			if err != nil {
				l.Error(err, "svc.MochiPay.GetProfileCustodialWallets() failed")
			}
			wg.Done()
		}(u)

		go func(u model.CommandUsageCounter) {
			_, err = j.entity.GetVaults(request.GetVaultsRequest{ProfileID: u.ProfileId})
			if err != nil {
				l.Error(err, "entity.GetVaults() failed")
			}
			wg.Done()
		}(u)
	}

	wg.Wait()

	return nil
}
