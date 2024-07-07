package model

import (
	"time"

	"github.com/google/uuid"
)

type ErcFormat string

const (
	ErcFormat721  = "ERC721"
	ErcFormat1155 = "ERC721"
)

type NFTCollection struct {
	ID          uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Address     string        `json:"address"`
	Name        string        `json:"name"`
	Symbol      string        `json:"symbol"`
	ChainID     string        `json:"chain_id"`
	ERCFormat   string        `json:"erc_format"`
	IsVerified  bool          `json:"is_verified"`
	CreatedAt   time.Time     `json:"created_at"`
	Image       string        `json:"image"`
	Author      string        `json:"author"`
	TotalSupply int           `json:"total_supply"`
}

type NewListedNFTCollection struct {
	ID         uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
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
	ID         uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
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
	ID        uuid.NullUUID `json:"id"`
	ERCFormat string        `json:"erc_format"`
	Address   string        `json:"address"`
	Name      string        `json:"name"`
	ChainID   string        `json:"chain_id"`
}
