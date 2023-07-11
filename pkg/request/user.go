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
	ProfileID string    `json:"profile_id"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	CustomXP  int64     `json:"-"`
}

type SendUserXPRequest struct {
	Recipients []string `json:"recipients"`
	GuildID    string   `json:"guild_id"`
	Each       bool     `json:"each"`
	Amount     int      `json:"amount"`
}

type GetUserProfileRequest struct {
	ProfileID string `form:"profile_id" binding:"required"`
	GuildID   string `form:"guild_id" binding:"required"`
}

type GetTopUsersRequest struct {
	GuildID   string `form:"guild_id" binding:"required"`
	ProfileID string `form:"profile_id" binding:"required"`
	Page      int    `form:"page,default=0"`
	Limit     int    `form:"limit,default=10"`
	Query     string `form:"query"`
	Sort      string `form:"sort"`
}
