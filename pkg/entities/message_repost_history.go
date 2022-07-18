package entities

import (
	"errors"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
)

func (e *Entity) IsRepostableMessage(req request.CreateMessageRepostHistRequest) bool {
	_, msgErr := e.repo.MessageRepostHistory.GetByMessageID(req.GuildID, req.MessageID)
	_, channelErr := e.repo.GuildConfigRepostReaction.GetByRepostChannelID(req.GuildID, req.ChannelID)
	if errors.Is(msgErr, gorm.ErrRecordNotFound) && errors.Is(channelErr, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (e *Entity) CreateRepostMessageHist(req request.CreateMessageRepostHistRequest, repostChannelID string) error {
	return e.repo.MessageRepostHistory.CreateIfNotExist(model.MessageRepostHistory{
		GuildID:         req.GuildID,
		OriginMessageID: req.MessageID,
		OriginChannelID: req.ChannelID,
		RepostChannelID: repostChannelID,
	})
}
