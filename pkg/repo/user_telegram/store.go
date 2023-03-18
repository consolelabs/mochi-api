package usertelegram

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByUsername(username string) (*model.UserTelegram, error)
}
