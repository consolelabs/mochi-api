package usertelegramdiscordassociation

import (
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

func (pg *pg) GetOneByTelegramID(telegramID string) (*model.UserTelegramDiscordAssociation, error) {
	model := &model.UserTelegramDiscordAssociation{}
	return model, pg.db.Where("telegram_id = ?", telegramID).First(model).Error
}

func (pg *pg) Upsert(model *model.UserTelegramDiscordAssociation) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "telegram_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"discord_id"}),
	}).Create(model).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
