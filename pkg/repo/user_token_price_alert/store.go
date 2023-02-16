package usertokenpricealert

import "github.com/defipod/mochi/pkg/model"

type UserTokenPriceAlertQuery struct {
	UserID    string
	CoincapID string
	Limit     int
	Offset    int
}

type Store interface {
	List(q UserTokenPriceAlertQuery) (items []model.UserTokenPriceAlert, total int64, err error)
	Create(item *model.UserTokenPriceAlert) error
	Delete(userID, CoincapID string) (rows int64, err error)
	Update(item *model.UserTokenPriceAlert) error
}
