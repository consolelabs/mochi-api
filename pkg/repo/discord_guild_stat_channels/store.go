package discord_guild_stat_channels

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(statChannel *model.DiscordGuildStatChannel) error
	GetStatChannelsByGuildID(guildID string) ([]model.DiscordGuildStatChannel, error)
	DeleteStatChannelByChannelID(channelID string) error
}
