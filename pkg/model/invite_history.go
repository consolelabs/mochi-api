package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	INVITE_TYPE_NORMAL = "normal"
	INVITE_TYPE_FAKE   = "fake"
	INVITE_TYPE_LEFT   = "left"
)

type InviteHistory struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   string        `json:"guild_id"`
	UserID    string        `json:"user_id"`
	InvitedBy string        `json:"invited_by"`
	CreatedAt *time.Time    `json:"created_at"`
	Metadata  JSON          `json:"metadata"`
	Type      string        `json:"type"`
}
