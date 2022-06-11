package entities

import (
	"errors"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
)

func (e *Entity) IsRepostableMessage(guildID, messageID string) bool {
	_, err := e.repo.MessageRepostHistory.GetByMessageID(guildID, messageID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
