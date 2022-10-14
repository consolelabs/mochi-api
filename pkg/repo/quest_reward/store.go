package questreward

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type Store interface {
	GetQuestRewards(questID []uuid.UUID) ([]model.QuestReward, error)
}
