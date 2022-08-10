package model

import "time"

type GuildConfigTwitterHashtag struct {
	UserID          string    `json:"user_id"`
	GuildID         string    `json:"guild_id"`
	ChannelID       string    `json:"channel_id"`
	RuleID          string    `json:"rule_id"`
	TwitterUsername string    `json:"twitter_username"`
	Hashtag         string    `json:"hashtag"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
