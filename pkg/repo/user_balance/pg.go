package userbalance

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetOne(userID string, tokenID int) (*model.UserBalance, error) {
	balance := &model.UserBalance{}
	return balance, pg.db.First(balance, "user_id = ? AND token_id = ?", userID, tokenID).Error
}

func (pg *pg) GetUserBalances(userID string) ([]model.UserBalance, error) {
	var balances []model.UserBalance
	return balances, pg.db.Preload("Token").Where("user_id = ?", userID).Find(&balances).Error
}
