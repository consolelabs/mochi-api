package entities

import (
	"errors"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) IsRepostableMessage(req request.MessageReactionRequest) bool {
	_, msgErr := e.repo.MessageRepostHistory.GetByMessageID(req.GuildID, req.MessageID)
	_, channelErr := e.repo.GuildConfigRepostReaction.GetByRepostChannelID(req.GuildID, req.ChannelID)
	if errors.Is(msgErr, gorm.ErrRecordNotFound) && errors.Is(channelErr, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (e *Entity) CreateRepostMessageHistory(req request.MessageReactionRequest) (*model.MessageRepostHistory, error) {
	repostMsg, err := e.repo.MessageRepostHistory.GetByMessageID(req.GuildID, req.MessageID)
	if err == nil {
		return repostMsg, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Error(err, "[entities.CreateRepostMessageHist] - failed to get repost message")
		return nil, err
	}
	// case not exists
	history := &model.MessageRepostHistory{
		GuildID:         req.GuildID,
		OriginMessageID: req.MessageID,
		OriginChannelID: req.ChannelID,
		RepostChannelID: req.RepostChannelID,
		ReactionCount:   int64(req.ReactionCount),
		IsStart:         req.IsStart,
		IsStop:          req.IsStop,
	}

	err = e.repo.MessageRepostHistory.Upsert(*history)
	if err != nil {
		e.log.Error(err, "[entities.CreateRepostMessageHist] - failed to get repost message")
		return nil, err
	}

	return history, nil
}

func (e *Entity) GetMessageRepostHistory(req request.MessageReactionRequest) (*model.MessageRepostHistory, error) {
	repostMsg, err := e.repo.MessageRepostHistory.GetByMessageID(req.GuildID, req.MessageID)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetMessageRepostHistory] repo.MessageRepostHistory.GetByMessageID() failed")
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return repostMsg, nil
}
