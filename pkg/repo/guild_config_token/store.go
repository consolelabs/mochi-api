package guild_config_token

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) ([]model.GuildConfigToken, error)
	UpsertMany(configs []model.GuildConfigToken) error
	UpsertOne(configs model.GuildConfigToken) error
	CreateOne(token model.GuildConfigToken) error
	GetAll() ([]model.GuildConfigToken, error)
	GetByGuildIDAndTokenID(guildID string, tokenID int) (*model.GuildConfigToken, error)
}
