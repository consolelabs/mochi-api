package guild_config_gm_gn

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigGmGn, error)
	UpsertOne(config *model.GuildConfigGmGn) error
}
