package entities

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateConfigNotify(req request.CreateTipConfigNotify) error {
	if req.Token == "all" {
		req.Token = "*"
	}

	if req.Token != "*" {
		_, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(req.Token))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[repo.OffchainTipBotTokens.GetBySymbol] - Unsupported tokens")
				return fmt.Errorf("%s hasn't been supported.\n:point_right: Please choose one in our supported `$token list`!\n:point_right: Add your token by `$token add-custom` or `$token add`.", strings.ToUpper(req.Token))
			}
			e.log.Error(err, "[entities.CreateConfigNotify] - failed to get token")
			return err
		}
	}

	return e.repo.OffchainTipBotConfigNotify.Create(&model.OffchainTipBotConfigNotify{
		GuildID:   req.GuildId,
		ChannelID: req.ChannelId,
		Token:     strings.ToUpper(req.Token),
	})
}

func (e *Entity) ListConfigNotify(guildId string) (rs []response.ConfigNotifyResponse, err error) {
	configs, err := e.repo.OffchainTipBotConfigNotify.GetByGuildID(guildId)
	if err != nil {
		e.log.Error(err, "[entities.ListConfigNotify] - failed to get config notify")
		return nil, err
	}

	configModels := make([]response.ConfigNotifyResponse, 0)
	for _, config := range configs {
		configModels = append(configModels, response.ConfigNotifyResponse{
			Id:        config.ID.String(),
			GuildId:   config.GuildID,
			ChannelId: config.ChannelID,
			Token:     config.Token,
			CreatedAt: config.CreatedAt.String(),
			UpdatedAt: config.UpdatedAt.String(),
		})
	}
	return configModels, nil
}

func (e *Entity) DeleteConfigNotify(id string) error {
	return e.repo.OffchainTipBotConfigNotify.Delete(id)
}
