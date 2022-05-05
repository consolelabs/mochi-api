package model

import "time"

type DiscordWalletVerification struct {
	UserDiscordID string    `json:"user_discord_id"`
	GuildID       string    `json:"guild_id"`
	Code          string    `json:"code"`
	CreatedAt     time.Time `json:"created_at"`
}
