package job

import (
	"github.com/consolelabs/mochi-typeset/typeset"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/kafka/message"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/shopspring/decimal"
)

type watchKeyPriceChanges struct {
	entity *entities.Entity
	log    logger.Logger
	cfg    config.Config
}

const (
	defaultInterval = "hour"
)

func NewWatchKeyPriceChange(e *entities.Entity, l logger.Logger) Job {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())

	return &watchKeyPriceChanges{
		entity: e,
		log:    l,
		cfg:    cfg,
	}
}

func (job *watchKeyPriceChanges) Run() error {
	watchList, err := job.entity.GetFriendTechKeyWatchlist()
	if err != nil {
		job.log.Error(err, "[watchKeyPriceChanges.Run] entity.GetFriendTechKeyWatchlist failed")
		return err
	}

	if len(watchList) == 0 {
		job.log.Info("[watchKeyPriceChanges.Run] No key to watch")
		return nil
	}

	profileIDAlertMap := make(map[string][]model.FriendTechKeyPriceChangeAlertItem)
	keyAddressMap := make(map[string][]model.FriendTechKeyWatchlistItem)

	for _, item := range watchList {
		keyAddressMap[item.KeyAddress] = append(keyAddressMap[item.KeyAddress], item)
	}

	currentDay := time.Now().UTC().Day()
	currentPrice := decimal.NewFromInt(0)
	yesterdayClosedPrice := decimal.NewFromInt(0)
	priceChangePercentage := decimal.NewFromInt(0)
	for k, items := range keyAddressMap {
		priceHistories, err := job.entity.GetFriendTechKeyPriceHistory(k, defaultInterval)
		if err != nil {
			job.log.Error(err, "[watchKeyPriceChanges.Run] entity.GetFriendTechKeyPriceHistory failed")
			continue
		}

		currentPrice = priceHistories.Data[len(priceHistories.Data)-1].Price

		for i := len(priceHistories.Data) - 1; i >= 0; i-- {
			if priceHistories.Data[i].Time.Day() != currentDay {
				yesterdayClosedPrice = priceHistories.Data[i].Price
				break
			}
		}

		priceChangePercentage = calculatePercentageChange(yesterdayClosedPrice, currentPrice)
		for _, item := range items {
			increaseAlertAt := decimal.NewFromInt(int64(item.IncreaseAlertAt))
			decreaseAlertAt := decimal.NewFromInt(int64(item.DecreaseAlertAt))

			addAlert := false
			if !increaseAlertAt.Equals(decimal.NewFromInt(0)) && priceChangePercentage.GreaterThan(increaseAlertAt) {
				addAlert = true
			}

			if !decreaseAlertAt.Equals(decimal.NewFromInt(0)) && priceChangePercentage.LessThan(decreaseAlertAt.Mul(decimal.NewFromInt(-1))) {
				addAlert = true
			}

			if addAlert {
				profileIDAlertMap[item.ProfileId] = append(profileIDAlertMap[item.ProfileId], model.FriendTechKeyPriceChangeAlertItem{
					KeyAddressID:   item.KeyAddress,
					Change:         priceChangePercentage,
					CurrentPrice:   currentPrice,
					YesterdayPrice: yesterdayClosedPrice,
				})
			}
		}
	}

	job.publishMessage(profileIDAlertMap)
	return nil
}

func calculatePercentageChange(yesterdayPrice, todayPrice decimal.Decimal) decimal.Decimal {
	if yesterdayPrice.IsZero() {
		return decimal.Zero
	}
	return ((todayPrice.Sub(yesterdayPrice)).Div(yesterdayPrice)).Mul(decimal.NewFromInt(100))
}

func (job *watchKeyPriceChanges) publishMessage(in map[string][]model.FriendTechKeyPriceChangeAlertItem) {
	messages := make([]message.KeyPriceChangeAlertMessage, 0)
	for profileID, pc := range in {
		for _, itm := range pc {
			messages = append(messages, message.KeyPriceChangeAlertMessage{
				Type: typeset.NOTIFICATION_KEY_PRICE_CHANGE,
				KeyPriceChangeAlertMetadata: message.KeyPriceChangeAlertMetadata{
					ProfileID:      profileID,
					KeyAddressID:   itm.KeyAddressID,
					Change:         itm.Change,
					CurrentPrice:   itm.CurrentPrice,
					YesterdayPrice: itm.YesterdayPrice,
				},
			})
		}
	}

	job.entity.PublishKeyPriceChangeMessage(messages)
}
