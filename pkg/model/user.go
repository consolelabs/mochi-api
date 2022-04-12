package model

import (
	"time"
)

type User struct {
	ID                     int64          `json:"id" gorm:"default:uuid_generate_v4()"`
	ReferralCode           string         `json:"referral_code" gorm:"default:substring(md5(random()::text), 0, 9)"`
	InvitedBy              JSONNullInt64  `json:"invited_by"`
	Username               string         `json:"username"`
	Nickname               string         `json:"nickname"`
	JoinDate               *time.Time     `json:"join_date"`
	InDiscordWalletAddress JSONNullString `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  JSONNullInt64  `json:"in_discord_wallet_number"`
}
