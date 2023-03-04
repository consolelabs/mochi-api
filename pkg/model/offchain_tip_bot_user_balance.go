package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotUserBalance struct {
	ID            uuid.UUID            `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	UserID        string               `json:"user_id"`
	TokenID       uuid.UUID            `json:"token_id"`
	Token         *OffchainTipBotToken `json:"token"`
	Amount        float64              `json:"amount"`
	ChangedAmount float64              `json:"-" gorm:"-"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	DeletedAt     *time.Time           `json:"-"`
}

func (OffchainTipBotUserBalance) TableName() string {
	return "offchain_tip_bot_user_balances"
}
