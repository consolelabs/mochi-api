package model

import "time"

type UserNftWatchlistItem struct {
	ProfileID         string     `json:"profile_id"`
	Symbol            string     `json:"symbol"`
	CollectionAddress string     `json:"collection_address"`
	ChainId           int64      `json:"chain_id"`
	CreatedAt         *time.Time `json:"created_at"`
}
