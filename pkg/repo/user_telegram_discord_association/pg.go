package usertelegramdiscordassociation

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetOneByTelegramUsername(telegramUsername string) (*model.UserTelegramDiscordAssociation, error) {
	model := &model.UserTelegramDiscordAssociation{}
	return model, pg.db.Where("telegram_username = ?", telegramUsername).First(model).Error
}

func (pg *pg) Upsert(model *model.UserTelegramDiscordAssociation) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "discord_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"telegram_username"}),
	}).Create(model).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
