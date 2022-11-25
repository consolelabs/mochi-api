package guild_config_default_currency

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigDefaultCurrency, error)
	Upsert(*model.UpsertGuildConfigDefaultCurrency) error
	DeleteByGuildID(guildID string) error
}
