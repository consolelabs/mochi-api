package guild_config_default_ticker

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOneByGuildIDAndQuery(guildID, query string) (*model.GuildConfigDefaultTicker, error)
	UpsertOne(config *model.GuildConfigDefaultTicker) error
}
