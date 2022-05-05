package guildconfigverification

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(guildID string) (*model.GuildConfigVerification, error)
	UpsertOne(gcv model.GuildConfigVerification) error
}
