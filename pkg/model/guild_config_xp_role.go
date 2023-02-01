package model

import "time"

type GuildConfigXPRole struct {
	ID         int       `json:"id"`
	GuildID    string    `json:"guild_id"`
	RoleID     string    `json:"role_id"`
	RequiredXP int       `json:"required_xp"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (GuildConfigXPRole) TableName() string {
	return "guild_config_xp_roles"
}

type MemberXPRole struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}
