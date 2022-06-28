package model

import "github.com/google/uuid"

type NFTCollection struct {
	ID         uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Address    string        `json:"address"`
	Name       string        `json:"name"`
	Symbol     string        `json:"symbol"`
	ChainID    string        `json:"chain_id"`
	ERCFormat  string        `json:"erc_format"`
	IsVerified bool          `json:"is_verified"`
}

type NFTCollectionConfig struct {
	NFTCollection
	TokenID string `json:"token_id"`
}
