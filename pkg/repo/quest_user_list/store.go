package questuserlist

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	UpsertMany(list []model.QuestUserList) error
	List(ListQuery) ([]model.QuestUserList, error)
}
