package discordwalletverification

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(dicordID, guildID string) (*model.DiscordWalletVerification, error)
	UpsertOne(v model.DiscordWalletVerification) error
	GetByValidCode(code string) (*model.DiscordWalletVerification, error)
	DeleteByCode(code string) error
}
