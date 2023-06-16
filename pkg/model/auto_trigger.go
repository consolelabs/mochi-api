package model

import (
	"time"

	"github.com/google/uuid"
)

type AutoTrigger struct {
	Id         uuid.NullUUID   `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildId    string          `json:"guild_id"`
	UserId     string          `json:"user_id"`
	Name       string          `json:"name"`
	Status     bool            `json:"status"`
	Conditions []AutoCondition `json:"auto_condition" gorm:"foreignKey:TriggerId;references:Id"`
	Actions    []AutoAction    `json:"auto_action" gorm:"foreignKey:TriggerId;references:Id"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
