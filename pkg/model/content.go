package model

import "time"

type Content struct {
	Id          int64     `json:"id"`
	Type        string    `json:"type"`
	Description []byte    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
