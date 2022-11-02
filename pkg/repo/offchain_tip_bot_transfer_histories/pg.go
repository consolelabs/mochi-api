package offchain_tip_bot_transfer_histories

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

func (pg *pg) CreateTransferHistories(transferHistories []model.OffchainTipBotTransferHistory) ([]model.OffchainTipBotTransferHistory, error) {
	return transferHistories, pg.db.Create(transferHistories).Error
}

func (pg *pg) GetByUserDiscordId(userDiscordId string) (transferHistories []model.OffchainTipBotTransferHistory, err error) {
	return transferHistories, pg.db.Where("sender_id = ?", userDiscordId).Find(&transferHistories).Error
}
