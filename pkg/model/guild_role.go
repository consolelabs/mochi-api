package model

type GuildRole struct {
	ID      int64  `json:"role_id"`
	Name    string `json:"name"`
	GuildID string `json:"guild_id"`
}
