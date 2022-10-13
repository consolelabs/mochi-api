package questuserlog

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateOne(log *model.QuestUserLog) error
}
