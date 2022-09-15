package entities

import (
	"fmt"
	"strings"
	"time"

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
