package model

import "github.com/google/uuid"

type GuildConfigNFTRole struct {
	ID              uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	NFTCollectionID uuid.NullUUID `json:"nft_collection_id" swaggertype:"string"`
	NumberOfTokens  int           `json:"number_of_tokens"`
	GroupID         uuid.NullUUID `json:"group_id"`
}

type MemberNFTRole struct {
	UserID             string `json:"user_id"`
	RoleID             string `json:"role_id"`
	TotalBalance       int64  `json:"total_balance"`
	GroupConfigBalance int64  `json:"group_config_balance"`
	GroupConfigID      string `json:"group_config_id"`
}
