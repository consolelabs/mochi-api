package guild_blacklist_channel_repost_configs

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) UpsertOne(config model.GuildBlacklistChannelRepostConfig) error {
	tx := pg.db.Begin()
	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "channel_id"}},
		UpdateAll: true,
	}).Create(&config).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildBlacklistChannelRepostConfig, error) {
	var configs []model.GuildBlacklistChannelRepostConfig
	return configs, pg.db.Where("guild_id = ?", guildID).Find(&configs).Error
}

func (pg *pg) DeleteOne(guildID, channelID string) error {
	return pg.db.Where("guild_id = ? AND channel_id = ?", guildID, channelID).Delete(&model.GuildBlacklistChannelRepostConfig{}).Error
}
