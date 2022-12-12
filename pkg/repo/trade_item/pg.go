package trade_offer

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

func (pg *pg) GetOne(id string) (*model.TradeItem, error) {
	item := &model.TradeItem{}
	return item, pg.db.Table("trade_items").Where("id = ?", id).First(item).Error
}

func (pg *pg) UpsertOne(item *model.TradeItem) error {
	tx := pg.db.Begin()
	err := tx.Table("trade_items").Clauses(clause.OnConflict{DoNothing: true}).Create(&item).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
