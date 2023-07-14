package model

import (
	"time"

	"github.com/google/uuid"
)

type GuildConfigGroupNFTRole struct {
	ID                 uuid.NullUUID        `json:"id" gorm:"default:uuid_generate_v4()"`
	GroupName          string               `json:"group_name"`
	GuildID            string               `json:"guild_id"`
	RoleID             string               `json:"role_id"`
	NumberOfTokens     int                  `json:"number_of_tokens"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	GuildConfigNFTRole []GuildConfigNFTRole `json:"guild_config_nft_role" gorm:"foreignkey:group_id"`
}

type UpdateUserRolesOptions struct {
	// GuildID is the guild ID to update token roles
	GuildID string
}
