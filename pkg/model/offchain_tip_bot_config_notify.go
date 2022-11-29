package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotConfigNotify struct {
	ID        uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   string    `json:"guild_id"`
	ChannelID string    `json:"channel_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (OffchainTipBotConfigNotify) TableName() string {
	return "offchain_tip_bot_config_notify"
}
