package users

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(user *model.User) error
}
