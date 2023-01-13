package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildConfigSalesTracker struct {
	ID              uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string" swaggertype:"string"`
	GuildID         string        `json:"guild_id"`
	ChannelID       string        `json:"channel_id"`
	ContractAddress string        `json:"contract_address"`
	Chain           string        `json:"chain"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
