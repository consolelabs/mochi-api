package model

import (
	"time"
)

type AutoTrigger struct {
	Id             int64           `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	DiscordGuildId string          `json:"discord_guild_id"`
	UserDiscordId  string          `json:"user_discord_id"`
	Name           string          `json:"name"`
	Status         bool            `json:"status"`
	Conditions     []AutoCondition `json:"auto_condition" gorm:"foreignKey:TriggerId;references:Id"`
	Actions        []AutoAction    `json:"auto_action" gorm:"foreignKey:TriggerId;references:Id"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
