package entities

import (
	"errors"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) UpdateGuildStats(stat model.DiscordGuildStat) error {
	err := e.repo.DiscordGuildStats.UpsertOne(stat)
	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) GetByGuildID(guildID string) (*model.DiscordGuildStat, error) {
	guildStat, err := e.repo.DiscordGuildStats.GetByGuildID(guildID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &model.DiscordGuildStat{
		ID:          guildStat.ID,
		GuildID:     guildStat.GuildID,
		NrOfMembers: guildStat.NrOfMembers,
		NrOfUsers:   guildStat.NrOfUsers,
		NrOfBots:    guildStat.NrOfBots,

		NrOfChannels:             guildStat.NrOfChannels,
		NrOfTextChannels:         guildStat.NrOfTextChannels,
		NrOfVoiceChannels:        guildStat.NrOfVoiceChannels,
		NrOfStageChannels:        guildStat.NrOfStageChannels,
		NrOfCategories:           guildStat.NrOfCategories,
		NrOfAnnouncementChannels: guildStat.NrOfAnnouncementChannels,

		NrOfEmojis:         guildStat.NrOfEmojis,
		NrOfStaticEmojis:   guildStat.NrOfStaticEmojis,
		NrOfAnimatedEmojis: guildStat.NrOfAnimatedEmojis,

		NrOfStickers:       guildStat.NrOfStickers,
		NrOfCustomStickers: guildStat.NrOfCustomStickers,
		NrOfServerStickers: guildStat.NrOfServerStickers,
		NrOfRoles:          guildStat.NrOfRoles,
		CreatedAt:          guildStat.CreatedAt,
	}, nil
}
