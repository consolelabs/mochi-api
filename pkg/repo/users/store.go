package users

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(user *model.User) error
	UpsertOne(user *model.User) error

	GetLatestWalletNumber() int
	GetOne(discordID string) (*model.User, error)
	GetByDiscordIDs(discordIDs []string) ([]model.User, error)
}
