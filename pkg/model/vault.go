package model

import "time"

type Vault struct {
	Id            int64     `json:"id"`
	GuildId       string    `json:"guild_id"`
	Name          string    `json:"name"`
	Threshold     string    `json:"threshold"`
	WalletAddress string    `json:"wallet_address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
