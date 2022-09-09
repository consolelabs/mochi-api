package model

import "github.com/google/uuid"

type GuildConfigNFTRole struct {
	ID              uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	NFTCollectionID uuid.NullUUID `json:"nft_collection_id" swaggertype:"string"`
	NumberOfTokens  int           `json:"number_of_tokens"`
	GroupID         uuid.NullUUID `json:"group_id"`
}

type MemberNFTRole struct {
	UserDiscordID string `json:"user_id"`
	RoleID        string `json:"role_id"`
}
