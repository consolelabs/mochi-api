package model

import "time"

type GuildConfigTokenRole struct {
	ID             int       `json:"id"`
	GuildID        string    `json:"guild_id"`
	RoleID         string    `json:"channel_id"`
	RequiredAmount float64   `json:"required_amount" gorm:"type:numeric"`
	TokenID        int       `json:"token_id"`
	Token          *Token    `json:"token,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (GuildConfigTokenRole) TableName() string {
	return "guild_config_token_roles"
}
