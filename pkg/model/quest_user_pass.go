package model

import (
	"time"

	"github.com/google/uuid"
)

type QuestUserPass struct {
	UserID    string    `json:"user_id"`
	PassID    uuid.UUID `json:"pass_id" swaggertype:"string"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
	Active    bool      `json:"active"`
}
