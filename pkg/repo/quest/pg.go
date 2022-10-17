package quest

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

func (pg *pg) List(q ListQuery) ([]model.Quest, error) {
	var quests []model.Quest
	db := pg.db
	if q.ID != nil {
		db = db.Where("id = ?", q.ID)
	}
	if q.Action != nil {
		db = db.Where("action::TEXT = ?", *q.Action)
	}
	if q.NotAction != nil {
		db = db.Where("action::TEXT != ?", *q.NotAction)
	}
	if q.Routine != nil {
		db = db.Where("routine::TEXT = ?", *q.Routine)
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
