package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type FriendTechKeyWatchlistItem struct {
	Id              int
	KeyAddress      string
	ProfileId       string
	IncreaseAlertAt int
	DecreaseAlertAt int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type FriendTechKeyPriceChangeAlertItem struct {
	Timestamp      time.Time       `json:"timestamp"`
	KeyAddressID   string          `json:"key_address_id"`
	KeyAddress     FriendTechKey   `json:"key_address"`
	Change         decimal.Decimal `json:"change"`
	CurrentPrice   decimal.Decimal `json:"current_price"`
	YesterdayPrice decimal.Decimal `json:"yesterday_price"`
}

type FriendTechKey struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Address         string    `json:"address"`
	TwitterUsername string    `json:"twitter_username"`
	TwitterPfpUrl   string    `json:"twitter_pfp_url"`
	ProfileChecked  bool      `json:"profile_checked"`
	Price           float64   `json:"price"`
	Supply          int       `json:"supply"`
	Holders         int       `json:"holders"`
}
