package request

type GiftXPRequest struct {
	AdminDiscordID string `json:"admin_discord_id"`
	UserDiscordID  string `json:"user_discord_id"`
	GuildID        string `json:"guild_id"`
	XPAmount       uint64 `json:"xp_amount"`
	ChannelID      string `json:"channel_id"`
}
