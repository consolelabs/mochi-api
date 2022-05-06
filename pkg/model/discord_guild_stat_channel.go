package model

import (
	"time"

	"github.com/google/uuid"
)

type DiscordGuildStatChannel struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID   string        `json:"guild_id"`
	ChannelID string        `json:"channel_id"`
	CountType string        `json:"count_type"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
