package guild_config_prune_exclude

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) ([]model.GuildConfigWhitelistPrune, error)
	UpsertOne(config *model.GuildConfigWhitelistPrune) error
	DeleteOne(config *model.GuildConfigWhitelistPrune) error
}
