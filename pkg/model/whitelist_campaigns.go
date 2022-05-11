package model

import "time"

type WhitelistCampaign struct {
	ID        int64      `json:"role_id"`
	Name      string     `json:"name"`
	GuildID   string     `json:"guild_id"`
	CreatedAt time.Time `json:"created_at"`
}
