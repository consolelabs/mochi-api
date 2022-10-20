package model

import (
	"github.com/google/uuid"
)

type TradeItem struct {
	ID           uuid.UUID       `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	IsFrom       bool            `json:"is_from"`
	TokenAddress string          `json:"token_address"`
	TokenIds     JSONArrayString `json:"token_ids"`
	TradeOfferID string          `json:"trade_offer_id"`
}

func (TradeItem) TableName() string {
	return "trade_items"
}
