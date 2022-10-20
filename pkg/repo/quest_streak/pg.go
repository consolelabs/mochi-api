package queststreak

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
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
	if q.StreakCount != 0 {
		db = db.Where("streak_from <= ? AND (streak_to IS NULL OR streak_to >= ?)", q.StreakCount, q.StreakCount)
	}
	return list, db.Find(&list).Error
}
