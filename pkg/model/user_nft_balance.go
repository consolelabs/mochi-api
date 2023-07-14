package model

import "github.com/google/uuid"

type UserNFTBalance struct {
	UserAddress     string         `json:"user_address"`
	ChainType       JSONNullString `json:"chain_type"`
	NFTCollectionID uuid.NullUUID  `json:"nft_collection_id" swaggertype:"string"`
	TokenID         string         `json:"token_id,omitempty"`
	Balance         int            `json:"balance"`
}

type UserNFTBalancesByGuild struct {
	UserDiscordId string `json:"user_discord_id"`
	TotalBalance  int64  `json:"total_balance"`
}

type UserAddressNFTBalancesByGuild struct {
	UserAddress     string `json:"user_address"`
	TotalBalance    int64  `json:"total_balance"`
	StakingNeko     int64  `json:"staking_neko"`
	NftCollectionID string `json:"nft_collection_id"`
}

type UserNFTBalanceIdentify struct {
	UserDiscordId   string `json:"user_discord_id"`
	NftCollectionId string `json:"nft_collection_id"`
}
