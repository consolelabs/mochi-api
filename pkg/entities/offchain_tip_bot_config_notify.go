package entities

import (
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateConfigNotify(req request.CreateTipConfigNotify) error {
	if req.Token == "all" {
		req.Token = "*"
	}

	return e.repo.OffchainTipBotConfigNotify.Create(&model.OffchainTipBotConfigNotify{
		GuildID:   req.GuildId,
		ChannelID: req.ChannelId,
		Token:     strings.ToUpper(req.Token),
	})
}

func (e *Entity) ListConfigNotify(guildId string) (rs []model.OffchainTipBotConfigNotify, err error) {
	return e.repo.OffchainTipBotConfigNotify.GetByGuildID(guildId)
}

func (e *Entity) DeleteConfigNotify(id string) error {
	return e.repo.OffchainTipBotConfigNotify.Delete(id)
}
