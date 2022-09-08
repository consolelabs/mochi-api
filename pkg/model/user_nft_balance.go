package model

import "github.com/google/uuid"

type UserNFTBalance struct {
	UserAddress     string         `json:"user_address"`
	ChainType       JSONNullString `json:"chain_type"`
	NFTCollectionID uuid.NullUUID  `json:"nft_collection_id" swaggertype:"string"`
	TokenID         string         `json:"token_id,omitempty"`
	Balance         int            `json:"balance"`
}
