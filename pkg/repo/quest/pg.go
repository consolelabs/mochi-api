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
	if q.Routine != nil {
		db = db.Where("routine::TEXT = ?", *q.Routine)
	}
	return quests, db.Find(&quests).Error
}
