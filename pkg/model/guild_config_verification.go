package model

import "time"

type GuildConfigVerification struct {
	GuildID         string    `json:"guild_id"`
	VerifyChannelID string    `json:"verify_channel_id"`
	VerifiedRoleID  string    `json:"verified_role_id"`
	Content         string    `json:"content"`
	EmbeddedMessage JSON      `json:"embedded_message"`
	CreatedAt       time.Time `json:"created_at"`
}
