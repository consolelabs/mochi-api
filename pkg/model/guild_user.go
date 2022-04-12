package model

import (
	"github.com/google/uuid"
)

type GuildUser struct {
	ID      uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID int64         `json:"guild_id"`
	UserID  int64         `json:"user_id"`
	Roles   []GuildRole   `json:"roles" gorm:"many2many:guild_user_role;foreignKey:UserID;joinForeignKey:UserID;References:ID;joinReferences:RoleID"`
}

type GuildUserRole struct {
	ID      uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID int64         `json:"guild_id"`
	UserID  int64         `json:"user_id"`
	RoleID  int64         `json:"role_id"`
}
