package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildConfigDefaultCurrency struct {
	ID            uuid.NullUUID       `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID       string              `json:"guild_id"`
	TipBotTokenID string              `json:"tip_bot_token_id"`
	TipBotToken   OffchainTipBotToken `json:"tip_bot_token" gorm:"foreignKey:TipBotTokenID"`
	UpdatedAt     time.Time           `json:"updated_at"`
	CreatedAt     time.Time           `json:"created_at"`
}

type UpsertGuildConfigDefaultCurrency struct {
	ID            uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID       string        `json:"guild_id"`
	TipBotTokenID string        `json:"tip_bot_token_id"`
	UpdatedAt     time.Time     `json:"updated_at"`
	CreatedAt     time.Time     `json:"created_at"`
}
