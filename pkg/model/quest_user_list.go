package model

import (
	"time"

	"github.com/google/uuid"
)

type QuestUserList struct {
	ID          uuid.UUID    `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	UserID      string       `json:"user_id"`
	QuestID     uuid.UUID    `json:"quest_id" swaggertype:"string"`
	Action      QuestAction  `json:"action"`
	Routine     QuestRoutine `json:"routine"`
	Current     int          `json:"current"`
	Target      int          `json:"target"`
	IsCompleted bool         `json:"is_completed"`
	IsClaimed   bool         `json:"is_claimed"`
	StartTime   time.Time    `json:"start_time"`
	EndTime     time.Time    `json:"end_time"`
}

func (QuestUserList) TableName() string {
	return "quests_user_list"
}
