package entities

import (
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
	statChannels, err := e.repo.DiscordGuildStatChannels.GetStatChannelsByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return statChannels, nil
}

func (e *Entity) DeleteStatChannelByChannelID(channelID string) error {
	err := e.repo.DiscordGuildStatChannels.DeleteStatChannelByChannelID(channelID)
	if err != nil {
		return err
	}
	return nil
}
