package model

import "time"

type DiscordUserUpvoteStreak struct {
	DiscordID      string    `json:"discord_id"`
	StreakCount    int       `json:"streak_count"`
	TotalCount     int       `json:"total_count"`
	LastStreakDate time.Time `json:"last_streak_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
