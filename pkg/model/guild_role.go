package model

type GuildRole struct {
	ID      int64  `json:"role_id"`
	Name    string `json:"name"`
	GuildID int64  `json:"guild_id"`
}
