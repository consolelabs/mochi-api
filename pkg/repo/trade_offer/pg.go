package trade_offer

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

func (pg *pg) GetOne(id string) (*model.TradeOffer, error) {
	offer := &model.TradeOffer{}
	return offer, pg.db.Table("trade_offers").Where("id = ?", id).Preload("FromItems").Preload("ToItems").First(offer).Error
}

func (pg *pg) Create(offer *model.TradeOffer) (*model.TradeOffer, error) {
	return offer, pg.db.Table("trade_offers").Create(offer).Error
}
