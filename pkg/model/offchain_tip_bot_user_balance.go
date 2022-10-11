package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotUserBalance struct {
	ID        uuid.UUID  `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	UserID    string     `json:"user_id"`
	TokenID   string     `json:"token_id"`
	Amount    *float64   `json:"amount"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (OffchainTipBotUserBalance) TableName() string {
	return "offchain_tip_bot_user_balances"
}
