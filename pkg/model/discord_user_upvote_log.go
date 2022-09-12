package model

import "time"

type DiscordUserUpvoteLog struct {
	DiscordID        string    `json:"discord_id"`
	Source           string    `json:"source"`
	LatestUpvoteTime time.Time `json:"last_upvote_time"`
	CreatedAt        time.Time `json:"created_at"`
}
