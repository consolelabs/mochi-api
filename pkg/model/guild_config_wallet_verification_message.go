package model

import "time"

type GuildConfigWalletVerificationMessage struct {
	GuildID          string    `json:"guild_id"`
	VerifyChannelID  string    `json:"verify_channel_id"`
	Content          string    `json:"content"`
	EmbeddedMessage  JSON      `json:"embedded_message"`
	CreatedAt        time.Time `json:"created_at"`
	DiscordMessageID string    `json:"discord_message_id"`
}
