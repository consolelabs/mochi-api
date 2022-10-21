package guild_blacklist_channel_repost_configs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(config model.GuildBlacklistChannelRepostConfig) error
	GetByGuildID(guildID string) ([]model.GuildBlacklistChannelRepostConfig, error)
	DeleteOne(guildID, channelID string) error
}
