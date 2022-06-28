package request

import "time"

type CreateUserRequest struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	GuildID   string `json:"guild_id"`
	InvitedBy string `json:"invited_by"`
}

type HandleUserActivityRequest struct {
	GuildID   string    `json:"guild_id"`
	ChannelID string    `json:"channel_id"`
	UserID    string    `json:"user_id"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	CustomXP  int64     `json:"-"`
}
