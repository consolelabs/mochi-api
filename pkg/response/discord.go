package response

import "time"

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Bot           bool   `json:"bot"`
	Discriminator string `json:"discriminator"`
}

type DiscordGuildUser struct {
	User     *DiscordUser `json:"user"`
	GuildID  string       `json:"guild_id"`
	Nickname string       `json:"nickname"`
	Avatar   string       `json:"avatar"`
	Roles    []string     `json:"roles"`
	JoinedAt time.Time    `json:"joined_at"`
}

type DiscordGuildSticker struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Tags string `json:"tags"`
	Type int    `json:"type"`
}
