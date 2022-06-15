package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DefiType string

const (
	DEPOSIT  DefiType = "deposit"
	WITHDRAW DefiType = "withdraw"
	TIP      DefiType = "tip"
	AIRDROP  DefiType = "airdrop"
)

type DefiRole string

const (
	SENDER    DefiRole = "sender"
	RECIPIENT DefiRole = "recipient"
)

type UserDefiActivityLog struct {
	ID        uuid.NullUUID  `json:"id" gorm:"default:uuid_generate_v4()"`
	Type      DefiType       `json:"type"`
	Role      DefiRole       `json:"role"`
	TokenID   int            `json:"token_id"`
	UserID    string         `json:"user_id"`
	GuildID   JSONNullString `json:"guild_id"`
	Amount    float64        `json:"amount"`
	CreatedAt time.Time      `json:"created_at"`
}

func (log *UserDefiActivityLog) BeforeCreate(tx *gorm.DB) (err error) {
	if log.Role == SENDER && log.Amount > 0 {
		log.Amount = 0 - log.Amount
	}
	return nil
}
