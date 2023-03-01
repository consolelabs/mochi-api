package job

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/util"
)

type watchDMNotifyPriceAlert struct {
	entity *entities.Entity
	log    logger.Logger
	cache  cache.Cache
	svc    *service.Service
}

func NewWatchDMNotifyPriceAlert(e *entities.Entity, l logger.Logger) Job {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())

	// init redis
	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		l.Fatal(err, "[WatchDMNotifyPriceAlert] failed to init redis")
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		l.Fatal(err, "[WatchDMNotifyPriceAlert] failed to init redis cache")
	}

	service, err := service.NewService(cfg, l)
	if err != nil {
		l.Fatal(err, "failed to init service")
	}

	return &watchDMNotifyPriceAlert{
		entity: e,
		log:    l,
		svc:    service,
		cache:  cache,
	}
}

func (job *watchDMNotifyPriceAlert) Run() error {
	var COMMUNICATION_CHANNEL = "channel_dm_price_alert"
	subcriber := job.cache.Subcribe(COMMUNICATION_CHANNEL)
	defer subcriber.Close()

	ch := subcriber.Channel()

	for msg := range ch {
		var data watchCoinPriceChangePayload
		if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
			panic(err)
		}
		if err := job.HandleNotifyDiscordUser(data); err == nil {
			job.log.Infof("Received %v", data)
		}
	}

	return nil
}

func (job *watchDMNotifyPriceAlert) HandleNotifyDiscordUser(payload watchCoinPriceChangePayload) error {
	alert, err := job.entity.GetSpecificAlert(payload.AlertID)
	if err != nil {
		return err
	}

	nowTime := time.Now().UTC()
	if nowTime.Before(alert.SnoozedTo) {
		job.log.Infof("User ID %v - Symbol %v - Next available to DM date is after %v", alert.UserDiscordID, alert.Symbol, alert.SnoozedTo)
		return nil
	}

	tradingPair := alert.Symbol + "/" + alert.Currency
	err = job.svc.Discord.SendDMUserPriceAlert(alert.UserDiscordID, tradingPair, alert.AlertType, alert.Value)
	if err != nil {
		job.log.Fields(logger.Fields{"user_discord_id": alert.UserDiscordID}).Error(err, "[job.HandleNotifyDiscordUser] svc.Discord.SendDMUserPriceAlert() failed")
		return err
	}

	if alert.Frequency == model.OnlyOnce {
		err = job.entity.RemoveTokenPriceAlert(payload.AlertID)
		if err != nil {
			job.log.Fields(logger.Fields{}).Error(err, "[job.HandleNotifyDiscordUser] entity.RemoveTokenPriceAlert() failed")
			return err
		}
		return nil
	}
	if alert.Frequency == model.OnceADay {
		alert.SnoozedTo = util.StartOfDay(time.Now().AddDate(0, 0, 1)) // update to the start of next day
		job.entity.UpdateSpecificPriceAlert(*alert)
		return nil
	}

	return nil
}
