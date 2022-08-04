package guild_config_twitter_hashtag

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(*model.GuildConfigTwitterHashtag) error
	GetByGuildID(guildID string) (*model.GuildConfigTwitterHashtag, error)
	DeleteByGuildID(guildID string) error
}
