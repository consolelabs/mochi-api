package auto_action_history

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

func (pg *pg) Update(id string, field string, value interface{}) error {
	return pg.db.Model(&model.AutoActionHistory{}).Where("id = ?", id).Update(field, value).Error
}

func (pq *pg) CountByTriggerActionUserMessage(triggerId int64, actionId int64, userId string, messageId string) (int64, error) {
	var count int64
	err := pq.db.Model(&model.AutoActionHistory{}).Where("trigger_id = ? AND action_id = ? AND user_discord_id = ? AND message_id = ?", triggerId, actionId, userId, messageId).Count(&count).Error
	return count, err
}

func (pg *pg) FirstOrCreate(actionHistory *model.AutoActionHistory) error {
	return pg.db.Where("trigger_id = ? AND action_id = ? AND user_id = ? AND message_id = ?", actionHistory.TriggerId, actionHistory.ActionId, actionHistory.UserId, actionHistory.MessageId).FirstOrCreate(actionHistory).Error
}

func (pg *pg) GetById(id string) (*model.AutoActionHistory, error) {
	var result *model.AutoActionHistory
	db := pg.db.Where("id = ?", id)
	return result, db.First(&result).Error
}

func (pg *pg) Create(actionHistory *model.AutoActionHistory) error {
	return pg.db.Create(actionHistory).Error
}
