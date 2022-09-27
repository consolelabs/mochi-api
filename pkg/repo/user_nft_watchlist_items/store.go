package usernftwatchlistitem

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	// List(userID string) (items []model.UserNftWatchlistItem, total int64, err error)
	Create(item *model.UserNftWatchlistItem) error
	Delete(userID, symbol string) (rows int64, err error)
}
