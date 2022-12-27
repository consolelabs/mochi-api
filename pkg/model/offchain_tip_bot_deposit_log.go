package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotDepositLog struct {
	ChainID     uuid.UUID `json:"chain_id"`
	TxHash      string    `json:"tx_hash"`
	TokenID     uuid.UUID `json:"token_id"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Amount      float64   `json:"amount"`
	AmountInUSD float64   `json:"amount_in_usd"`
	UserID      string    `json:"user_id"`
	BlockNumber int64     `json:"block_number"`
	SignedAt    time.Time `json:"signed_at"`
}

func (OffchainTipBotDepositLog) TableName() string {
	return "offchain_tip_bot_deposit_logs"
}
