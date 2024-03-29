package guild_config_default_collection

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) ([]model.GuildConfigDefaultCollection, error)
	Upsert(*model.GuildConfigDefaultCollection) error
	GetOneByGuildIDAndQuery(guildID, symbol string) (*model.GuildConfigDefaultCollection, error)
}
