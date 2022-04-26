package response

type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type DiscordGuildUser struct {
	User     *DiscordUser `json:"user"`
	GuildID  string       `json:"guild_id"`
	Nickname string       `json:"nickname"`
}
