package usertokenpricealert

import "github.com/defipod/mochi/pkg/model"

type UserTokenPriceAlertQuery struct {
	UserDiscordID string
	Symbol        string
	Price         float64
	Limit         int
	Offset        int
}

type Store interface {
	List(q UserTokenPriceAlertQuery) (items []model.UserTokenPriceAlert, total int64, err error)
	Create(item *model.UserTokenPriceAlert) error
	Delete(userID, symbol string, price float64) (rows int64, err error)
	Update(item *model.UserTokenPriceAlert) error
}
