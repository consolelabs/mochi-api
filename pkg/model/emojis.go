package model

type ProductMetadataEmojis struct {
	ID         int     `json:"id"`
	Code       string  `json:"code"`
	DiscordId  *string `json:"discord_id"`
	TelegramId *string `json:"telegram_id"`
	TwitterId  *string `json:"twitter_id"`
}
