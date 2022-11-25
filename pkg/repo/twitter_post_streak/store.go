package twitterpoststreak

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) (list []model.TwitterPostStreak, total int64, err error)
	UpsertOne(*model.TwitterPostStreak) error
}
