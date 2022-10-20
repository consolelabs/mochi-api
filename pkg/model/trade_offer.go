package model

import (
	"time"

	"github.com/google/uuid"
)

type TradeOffer struct {
	ID          uuid.UUID   `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	FromAddress string      `json:"from_address"`
	ToAddress   string      `json:"to_address"`
	FromItems   []TradeItem `json:"from_items"`
	ToItems     []TradeItem `json:"to_items"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

func (TradeOffer) TableName() string {
	return "trade_offers"
}
