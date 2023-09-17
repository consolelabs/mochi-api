package job

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/consolelabs/mochi-typeset/typeset"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/kafka/message"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
)

type watchKeyPriceChanges struct {
	entity *entities.Entity
	log    logger.Logger
	cfg    config.Config
	cache  cache.Cache
}

const (
	defaultInterval = "hour"
)

func NewWatchKeyPriceChange(e *entities.Entity, l logger.Logger) Job {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())

	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal(err, "[WatchCoinPriceChanges] failed to init redis")
	}

	c, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatal(err, "[WatchCoinPriceChanges] failed to init redis cache")
	}

	return &watchKeyPriceChanges{
		entity: e,
		log:    l,
		cfg:    cfg,
		cache:  c,
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

	now := time.Now().UTC()
	currentDay := now.Day()
	currentPrice := decimal.NewFromInt(0)
	yesterdayClosedPrice := decimal.NewFromInt(0)
	for k, items := range keyAddressMap {
		keyAddress, err := job.entity.SearchFriendTechKeys(request.SearchFriendTechKeysRequest{
			Query: k,
			Limit: 1,
		})
		if err != nil {
			job.log.Error(err, "[watchKeyPriceChanges.Run] entity.SearchFriendTechKeys failed")
			continue
		}

		if len(keyAddress.Data) == 0 {
			job.log.Error(err, "[watchKeyPriceChanges.Run] not found key address")
			continue
		}

		priceHistories, err := job.entity.GetFriendTechKeyPriceHistory(k, defaultInterval)
		if err != nil {
			job.log.Error(err, "[watchKeyPriceChanges.Run] entity.GetFriendTechKeyPriceHistory failed")
			continue
		}

		if priceHistories != nil && len(priceHistories.Data) == 0 {
			job.log.Error(err, "[watchKeyPriceChanges.Run] not found price history")
			continue
		}

		currentPrice = priceHistories.Data[len(priceHistories.Data)-1].Price

		for i := len(priceHistories.Data) - 1; i >= 0; i-- {
			if priceHistories.Data[i].Time.Day() != currentDay {
				yesterdayClosedPrice = priceHistories.Data[i].Price
				break
			}
		}

		priceChangePercentage := calculatePercentageChange(yesterdayClosedPrice, currentPrice)

		key := fmt.Sprintf("key-price-alert-%s-%v", strings.ToLower(k), currentDay)
		val := fmt.Sprintf("%v-%v", currentPrice, yesterdayClosedPrice)

		zValue, err := job.cache.GetString(key)
		if err != nil {
			job.log.Error(err, "[watchKeyPriceChanges.Run] cache.GetString failed")
			continue
		}

		if zValue == val {
			continue
		}

		expTime := time.Duration(24-now.Hour()) * time.Hour
		err = job.cache.Set(key, val, expTime)
		if err != nil {
			job.log.Error(err, "[watchKeyPriceChanges.Run] cache.Set failed")
			continue
		}

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
				addr := model.FriendTechKey{
					ID:              keyAddress.Data[0].ID,
					CreatedAt:       keyAddress.Data[0].CreatedAt,
					UpdatedAt:       keyAddress.Data[0].UpdatedAt,
					Address:         keyAddress.Data[0].Address,
					TwitterUsername: keyAddress.Data[0].TwitterUsername,
					TwitterPfpUrl:   keyAddress.Data[0].TwitterPfpUrl,
					ProfileChecked:  keyAddress.Data[0].ProfileChecked,
					Price:           keyAddress.Data[0].Price,
					Supply:          keyAddress.Data[0].Supply,
					Holders:         keyAddress.Data[0].Holders,
				}

				profileIDAlertMap[item.ProfileId] = append(profileIDAlertMap[item.ProfileId], model.FriendTechKeyPriceChangeAlertItem{
					Timestamp:      time.Now().UTC(),
					KeyAddressID:   item.KeyAddress,
					KeyAddress:     addr,
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
			addr := message.FriendTechKey{
				ID:              itm.KeyAddress.ID,
				CreatedAt:       itm.KeyAddress.CreatedAt,
				UpdatedAt:       itm.KeyAddress.UpdatedAt,
				Address:         itm.KeyAddress.Address,
				TwitterUsername: itm.KeyAddress.TwitterUsername,
				TwitterPfpUrl:   itm.KeyAddress.TwitterPfpUrl,
				ProfileChecked:  itm.KeyAddress.ProfileChecked,
				Price:           decimal.NewFromFloat(itm.KeyAddress.Price),
				Supply:          itm.KeyAddress.Supply,
				Holders:         itm.KeyAddress.Holders,
			}
			messages = append(messages, message.KeyPriceChangeAlertMessage{
				Type: typeset.NOTIFICATION_KEY_PRICE_CHANGE,
				KeyPriceChangeAlertMetadata: message.KeyPriceChangeAlertMetadata{
					Timestamp:      itm.Timestamp,
					ProfileID:      profileID,
					KeyAddressID:   itm.KeyAddressID,
					KeyAddress:     addr,
					Change:         itm.Change,
					CurrentPrice:   itm.CurrentPrice,
					YesterdayPrice: itm.YesterdayPrice,
				},
			})
		}
	}

	job.entity.PublishKeyPriceChangeMessage(messages)
}
