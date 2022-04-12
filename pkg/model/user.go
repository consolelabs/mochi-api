package model

import (
	"time"
)

type User struct {
	ID                     int64          `json:"id" gorm:"default:uuid_generate_v4()"`
	Username               string         `json:"username"`
	Nickname               JSONNullString `json:"nickname"`
	JoinDate               *time.Time     `json:"join_date"`
	InDiscordWalletAddress JSONNullString `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  JSONNullInt64  `json:"in_discord_wallet_number"`
}
