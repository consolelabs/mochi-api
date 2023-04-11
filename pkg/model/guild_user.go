package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildUser struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   string        `json:"guild_id" gorm:"unique_index:idx_guild_user_guild_id_user_id"`
	UserID    string        `json:"user_id" gorm:"uique_index:idx_guild_user_guild_id_user_id"`
	Nickname  string        `json:"nickname"`
	InvitedBy string        `json:"invited_by"`
	Avatar    string        `json:"avatar"`
	JoinedAt  time.Time     `json:"joined_at"`
	Roles     []byte        `json:"-" gorm:"roles type:jsonb;default:'[]';not null"`
	RoleSlice []string      `json:"roles" gorm:"-"`
}

type GuildUserRole struct {
	ID      uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID int64         `json:"guild_id"`
	UserID  int64         `json:"user_id"`
	RoleID  int64         `json:"role_id"`
}
