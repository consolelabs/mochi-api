package request

import (
	"time"
)

type CreateUserRequest struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	JoinDate  time.Time `json:"join_date"`
	GuildID   string    `json:"guild_id"`
	InvitedBy string    `json:"invited_by"`
}
