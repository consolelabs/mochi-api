package model

import (
	"time"

	"gorm.io/gorm"
)

type TwitterPostStreak struct {
	GuildID        string    `json:"guild_id"`
	TwitterID      string    `json:"twitter_id"`
	TwitterHandle  string    `json:"twitter_handle"`
	StreakCount    int       `json:"streak_count"`
	TotalCount     int       `json:"total_count"`
	LastStreakDate time.Time `json:"last_streak_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (s *TwitterPostStreak) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().UTC()
	s.CreatedAt = now
	s.UpdatedAt = now
	return nil
}

func (s *TwitterPostStreak) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now().UTC()
	return nil
}
