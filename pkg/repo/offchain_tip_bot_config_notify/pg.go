package offchain_tip_bot_config_notify

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
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

func (pg *pg) Create(config *model.OffchainTipBotConfigNotify) error {
	return pg.db.Create(&config).Error
}

func (pg *pg) Delete(id string) error {
	return pg.db.Where("id = ?", id).Delete(&model.OffchainTipBotConfigNotify{}).Error
}
