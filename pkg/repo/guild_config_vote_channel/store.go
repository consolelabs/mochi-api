package guild_config_vote_channel

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigVoteChannel, error)
	UpsertOne(config *model.GuildConfigVoteChannel) (*model.GuildConfigVoteChannel, error)
	DeleteOne(config *model.GuildConfigVoteChannel) error
}
