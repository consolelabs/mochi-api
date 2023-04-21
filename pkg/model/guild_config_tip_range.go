package model

import "time"

type GuildConfigTipRange struct {
	Id        int64     `json:"id"`
	GuildID   string    `json:"guild_id"`
	Min       *float64  `json:"min"` // fiat_usd
	Max       *float64  `json:"max"` // fiat_usd
	UpdatedAt time.Time `json:"updated_at"`
}
