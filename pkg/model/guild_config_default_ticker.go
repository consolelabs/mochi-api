package model

type GuildConfigDefaultTicker struct {
	GuildID       string `json:"guild_id"`
	Query         string `json:"query"`
	DefaultTicker string `json:"default_ticker"`
}
