package user_tag

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(q GetOneQuery) (*model.UserTag, error)
	UpsertOne(tag model.UserTag) (*model.UserTag, error)
}
