package model

import (
	"time"

	"github.com/google/uuid"
)

type InviteHistory struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID   int64         `json:"guild_id"`
	UserID    int64         `json:"user_id"`
	InvitedBy int64         `json:"invited_by"`
	CreatedAt *time.Time    `json:"created_at"`
	Metadata  JSON          `json:"metadata"`
}
