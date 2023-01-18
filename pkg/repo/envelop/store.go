package envelop

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(envelop *model.Envelop) error
	GetUserStreak(userID string) (model *model.UserEnvelopStreak, err error)
}
