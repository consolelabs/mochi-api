package quest

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

func (pg *pg) List(q ListQuery) ([]model.Quest, error) {
	var quests []model.Quest
	db := pg.db
	if q.ID != nil {
		db = db.Where("id = ?", q.ID)
	}
	if q.Action != "" {
		db = db.Where("action::TEXT = ?", q.Action)
	}
	if q.NotActions != nil && len(q.NotActions) > 0 {
		db = db.Where("action::TEXT NOT IN ?", q.NotActions)
	}
	if q.Routine != "" {
		db = db.Where("routine::TEXT = ?", q.Routine)
	}
	if q.Sort == "" {
		q.Sort = "title"
	}
	db = db.Order(q.Sort).Preload("Rewards").Preload("Rewards.RewardType")
	return quests, db.Find(&quests).Error
}

func (pg *pg) GetAvailableRoutines() ([]model.QuestRoutine, error) {
	var routines []model.QuestRoutine
	return routines, pg.db.Table("quests").Distinct().Pluck("routine", &routines).Error
}
