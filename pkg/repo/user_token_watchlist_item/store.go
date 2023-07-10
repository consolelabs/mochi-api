package usertokenwatchlistitem

import "github.com/defipod/mochi/pkg/model"

type UserWatchlistQuery struct {
	ProfileID   string
	CoinGeckoID string
	Symbol      string
	Limit       int
	Offset      int
}

type CountQuery struct {
	CoingeckoId string
	Symbol      string
	Distinct    string
}

type Store interface {
	List(q UserWatchlistQuery) (items []model.UserTokenWatchlistItem, total int64, err error)
	Create(item *model.UserTokenWatchlistItem) error
	Delete(profileID, symbol string) (rows int64, err error)
	Count(CountQuery) (count int64, err error)
}
