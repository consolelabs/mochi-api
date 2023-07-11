package model

import (
	"time"
)

type GuildUserActivityLog struct {
	ID           int       `json:"id"`
	GuildID      string    `json:"guild_id"`
	UserID       string    `json:"user_id"`
	ProfileID    string    `json:"profile_id"`
	ActivityName string    `json:"activity_name"`
	EarnedXP     int       `json:"earned_xp"`
	CreatedAt    time.Time `json:"created_at"`
}
