package model

import (
	"time"

	"gorm.io/gorm"
)

type SaleBotTwitterConfig struct {
	ID             int                 `json:"id"`
	Address        string              `json:"address"`
	CollectionName string              `json:"collection_name"`
	Slug           string              `json:"slug"`
	ChainID        int                 `json:"chain_id"`
	MarketplaceID  int                 `json:"marketplace_id"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
	Marketplace    *SaleBotMarketplace `json:"marketplace" gorm:"foreignKey:MarketplaceID"`
}

func (s *SaleBotTwitterConfig) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().UTC()
	s.CreatedAt = now
	s.UpdatedAt = now
	return nil
}

func (s *SaleBotTwitterConfig) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now().UTC()
	return nil
}
