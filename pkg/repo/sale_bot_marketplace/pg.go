package salebotmarketplace

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

func (pg *pg) List() ([]model.SaleBotMarketplace, error) {
	var marketplaces []model.SaleBotMarketplace
	return marketplaces, pg.db.Find(&marketplaces).Error
}

func (pg *pg) GetOne(name string) (*model.SaleBotMarketplace, error) {
	var marketplace model.SaleBotMarketplace
	return &marketplace, pg.db.First(&marketplace, "name ILIKE ?", name).Error
}
