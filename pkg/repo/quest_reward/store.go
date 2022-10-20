package questreward

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	List(q ListQuery) ([]model.QuestReward, error)
}
