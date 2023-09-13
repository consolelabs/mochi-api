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
	KeyAddressID   string          `json:"key_address_id"`
	Change         decimal.Decimal `json:"change"`
	CurrentPrice   decimal.Decimal `json:"current_price"`
	YesterdayPrice decimal.Decimal `json:"yesterday_price"`
}
