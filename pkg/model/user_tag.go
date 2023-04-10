package model

import (
	"time"
)

type UserTag struct {
	ID              int64     `json:"id"`
	UserId          string    `json:"user_id"`
	GuildId         string    `json:"guild_id"`
	MentionUsername bool      `json:"mention_username"`
	MentionRole     bool      `json:"mention_role"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
