package model

import (
	"time"

	"github.com/google/uuid"
)

type InviteHistory struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID   int64         `json:"guild_id"`
	UserID    int64         `json:"user_id"`
	CreatedAt *time.Time    `json:"created_at"`
	Metadata  interface{}   `json:"metadata"`
}
