package model

type Activity struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	XP           int    `json:"xp"`
	GuildDefault bool   `json:"guild_default"`
}
