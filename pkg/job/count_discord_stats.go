package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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
		// update data stats in database
		err := c.entity.UpdateOneGuildStats(guild.ID)
		if err != nil {
			return err
		}
		// update channel name in guilds
		statChannels, err := c.entity.GetStatChannelsByGuildID(guild.ID)
		if err != nil {
			return err
		}
		for _, statChannel := range statChannels {
			err = c.entity.EditGuildChannel(guild.ID, statChannel)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
