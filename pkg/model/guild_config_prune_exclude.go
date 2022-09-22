package model

import "github.com/google/uuid"

type GuildConfigWhitelistPrune struct {
	ID      uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID string        `json:"guild_id"`
	RoleID  string        `json:"role_id"`
}
