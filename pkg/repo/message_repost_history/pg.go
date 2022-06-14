package message_repost_history

import (
	"fmt"
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetByMessageID(guildID, messageID string) (model.MessageRepostHistory, error) {
	var config model.MessageRepostHistory
	err := pg.db.Model(model.MessageRepostHistory{}).Where("guild_id = ? AND origin_message_id = ?", guildID, messageID).First(&config).Error
	if err != nil {
		return config, fmt.Errorf("failed to get message repost history: %w", err)
	}
	return config, nil
}

func (pg *pg) CreateIfNotExist(record model.MessageRepostHistory) error {
	//return pg.db.Create(&config).Error
	tx := pg.db.Begin()

	err := tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "origin_message_id"}},
		DoNothing: true,
	}).Create(&record).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
