package model

import "time"

type GuildConfigTwitterBlacklist struct {
	ID              int       `json:"-"`
	GuildID         string    `json:"guild_id"`
	TwitterUsername string    `json:"twitter_username"`
	TwitterID       string    `json:"twitter_id"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
}

func (GuildConfigTwitterBlacklist) TableName() string {
	return "guild_config_twitter_blacklist"
}
