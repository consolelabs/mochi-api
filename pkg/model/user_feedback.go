package model

import (
	"time"

	"github.com/google/uuid"
)

type UserFeedback struct {
	ID          uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	DiscordID   string        `json:"discord_id"`
	MessageID   string        `json:"message_id"`
	Command     string        `json:"command"`
	Feedback    string        `json:"feedback"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	ConfirmedAt *time.Time    `json:"confirmed_at"`
	CompletedAt *time.Time    `json:"completed_at"`
}
