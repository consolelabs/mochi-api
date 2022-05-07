package guild_config_token

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) ([]model.GuildConfigToken, error)
	UpsertOne(config *model.GuildConfigToken) error
}
