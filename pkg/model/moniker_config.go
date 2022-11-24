package model

import (
	"time"

	"github.com/google/uuid"
)

type MonikerConfig struct {
	ID        uuid.UUID            `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Moniker   string               `json:"moniker"`
	Plural    string               `json:"plural"`
	GuildID   string               `json:"guild_id"`
	TokenID   uuid.UUID            `json:"token_id"`
	Token     *OffchainTipBotToken `json:"token"`
	Amount    float64              `json:"amount"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
