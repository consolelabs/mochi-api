package model

import "time"

type Envelop struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Command   string    `json:"command"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
