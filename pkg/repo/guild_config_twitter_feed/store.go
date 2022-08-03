package guild_config_twitter_feed

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(*model.GuildConfigTwitterFeed) error
	GetAll() ([]model.GuildConfigTwitterFeed, error)
}
