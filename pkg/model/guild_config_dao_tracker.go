package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildConfigDaoTracker struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   string        `json:"guild_id"`
	Space     string        `json:"space"`
	Source    string        `json:"source"`
	ChannelID string        `json:"channel_id"`
	UpdatedAt time.Time     `json:"updated_at"`
	CreatedAt time.Time     `json:"created_at"`
}
