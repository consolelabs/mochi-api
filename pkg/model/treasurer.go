package model

import "time"

type Treasurer struct {
	Id            int64     `json:"id"`
	VaultId       int64     `json:"vault_id"`
	GuildId       string    `json:"guild_id"`
	UserDiscordId string    `json:"user_discord_id"`
	Role          string    `json:"role"`
	Vault         Vault     `json:"vault" gorm:"foreignKey:VaultId;references:Id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
