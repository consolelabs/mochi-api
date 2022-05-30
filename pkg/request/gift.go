package request

type GiftXpRequest struct {
	AdminDiscordId string `json:"admin_discord_id"`
	UserDiscordId string `json:"user_discord_id"`
	GuildId string `json:"guild_id"`
	XpAmount string `json:"xp_amount"`
}