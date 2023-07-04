package entities

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetTokenAlertZCache(symbol, trend, min, max string) []string {
	return e.cache.GetStringSorted(fmt.Sprintf("alert_%s_%s", symbol, trend), min, max)
}

func (e *Entity) DeleteTokenAlertZCache(symbol, trend, value string) error {
	return e.cache.ZRemove(fmt.Sprintf("alert_%s_%s", symbol, trend), value)
}

func (e *Entity) GetPriceAlertZCache(symbol, trend, min, max string) []response.ZSetWithScoreData {
	return e.cache.GetStringSortedWithScores(fmt.Sprintf("alert_direction_%s:%s", trend, symbol), min, max)
}

func (e *Entity) RemovePriceAlertZCache(symbol, trend, price string) error {
	if trend == "up" {
		return e.cache.ZRemoveByScore(fmt.Sprintf("alert_direction_%s:%s", trend, symbol), "0", price)
	} else {
		return e.cache.ZRemoveByScore(fmt.Sprintf("alert_direction_%s:%s", trend, symbol), price, "inf")
	}
}

func (e *Entity) Publish(channel string, payload interface{}) error {
	return e.cache.Publish(channel, payload)
}

func (e *Entity) Subcribe(channel string) *redis.PubSub {
	return e.cache.Subcribe(channel)
}
