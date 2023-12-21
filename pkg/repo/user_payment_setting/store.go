package userpaymentsetting

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	FirstOrCreate(model.UserPaymentSetting) (*model.UserPaymentSetting, error)
	Update(*model.UserPaymentSetting) error
}
