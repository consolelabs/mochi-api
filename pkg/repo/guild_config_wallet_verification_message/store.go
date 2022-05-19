package guildconfigwalletverificationmessage

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(guildID string) (*model.GuildConfigWalletVerificationMessage, error)
	UpsertOne(gcv model.GuildConfigWalletVerificationMessage) error
	DeleteOne(guildID string) error
}
