package onchainassetavgcost

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Upsert(asset *model.OnchainAssetAvgCost) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "wallet_address"}, {Name: "token_address"}, {Name: "blockchain"}},
		DoUpdates: clause.AssignmentColumns([]string{"average_cost", "updated_at"}),
	}).Create(asset).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) UpsertMany(assets []model.OnchainAssetAvgCost) error {
	log := logger.NewLogrusLogger()
	tx := pg.db.Begin()
	for _, asset := range assets {
		err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "wallet_address"}, {Name: "token_address"}, {Name: "blockchain"}},
			DoUpdates: clause.AssignmentColumns([]string{"average_cost", "updated_at"}),
		}).Create(&asset).Error
		if err != nil {
			log.Error(err, "[onchainassetavgcost.UpsertMany] failed")
			continue
		}
	}
	return tx.Commit().Error
}

func (pg *pg) GetByWalletAddr(walletAddr string) ([]model.OnchainAssetAvgCost, error) {
	assets := make([]model.OnchainAssetAvgCost, 0)
	return assets, pg.db.Where("LOWER(wallet_address) = LOWER(?)", walletAddr).Find(&assets).Error
}
