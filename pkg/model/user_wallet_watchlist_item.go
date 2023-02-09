package model

import "time"

type UserWalletWatchlistItem struct {
	UserID    string    `json:"user_id"`
	Address   string    `json:"address"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"created_at"`
	NetWorth  float64   `json:"net_worth" gorm:"-"`
}
