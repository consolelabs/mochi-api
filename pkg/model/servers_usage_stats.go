package model

type UsageStat struct {
	ID      int    `json:"id"`
	UserID  string `json:"user_id"`
	GuildID string `json:"guild_id"`
	Command string `json:"command"`
	Args    string `json:"args"`
}
