package model

import (
	"github.com/google/uuid"
	"time"
)

type GuildConfigDefaultRole struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	RoleID    string        `json:"role_id"`
	GuildID   string        `json:"guild_id"`
	CreatedAt time.Time     `json:"created_at"`
}
