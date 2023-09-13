package model

import "time"

type FriendTechKeyWatchlistItem struct {
	Id              int
	KeyAddress      string
	ProfileId       string
	IncreaseAlertAt int
	DecreaseAlertAt int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
