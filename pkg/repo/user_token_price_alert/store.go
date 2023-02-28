package usertokenpricealert

import "github.com/defipod/mochi/pkg/model"

type UserTokenPriceAlertQuery struct {
	UserDiscordID  string
	Symbol         string
	Value          float64
	PriceByPercent float64
	Limit          int
	Offset         int
}

type Store interface {
	List(q UserTokenPriceAlertQuery) (items []model.UserTokenPriceAlert, total int64, err error)
	GetOne(q UserTokenPriceAlertQuery) (item model.UserTokenPriceAlert, err error)
	Create(item *model.UserTokenPriceAlert) error
	Delete(userID, symbol string, value float64) (rows int64, err error)
	Update(item *model.UserTokenPriceAlert) error
	FetchListSymbol() ([]string, error)
}
