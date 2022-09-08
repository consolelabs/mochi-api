package model

import (
	"github.com/google/uuid"
)

type NFTSalesTracker struct {
	ID                      uuid.NullUUID           `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	ContractAddress         string                  `json:"contract_address"`
	Platform                string                  `json:"platform"`
	SalesConfigID           string                  `json:"sales_config_id"`
	GuildConfigSalesTracker GuildConfigSalesTracker `json:"guild_config_sales_tracker" gorm:"foreignKey:sales_config_id"`
}
type InsertNFTSalesTracker struct {
	ID              uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	ContractAddress string        `json:"contract_address"`
	Platform        string        `json:"platform"`
	SalesConfigID   string        `json:"sales_config_id" gorm:"foreignKey:ID"`
}

func (InsertNFTSalesTracker) TableName() string {
	return "nft_sales_trackers"
}
