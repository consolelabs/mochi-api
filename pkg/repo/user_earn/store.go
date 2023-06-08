package userearn

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(*model.UserEarn) (*model.UserEarn, error)
	Delete(*model.UserEarn) (*model.UserEarn, error)
	GetByUserId(ListQuery) ([]model.UserEarn, int64, error)
}
