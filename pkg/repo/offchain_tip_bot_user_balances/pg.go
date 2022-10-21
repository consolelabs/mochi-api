package offchain_tip_bot_user_balances

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

func (pg *pg) GetUserBalances(userID string) ([]model.OffchainTipBotUserBalance, error) {
	var balance []model.OffchainTipBotUserBalance
	return balance, pg.db.Preload("Token").Find(&balance, "user_id = ?", userID).Error
}
