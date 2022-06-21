package nft_sales_tracker

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

func (pg *pg) FirstOrCreate(tracker *model.InsertNFTSalesTracker) error {
	return pg.db.Where("contract_address=? AND platform=? AND sales_config_id=?",tracker.ContractAddress,tracker.Platform,tracker.SalesConfigID).FirstOrCreate(tracker).Error
} 