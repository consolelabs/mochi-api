package userwatchlistitem

import "github.com/defipod/mochi/pkg/model"

type UserWatchlistQuery struct {
	UserID      string
	CoinGeckoID string
	Symbol      string
	Limit       int
	Offset      int
}

type Store interface {
	List(q UserWatchlistQuery) (items []model.UserWatchlistItem, total int64, err error)
	Create(item *model.UserWatchlistItem) error
	Delete(userID, symbol string) error
}
