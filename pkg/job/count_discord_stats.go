package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type countDiscordStats struct {
	entity *entities.Entity
}

func NewCountDiscordStatsJob(e *entities.Entity, l logger.Logger) Job {
	return &countDiscordStats{
		entity: e,
	}
}

func (c *countDiscordStats) Run() error {
	guilds, err := c.entity.GetGuilds()
	if err != nil {
		return err
	}

	for _, guild := range guilds.Data {
		nr_of_members, nr_of_user, nr_of_bots, err := c.entity.CountGuildMembers(guild.ID)
		if err != nil {
			return err
		}
		nr_of_channels, nr_of_text_channels, nr_of_voice_channels, nr_of_stage_channels, nr_of_categories, nr_of_announcement_channels, err := c.entity.CountGuildChannels(guild.ID)
		if err != nil {
			return err
		}
		nr_of_emojis, nr_of_static_emojis, nr_of_animated_emojis, err := c.entity.CountGuildEmojis(guild.ID)
		if err != nil {
			return err
		}
		nr_of_stickers, nr_of_standard_stickers, nr_of_guild_stickers, err := c.entity.CountGuildStickers(guild.ID)
		if err != nil {
			return err
		}
		nr_of_roles, err := c.entity.CountGuildRoles(guild.ID)
		if err != nil {
			return err
		}

		// update stats to database
		err = c.entity.UpdateGuildStats(model.DiscordGuildStat{
			GuildID:     guild.ID,
			NrOfMembers: nr_of_members,
			NrOfUsers:   nr_of_user,
			NrOfBots:    nr_of_bots,

			NrOfChannels:             nr_of_channels,
			NrOfTextChannels:         nr_of_text_channels,
			NrOfVoiceChannels:        nr_of_voice_channels,
			NrOfStageChannels:        nr_of_stage_channels,
			NrOfCategories:           nr_of_categories,
			NrOfAnnouncementChannels: nr_of_announcement_channels,

			NrOfEmojis:         nr_of_emojis,
			NrOfStaticEmojis:   nr_of_static_emojis,
			NrOfAnimatedEmojis: nr_of_animated_emojis,

			NrOfStickers:         nr_of_stickers,
			NrOfStandardStickers: nr_of_standard_stickers,
			NrOfGuildStickers:    nr_of_guild_stickers,
			NrOfRoles:            nr_of_roles,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
