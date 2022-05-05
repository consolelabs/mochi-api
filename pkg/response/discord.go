package response

type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Bot      bool   `json:"bot"`
}

type DiscordGuildUser struct {
	User     *DiscordUser `json:"user"`
	GuildID  string       `json:"guild_id"`
	Nickname string       `json:"nickname"`
}

type DiscordGuildSticker struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Tags string `json:"tags"`
	Type int    `json:"type"`
}
