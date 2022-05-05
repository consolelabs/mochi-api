package userwallet

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOneByDiscordIDAndGuildID(discordID, guildID string) (*model.UserWallet, error)
	GetOneByGuildIDAndAddress(guildID, address string) (*model.UserWallet, error)
	UpsertOne(uw model.UserWallet) error
}
