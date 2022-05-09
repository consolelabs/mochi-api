package discord_guild_stat_channels

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) Create(statChannel *model.DiscordGuildStatChannel) error {
	return pg.db.Create(statChannel).Error
}

func (pg *pg) GetStatChannelsByGuildID(guildID string) ([]model.DiscordGuildStatChannel, error) {
	statChannels := []model.DiscordGuildStatChannel{}
	err := pg.db.Where("guild_id = ?", guildID).Find(&statChannels).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get guilds: %w", err)
	}

	return statChannels, nil
}

func (pg *pg) DeleteStatChannelByChannelID(channelID string) error {
	statChannel := model.DiscordGuildStatChannel{}
	err := pg.db.Where("channel_id = ?", channelID).Delete(statChannel).Error
	if err != nil {
		return fmt.Errorf("failed to delete stat channel: %w", err)
	}
	return nil
}
