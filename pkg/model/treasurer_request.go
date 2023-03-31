package model

import "time"

type TreasurerRequest struct {
	Id            int64      `json:"id"`
	VaultId       int64      `json:"vault_id"`
	GuildId       string     `json:"guild_id"`
	UserDiscordId string     `json:"user_discord_id"`
	Message       string     `json:"message"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}
