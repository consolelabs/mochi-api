package model

import (
	"time"
)

type NftSoulbound struct {
	ID                int64     `json:"id"`
	CollectionAddress string    `json:"collection_address"`
	TraitType         string    `json:"trait_type"`
	Value             string    `json:"value"`
	TotalSoulbound    int64     `json:"total_soulbound"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAT         time.Time `json:"updated_at"`
}
