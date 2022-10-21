package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildBlacklistChannelRepostConfig struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   string        `json:"guild_id"`
	ChannelID string        `json:"channel_id"`
	CreatedAt time.Time     `json:"created_at"`
}
