package guild_config_join_leave_channel

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigJoinLeaveChannel, error)
	UpsertOne(config *model.GuildConfigJoinLeaveChannel) (*model.GuildConfigJoinLeaveChannel, error)
	DeleteOne(config *model.GuildConfigJoinLeaveChannel) error
}
