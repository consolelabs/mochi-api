package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

func (e *Entity) CreateDiscordGuildStatChannel(statChannel model.DiscordGuildStatChannel) error {
	err := e.repo.DiscordGuildStatChannels.Create(&statChannel)
	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) GetStatChannelsByGuildID(guildID string) ([]model.DiscordGuildStatChannel, error) {
	log := logger.NewLogrusLogger()
	log.Infof("Get stats channel for guild. GuildId: %v", guildID)
	statChannels, err := e.repo.DiscordGuildStatChannels.GetStatChannelsByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return statChannels, nil
}

func (e *Entity) DeleteStatChannelByChannelID(channelID string) error {
	log := logger.NewLogrusLogger()
	log.Infof("Delete unused stat channels, channelId: %v", channelID)
	err := e.repo.DiscordGuildStatChannels.DeleteStatChannelByChannelID(channelID)
	if err != nil {
		return err
	}
	return nil
}
