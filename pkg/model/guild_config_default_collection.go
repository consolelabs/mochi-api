package model

import "time"

type GuildConfigDefaultCollection struct {
	GuildID   string    `json:"guild_id"`
	Symbol    string    `json:"symbol"`
	Address   string    `json:"address"`
	ChainID   string    `json:"chain_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
