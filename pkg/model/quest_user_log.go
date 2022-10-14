package model

import (
	"time"

	"github.com/google/uuid"
)

type QuestUserLog struct {
	ID        uuid.UUID   `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   *string     `json:"guild_id"`
	UserID    string      `json:"user_id"`
	QuestID   uuid.UUID   `json:"quest_id"`
	Action    QuestAction `json:"action"`
	Target    int         `json:"target"`
	CreatedAt time.Time   `json:"created_at"`
}

func (QuestUserLog) TableName() string {
	return "quests_user_logs"
}
