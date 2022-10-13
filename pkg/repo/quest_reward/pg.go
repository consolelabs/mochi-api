package questreward

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetQuestRewards(questID []uuid.UUID) ([]model.QuestReward, error) {
	var rewards []model.QuestReward
	return rewards, pg.db.Where("quest_id IN ?", questID).Preload("Quest").Preload("RewardType").Find(&rewards).Error
}
