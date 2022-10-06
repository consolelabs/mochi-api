package request

type LinkUserTelegramWithDiscordRequest struct {
	TelegramUsername string `json:"telegram_username" binding:"required"`
	DiscordID        string `json:"discord_id" binding:"required"`
}
