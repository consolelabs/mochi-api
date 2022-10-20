package questreward

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

func (pg *pg) List(q ListQuery) ([]model.QuestReward, error) {
	var rewards []model.QuestReward
	db := pg.db
	if q.QuestIDs != nil {
		db = db.Where("quest_id IN ?", q.QuestIDs)
	}
	return rewards, db.Preload("Quest").Preload("RewardType").Find(&rewards).Error
}
