package model

import "time"

type UserWallet struct {
	UserDiscordID string         `json:"user_discord_id"`
	GuildID       string         `json:"guild_id"`
	Address       string         `json:"address"`
	ChainType     JSONNullString `json:"chain_type"`
	CreatedAt     time.Time      `json:"created_at"`
}
