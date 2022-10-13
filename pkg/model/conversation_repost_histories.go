package model

import (
	"time"

	"github.com/google/uuid"
)

type ConversationRepostHistories struct {
	ID                   uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID              string        `json:"guild_id"`
	OriginChannelID      string        `json:"origin_channel_id"`
	OriginStartMessageID string        `json:"origin_start_message_id"`
	OriginStopMessageID  string        `json:"origin_stop_message_id"`
	RepostChannelID      string        `json:"repost_channel_id"`
	CreatedAt            time.Time     `json:"created_at"`
}
