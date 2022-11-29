package offchain_tip_bot_config_notify

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (rs []model.OffchainTipBotConfigNotify, err error)
}
