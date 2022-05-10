package model

import "time"

type GuildUserActivityXP struct {
	ID         int       `json:"id"`
	GuildID    string    `json:"guild_id"`
	UserID     string    `json:"user_id"`
	ActivityID int       `json:"activity_id"`
	CreatedAt  time.Time `json:"created_at"`
}
