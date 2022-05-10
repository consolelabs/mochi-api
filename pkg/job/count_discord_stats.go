package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type countDiscordStats struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewCountDiscordStatsJob(e *entities.Entity, l logger.Logger) Job {
	return &countDiscordStats{
		entity: e,
		log:    l,
	}
}

func (c *countDiscordStats) Run() error {
	guilds, err := c.entity.GetGuilds()
	if err != nil {
		return err
	}

	for _, guild := range guilds.Data {
		// check if mochi bot has permission in guild
		guild, _ := c.entity.GetGuildById(guild.ID)
		if guild == nil {
			continue
		}

		// update data stats in database
		err := c.entity.UpdateOneGuildStats(guild.ID)
		if err != nil {
			c.log.Error(err, "failed to update guilds stat for "+guild.ID)
		}
		// update channel name in guilds
		statChannels, err := c.entity.GetStatChannelsByGuildID(guild.ID)
		if err != nil {
			c.log.Error(err, "failed to get stats channel for guild "+guild.ID)
		}

		// check if channels is deleted, then not update and delete from db
		existChannels := make([]model.DiscordGuildStatChannel, 0)
		for _, statChannel := range statChannels {
			channel, _ := c.entity.GetGuildChannel(statChannel.ChannelID)
			if channel != nil {
				existChannels = append(existChannels, statChannel)
			} else {
				_ = c.entity.DeleteStatChannelByChannelID(statChannel.ChannelID)
			}
		}
		// update channels exist
		for _, statChannel := range existChannels {
			err = c.entity.EditGuildChannel(guild.ID, statChannel)
			if err != nil {
				c.log.Error(err, "failed to edit channel name for "+guild.ID)
			}
		}
	}

	return nil
}
