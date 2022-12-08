package model

import "time"

type NftAddRequestHistory struct {
	Address   string    `json:"address"`
	ChainID   int64     `json:"chain_id"`
	GuildID   string    `json:"guild_id"`
	ChannelID string    `json:"channel_id"`
	MessageID string    `json:"message_id"`
	CreatedAt time.Time `json:"created_at"`
}
