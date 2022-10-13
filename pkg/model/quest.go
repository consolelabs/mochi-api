package model

import "github.com/google/uuid"

type QuestAction string

const (
	GM        QuestAction = "gm"
	VOTE      QuestAction = "vote"
	TRADE     QuestAction = "trade"
	GIFT      QuestAction = "gift"
	TICKER    QuestAction = "ticker"
	WATCHLIST QuestAction = "watchlist"
)

type QuestRoutine string

const (
	DAILY   QuestRoutine = "daily"
	WEEKLY  QuestRoutine = "weekly"
	MONTHLY QuestRoutine = "monthly"
	YEARLY  QuestRoutine = "yearly"
	ONCE    QuestRoutine = "once"
)

type Quest struct {
	ID        uuid.UUID    `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Title     string       `json:"title"`
	Action    QuestAction  `json:"action"`
	Frequency int          `json:"frequency"`
	Routine   QuestRoutine `json:"routine"`
}

func (Quest) TableName() string {
	return "quests"
}
