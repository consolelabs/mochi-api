package guild_config_gm_gn

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigGmGn, error)
	GetAllByGuildID(guildID string) ([]model.GuildConfigGmGn, error)
	GetLatestByGuildID(guildID string) (config []model.GuildConfigGmGn, err error)
	UpsertOne(config *model.GuildConfigGmGn) error
	CreateOne(config *model.GuildConfigGmGn) error
}
