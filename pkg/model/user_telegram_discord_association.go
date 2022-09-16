package model

type UserTelegramDiscordAssociation struct {
	TelegramID int64  `json:"telegram_id"`
	DiscordID  string `json:"discord_id"`
}
