package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildConfigDefaultRole struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	RoleID    string        `json:"role_id"`
	GuildID   string        `json:"guild_id"`
	CreatedAt time.Time     `json:"created_at"`
}
