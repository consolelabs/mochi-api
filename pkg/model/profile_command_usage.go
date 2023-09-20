package model

import "time"

type ProfileCommandUsage struct {
	ProfileId      string    `json:"profile_id"`
	UserPlatformId string    `json:"user_platform_id"`
	Command        string    `json:"command"`
	Params         string    `json:"params"`
	Platform       string    `json:"platform"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CommandUsageCounter struct {
	TotalUsage     int64
	ProfileId      string
	UserPlatformId string
	Platform       string
}
