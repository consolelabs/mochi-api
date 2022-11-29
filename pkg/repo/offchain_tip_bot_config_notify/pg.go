package offchain_tip_bot_config_notify

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetByGuildID(guildID string) (rs []model.OffchainTipBotConfigNotify, err error) {
	return rs, pg.db.Where("guild_id = ?", guildID).Find(&rs).Error
}
