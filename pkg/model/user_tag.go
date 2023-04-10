package model

import (
	"time"
)

type UserTag struct {
	ID        int64     `json:"id"`
	UserId    string    `json:"user_id"`
	GuildId   string    `json:"guild_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
