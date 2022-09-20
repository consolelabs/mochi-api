package request

type LinkUserTelegramWithDiscordRequest struct {
	TelegramUsername string `json:"telegram_username"`
	DiscordID        string `json:"discord_id"`
}
