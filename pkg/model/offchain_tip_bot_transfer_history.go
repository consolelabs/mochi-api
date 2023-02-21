package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotTransferHistory struct {
	ID         uuid.UUID  `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	SenderID   *string    `json:"sender_id"`
	ReceiverID string     `json:"receiver_id"`
	GuildID    string     `json:"guild_id"`
	LogID      string     `json:"log_id"`
	Status     string     `json:"status"`
	Amount     float64    `json:"amount"`
	Token      string     `json:"token"`
	Action     string     `json:"action"`
	ServiceFee float64    `json:"service_fee"`
	FeeAmount  float64    `json:"fee_amount"`
	TxHash     string     `json:"tx_hash"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-"`
}

func (OffchainTipBotTransferHistory) TableName() string {
	return "offchain_tip_bot_transfer_histories"
}
