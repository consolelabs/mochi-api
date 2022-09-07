package model

import "github.com/google/uuid"

type GuildConfigNFTRole struct {
	ID              uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	NFTCollectionID uuid.NullUUID `json:"nft_collection_id" swaggertype:"string"`
	GuildID         string        `json:"guild_id"`
	RoleID          string        `json:"role_id"`
	NumberOfTokens  int           `json:"number_of_tokens"`
	TokenID         string        `json:"token_id,omitempty"`
}

type MemberNFTRole struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}
