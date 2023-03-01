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
	GetById(ID int) (model.UserTokenPriceAlert, error)
	Create(item *model.UserTokenPriceAlert) (int, error)
	DeleteByID(alertID int) error
	Update(item *model.UserTokenPriceAlert) error
}
