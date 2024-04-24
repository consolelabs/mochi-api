package binancetracking

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

func (pg *pg) FirstOrCreate(binanceTracking *model.BinanceTracking) (*model.BinanceTracking, error) {
	return binanceTracking, pg.db.Where("profile_id = ?", binanceTracking.ProfileId).FirstOrCreate(&binanceTracking).Error
}
