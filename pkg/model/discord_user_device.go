package model

import (
	"time"
)

type DiscordUserDevice struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	IosNotiToken string    `json:"ios_noti_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
