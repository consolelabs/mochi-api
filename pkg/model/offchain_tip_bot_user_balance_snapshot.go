package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotUserBalanceSnapshot struct {
	UserID        string
	TokenID       uuid.UUID
	TokenSymbol   string
	Action        string
	ChangedAmount float64
	Amount        float64
	CreatedAt     time.Time
}

func (OffchainTipBotUserBalanceSnapshot) TableName() string {
	return "offchain_tip_bot_user_balance_snapshots"
}
