package gitbook_click_collectors

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

func (pg *pg) UpsertOne(info model.GitbookClickCollector) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"count_clicks": info.CountClicks}),
	}).Create(&info).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) GetByCommandAndAction(cmd, action string) (model.GitbookClickCollector, error) {
	var info model.GitbookClickCollector
	return info, pg.db.Where("command = ? AND action = ?", cmd, action).First(&info).Error
}
