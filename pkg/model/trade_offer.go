package model

import (
	"time"

	"github.com/google/uuid"
)

type TradeOffer struct {
	ID           uuid.UUID   `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	OwnerAddress string      `json:"owner_address"`
	HaveItems    []TradeItem `json:"have_items"`
	WantItems    []TradeItem `json:"want_items"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func (TradeOffer) TableName() string {
	return "trade_offers"
}
