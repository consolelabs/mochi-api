package offchaintipbotuserbalancesnapshot

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

func (pg *pg) CreateBatch(list []model.OffchainTipBotUserBalanceSnapshot) error {
	tx := pg.db.Begin()
	for _, item := range list {
		err := tx.Create(&item).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
