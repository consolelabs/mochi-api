package users

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(user *model.User) error
	GetOne(discordID string) (*model.User, error)
	GetByDiscordIDs(discordIDs []string) ([]model.User, error)
	UpdateNrOfJoin(discordId string, nrOfJoin int64) error
}
