package model

import "time"

type MixRoleTokenRequirement struct {
	ID             int       `json:"id"`
	TokenID        int       `json:"token_id"`
	RequiredAmount float64   `json:"required_amount"`
	Token          *Token    `json:"token,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MixRoleNFTRequirement struct {
	ID              int            `json:"id"`
	NFTCollectionID string         `json:"nft_collection_id"`
	RequiredAmount  int            `json:"required_amount"`
	NFTCollection   *NFTCollection `json:"nft_collection,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type GuildConfigMixRole struct {
	ID                 int                      `json:"id"`
	GuildID            string                   `json:"guild_id"`
	RoleID             string                   `json:"role_id"`
	TokenRequirementID *int                     `json:"token_requirement_id,omitempty"`
	NFTRequirementID   *int                     `json:"nft_requirement_id,omitempty"`
	RequiredLevel      int                      `json:"required_level"`
	TokenRequirement   *MixRoleTokenRequirement `json:"token_requirement,omitempty"`
	NFTRequirement     *MixRoleNFTRequirement   `json:"nft_requirement,omitempty"`
	CreatedAt          time.Time                `json:"created_at"`
	UpdatedAt          time.Time                `json:"updated_at"`
}

type MemberMixRole struct {
	UserDiscordID string `json:"user_id"`
	RoleID        string `json:"role_id"`
}
