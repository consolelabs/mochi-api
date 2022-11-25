package model

import (
	"time"

	"github.com/google/uuid"
)

type DiscordUserTokenAlert struct {
	ID                uuid.NullUUID      `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TokenID           string             `json:"token_id"`
	DiscordID         string             `json:"discord_id"`
	PriceSet          float64            `json:"price_set"`
	Trend             string             `json:"trend"`
	DeviceID          string             `json:"device_id"`
	IsEnable          bool               `json:"is_enable"`
	DiscordUserDevice *DiscordUserDevice `json:"device" gorm:"foreignkey:DeviceID"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

type UpsertDiscordUserTokenAlert struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TokenID   string        `json:"token_id"`
	DiscordID string        `json:"discord_id"`
	PriceSet  float64       `json:"price_set"`
	Trend     string        `json:"trend"`
	DeviceID  string        `json:"device_id"`
	IsEnable  bool          `json:"is_enable"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
