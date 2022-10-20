package queststreak

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(q ListQuery) ([]model.QuestStreak, error)
}
