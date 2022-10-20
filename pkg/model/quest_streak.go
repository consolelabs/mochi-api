package model

import "github.com/google/uuid"

type QuestStreak struct {
	ID         uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Title      string    `json:"title"`
	Action     string    `json:"action"`
	StreakFrom int       `json:"streak_from"`
	StreakTo   *int      `json:"streak_to"`
	Multiplier float64   `json:"multiplier"`
}

func (QuestStreak) TableName() string {
	return "quests_streak"
}
