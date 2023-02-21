package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) SetUpvoteMessageCache(req *request.SetUpvoteMessageCacheRequest) error {
	// "upvote_msg_<userID>": "<messageID>_<channelID>_<guildID>"

	key := fmt.Sprintf("upvote_msg_%s", req.UserID)
	value := fmt.Sprintf("%s_%s_%s", req.MessageID, req.ChannelID, req.GuildID)
	err := e.cache.Set(key, value, time.Duration(time.Minute*5))
	if err != nil {
		e.log.Fields(logger.Fields{"key": key, "value": value}).Errorf(err, "[e.SetUpvoteMessageCache] failed to set cache")
		return err
	}
	return nil
}

func (e *Entity) GetUpvoteMessageCache(userID string) (*response.SetUpvoteMessageCacheResponse, error) {
	key := fmt.Sprintf("upvote_msg_%s", userID)
	data, err := e.cache.GetString(key)
	if err != nil {
		e.log.Fields(logger.Fields{"key": key, "value": data}).Errorf(err, "[e.GetUpvoteMessageCache] failed to get cache")
		return nil, err
	}
	if data == "" {
		return nil, nil
	}

	ids := strings.Split(data, "_")
	if len(ids) != 3 {
		err = fmt.Errorf("cached data invalid: %s", data)
		e.log.Fields(logger.Fields{"key": key, "value": data}).Errorf(err, "[e.GetUpvoteMessageCache] cached data invalid form")
		return nil, err
	}

	return &response.SetUpvoteMessageCacheResponse{
		UserID:    userID,
		MessageID: ids[0],
		ChannelID: ids[1],
		GuildID:   ids[2],
	}, nil
}

func (e *Entity) RemoveUpvoteMessageCache(userID string) error {
	key := fmt.Sprintf("upvote_msg_%s", userID)
	err := e.cache.Remove(key)
	if err != nil {
		e.log.Fields(logger.Fields{"key": key}).Errorf(err, "[e.RemoveUpvoteMessageCache] failed to delete cache")
		return err
	}
	return nil
}

func (e *Entity) GetTokenAlertZCache(symbol, trend, min, max string) []string {
	return e.cache.GetStringSorted(fmt.Sprintf("alert_%s_%s", symbol, trend), min, max)
}

func (e *Entity) DeleteTokenAlertZCache(symbol, trend, value string) error {
	return e.cache.ZRemove(fmt.Sprintf("alert_%s_%s", symbol, trend), value)
}

func (e *Entity) GetPriceAlertZCache(symbol, trend, min, max string) []response.ZSetWithScoreData {
	return e.cache.GetStringSortedWithScores(fmt.Sprintf("alert_direction_%s:%s", trend, symbol), min, max)
}

func (e *Entity) RemovePriceAlertZCache(symbol, trend, value string) error {
	return e.cache.ZRemove(fmt.Sprintf("alert_direction_%s:%s", trend, symbol), value)
}

func (e *Entity) Publish(channel string, payload interface{}) error {
	return e.cache.Publish(channel, payload)
}

func (e *Entity) Subcribe(channel string) *redis.PubSub {
	return e.cache.Subcribe(channel)
}
