package questuserreward

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateMany([]model.QuestUserReward) error
}
