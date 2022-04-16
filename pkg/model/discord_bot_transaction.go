package model

import (
	"time"
)

type DiscordBotTransaction struct {
	TxHash        string    `json:"tx_hash"`
	FromDiscordID string    `json:"from_discord_id"`
	ToDiscordID   string    `json:"to_discord_id"`
	ToAddress     string    `json:"to_address"`
	Amount        float64   `json:"amount"`
	Reason        string    `json:"reason"`
	Type          string    `json:"type"`
	GuildID       string    `json:"guild_id"`
	ChannelID     string    `json:"channel_id"`
	TokenID       int       `json:"token_id"`
	CreatedAt     time.Time `json:"created_at"`
}
