package model

type UserTelegramDiscordAssociation struct {
	TelegramUsername string `json:"telegram_username"`
	DiscordID        string `json:"discord_id"`
}
