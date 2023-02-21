package model

import "time"

type UserWalletWatchlistItem struct {
	UserID      string    `json:"user_id"`
	Address     string    `json:"address"`
	Alias       string    `json:"alias"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	NetWorth    float64   `json:"net_worth" gorm:"-"`
	FetchedData bool      `json:"fetched_data" gorm:"-"`
}
