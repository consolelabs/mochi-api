package queststreak

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) ([]model.QuestStreak, error) {
	var list []model.QuestStreak
	db := pg.db
	if q.Action != "" {
		db = db.Where("action = ?", q.Action)
	}
	db = db.Where(
		pg.db.Where("streak_from <= ? AND (streak_to IS NULL OR streak_to >= ?)", q.StreakCount, q.StreakCount),
	).Or(
		pg.db.Where("? > streak_from AND ? > streak_to", q.StreakCount, q.StreakCount),
	)
	if q.Sort != "" {
		db = db.Order(q.Sort)
	}
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return list, db.Find(&list).Error
}
