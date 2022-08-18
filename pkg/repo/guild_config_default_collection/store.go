package guild_config_default_collection

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildIDandChainID(guildID string, chainID string) ([]model.GuildConfigDefaultCollection, error)
	Upsert(*model.GuildConfigDefaultCollection) error
}
