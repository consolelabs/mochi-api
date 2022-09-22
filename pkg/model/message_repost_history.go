package model

import (
	"time"

	"github.com/google/uuid"
)

type MessageRepostHistory struct {
	ID              uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID         string        `json:"guild_id"`
	OriginMessageID string        `json:"origin_message_id"`
	OriginChannelID string        `json:"origin_channel_id"`
	RepostChannelID string        `json:"repost_channel_id"`
	RepostMessageID string        `json:"repost_message_id"`
	CreatedAt       *time.Time    `json:"created_at"`
}
