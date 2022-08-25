package model

import (
	"time"

	"github.com/google/uuid"
)

type NFTCollection struct {
	ID         uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Address    string        `json:"address"`
	Name       string        `json:"name"`
	Symbol     string        `json:"symbol"`
	ChainID    string        `json:"chain_id"`
	ERCFormat  string        `json:"erc_format"`
	IsVerified bool          `json:"is_verified"`
	CreatedAt  time.Time     `json:"created_at"`
	Image      string        `json:"image"`
	Author     string        `json:"author"`
}
type NewListedNFTCollection struct {
	ID         uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Address    string        `json:"address"`
	Name       string        `json:"name"`
	Symbol     string        `json:"symbol"`
	ChainID    string        `json:"chain_id"`
	Chain      string        `json:"chain"`
	ERCFormat  string        `json:"erc_format"`
	IsVerified bool          `json:"is_verified"`
	CreatedAt  time.Time     `json:"created_at"`
	Image      string        `json:"image"`
	Author     string        `json:"author"`
}

type NFTCollectionDetail struct {
	ID         uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	Address    string        `json:"address"`
	Name       string        `json:"name"`
	Symbol     string        `json:"symbol"`
	ChainID    string        `json:"chain_id"`
	Chain      *Chain        `json:"chain"`
	ERCFormat  string        `json:"erc_format"`
	IsVerified bool          `json:"is_verified"`
	CreatedAt  time.Time     `json:"created_at"`
	Image      string        `json:"image"`
	Author     string        `json:"author"`
}

type NFTCollectionConfig struct {
	NFTCollection
	TokenID string `json:"token_id"`
}
