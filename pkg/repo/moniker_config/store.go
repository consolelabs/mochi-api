package moniker_config

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) ([]model.MonikerConfig, error)
	UpsertOne(record model.MonikerConfig) error
	DeleteOne(guildID, moniker string) error
	GetDefaultMoniker() ([]model.MonikerConfig, error)
}
