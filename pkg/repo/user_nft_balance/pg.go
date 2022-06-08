package user_nft_balance

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) Upsert(balance model.UserNFTBalance) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_address"},
			{Name: "chain_type"},
			{Name: "nft_collection_id"},
			{Name: "token_id"},
		},
		UpdateAll: true,
	}).Create(&balance).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
