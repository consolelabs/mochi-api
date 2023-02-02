package salebottwitterconfig

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

func (pg *pg) List(q ListQuery) ([]model.SaleBotTwitterConfig, error) {
	var configs []model.SaleBotTwitterConfig
	db := pg.db.Select("sale_bot_twitter_configs.*")
	if q.MarketplaceName != "" {
		db = db.
			Joins("JOIN sale_bot_marketplaces ON sale_bot_twitter_configs.marketplace_id = sale_bot_marketplaces.id").
			Where("sale_bot_marketplaces.name ILIKE ?", q.MarketplaceName)
	}
	return configs, db.Preload("Marketplace").Find(&configs).Error
}

func (pg *pg) Create(cfg *model.SaleBotTwitterConfig) error {
	return pg.db.Create(cfg).Error
}
