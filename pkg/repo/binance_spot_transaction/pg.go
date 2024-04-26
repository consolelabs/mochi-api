package binancespottransaction

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

func (pg *pg) Create(tx *model.BinanceSpotTransaction) error {
	return pg.db.Create(&tx).Error
}

func (pg *pg) List(q ListQuery) ([]model.BinanceSpotTransaction, error) {
	var txs []model.BinanceSpotTransaction
	db := pg.db
	if q.ProfileId != "" {
		db = db.Where("profile_id = ?", q.ProfileId)
	}
	return txs, pg.db.Find(&txs).Order("time DESC").Error
}
