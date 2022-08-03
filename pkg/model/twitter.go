package model

import "github.com/google/uuid"

type TwitterSalesMessage struct {
	ID                uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	IsNotifiedTwitter bool          `json:"is_notified_twitter"`
	TokenName         string        `json:"token_name"`
	CollectionName    string        `json:"collection_name"`
	Price             string        `json:"price"`
	SellerAddress     string        `json:"seller_address"`
	BuyerAddress      string        `json:"buyer_address"`
	Marketplace       string        `json:"marketplace"`
	MarketplaceURL    string        `json:"marketplace_url"`
	Image             string        `json:"image"`
}
