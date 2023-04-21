package guild_config_tip_range

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetByGuildID(guildID string) (*model.GuildConfigTipRange, error) {
	var configs model.GuildConfigTipRange
	return &configs, pg.db.Where("guild_id = ?", guildID).First(&configs).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigTipRange) (*model.GuildConfigTipRange, error) {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&config).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return config, tx.Commit().Error
}

func (pg *pg) Remove(guildID string) error {
	return pg.db.
		Where("guild_id = ?", guildID).
		Delete(&model.GuildConfigTipRange{}).Error
}
