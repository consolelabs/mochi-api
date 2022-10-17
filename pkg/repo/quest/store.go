package quest

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(q ListQuery) ([]model.Quest, error)
	GetAvailableRoutines() ([]model.QuestRoutine, error)
}
