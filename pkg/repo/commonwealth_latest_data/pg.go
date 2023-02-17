package commonwealth_lastest_data

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

func (pg *pg) UpsertOne(model model.CommonwealthLatestData) error {
	tx := pg.db.Begin()

	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "community_id"}},
		UpdateAll: true,
	}).Create(&model).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) GetAll() (models []model.CommonwealthLatestData, err error) {
	return models, pg.db.Find(&models).Error
}
