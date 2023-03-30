package model

import "time"

type VaultConfig struct {
	Id        int64     `json:"id"`
	GuildId   string    `json:"guild_id"`
	ChannelId string    `json:"channel_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
