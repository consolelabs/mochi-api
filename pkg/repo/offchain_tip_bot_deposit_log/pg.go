package offchain_tip_bot_deposit_log

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

func (pg *pg) GetOne(chainID, txHash string) (*model.OffchainTipBotDepositLog, error) {
	var rs model.OffchainTipBotDepositLog
	return &rs, pg.db.Where("chain_id::TEXT = ? AND tx_hash ILIKE ?", chainID, txHash).First(&rs).Error
}

func (pg *pg) CreateMany(list []model.OffchainTipBotDepositLog) error {
	tx := pg.db.Begin()
	for _, item := range list {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
