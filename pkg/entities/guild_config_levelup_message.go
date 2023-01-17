package entities

import (
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
)

func (e *Entity) UpsertLevelUpMessage(req request.UpsertGuildLevelUpMessageRequest) (*model.GuildConfigLevelupMessage, error) {
	config, err := e.repo.GuildConfigLevelUpMessage.UpsertOne(model.GuildConfigLevelupMessage{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		ImageURL:  req.ImageURL,
		Message:   req.Message,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.UpsertLevelUpMessage] e.repo.GuildConfigLevelUpMessage.UpsertOne failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) GetLevelUpMessage(guildId string) (*model.GuildConfigLevelupMessage, error) {
	config, err := e.repo.GuildConfigLevelUpMessage.GetByGuildId(guildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[entity.GetLevelUpMessage] e.repo.GuildConfigLevelUpMessage.GetByGuildId failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) DeleteLevelUpMessage(req request.GuildIDRequest) error {
	err := e.repo.GuildConfigLevelUpMessage.DeleteByGuildId(req.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[entity.DeleteLevelUpMessage] e.repo.GuildConfigLevelUpMessage.DeleteByGuildId failed")
		return err
	}
	return nil
}
