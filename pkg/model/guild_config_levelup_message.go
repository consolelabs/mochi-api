package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildConfigLevelupMessage struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   string        `json:"guild_id"`
	ImageURL  string        `json:"image_url"`
	Message   string        `json:"message"`
	ChannelID string        `json:"channel_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
