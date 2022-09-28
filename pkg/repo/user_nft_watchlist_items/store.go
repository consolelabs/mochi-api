package usernftwatchlistitem

import "github.com/defipod/mochi/pkg/model"

type UserNftWatchlistQuery struct {
	UserID            string
	CollectionAddress string
	ChainID           string
	Symbol            string
	Limit             int
	Offset            int
}
type Store interface {
	List(q UserNftWatchlistQuery) ([]model.UserNftWatchlistItem, int64, error)
	Create(item *model.UserNftWatchlistItem) error
	Delete(userID, symbol string) (rows int64, err error)
}
