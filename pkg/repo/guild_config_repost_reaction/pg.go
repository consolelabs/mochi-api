package guild_config_repost_reaction

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildConfigRepostReaction, error) {
	var configs []model.GuildConfigRepostReaction
	err := pg.db.Model(model.GuildConfigRepostReaction{}).Where("guild_id = ?", guildID).Scan(&configs).Error
	if err != nil {
		return configs, fmt.Errorf("failed to list repostable reaction configs: %w", err)
	}
	return configs, nil
}

func (pg *pg) GetByReaction(guildID string, emoji string) (model.GuildConfigRepostReaction, error) {
	var config model.GuildConfigRepostReaction
	return config, pg.db.Model(model.GuildConfigRepostReaction{}).Where("guild_id = ? AND emoji = ?", guildID, emoji).First(&config).Error
}

func (pg *pg) GetByReactionStartOrStop(guildID, emoji string) (model.GuildConfigRepostReaction, error) {
	var config model.GuildConfigRepostReaction
	return config, pg.db.Model(model.GuildConfigRepostReaction{}).Where("guild_id = ? AND (emoji_start = ? OR emoji_stop = ?)", guildID, emoji, emoji).First(&config).Error
}

func (pg *pg) GetByRepostChannelID(guildID, channelID string) (model.GuildConfigRepostReaction, error) {
	var config model.GuildConfigRepostReaction
	err := pg.db.Model(model.GuildConfigRepostReaction{}).Where("guild_id = ? AND repost_channel_id = ?", guildID, channelID).First(&config).Error
	if err != nil {
		return config, fmt.Errorf("failed to get message repost history: %w", err)
	}
	return config, nil
}

func (pg *pg) UpsertOne(config model.GuildConfigRepostReaction) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "emoji"}},
		UpdateAll: true,
	}).Create(&config).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) DeleteOne(guildID string, emoji string) error {
	return pg.db.Where("guild_id = ? AND emoji = ?", guildID, emoji).Delete(&model.GuildConfigRepostReaction{}).Error
}
