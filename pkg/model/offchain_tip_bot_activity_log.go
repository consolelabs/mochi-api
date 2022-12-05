package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type OffchainTipBotActivityLog struct {
	ID              uuid.UUID      `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	UserID          string         `json:"user_id"`
	GuildID         string         `json:"guild_id"`
	ChannelID       string         `json:"channel_id"`
	Action          *string        `json:"action"`
	Receiver        pq.StringArray `json:"receiver" gorm:"type:varchar(256)[];"`
	NumberReceivers int            `json:"number_receivers"`
	Duration        *int           `json:"duration"`
	TokenID         string         `json:"token_id"`
	Amount          float64        `json:"amount"`
	FullCommand     *string        `json:"full_command"`
	Status          string         `json:"status"`
	FailReason      string         `json:"fail_reason"`
	ServiceFee      float64        `json:"service_fee"`
	FeeAmount       float64        `json:"fee_amount"`
	Image           string         `json:"image"`
	Message         string         `json:"message"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       *time.Time     `json:"deleted_at"`
}

func (OffchainTipBotActivityLog) TableName() string {
	return "offchain_tip_bot_activity_logs"
}
