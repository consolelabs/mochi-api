package message_reaction

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

func (pg *pg) Create(record model.MessageReaction) error {
	return pg.db.Create(&record).Error
}

func (pg *pg) GetByMessageID(messageID string) ([]model.MessageReaction, error) {
	msgReaction := []model.MessageReaction{}
	return msgReaction, pg.db.Where("message_id = ?", messageID).Find(&msgReaction).Error
}

func (pg *pg) Delete(messageID string, userID string, reaction string) error {
	return pg.db.Where("message_id = ? AND user_id = ? AND reaction = ?", messageID, userID, reaction).Delete(&model.MessageReaction{}).Error
}

func (pg *pg) DeleteByMessageID(messageID string) error {
	return pg.db.Where("message_id = ?", messageID).Delete(&model.MessageReaction{}).Error
}
