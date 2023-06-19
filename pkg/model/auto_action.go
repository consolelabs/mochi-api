package model

import (
	"time"

	"github.com/google/uuid"
)

type AutoAction struct {
	Id           uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TriggerId    string        `json:"trigger_id" gorm:"not null"`
	TypeId       string        `json:"type_id" gorm:"not null"`
	ChannelId    string        `json:"channel_id"`
	ChildId      string        `json:"child_id"`
	Index        int           `json:"index"`
	Platform     string        `json:"platform"`
	Content      string        `json:"content"`
	IsPrimary    bool          `json:"is_primary"`
	CreatedAt    time.Time     `json:"created_at"`
	ActionData   string        `json:"action_data"`
	LimitPerUser int           `json:"limit_per_user"`

	Type  AutoType  `json:"auto_type" gorm:"foreignKey:TypeId;references:Id"`
	Embed AutoEmbed `json:"auto_embed" gorm:"foreignKey:Id;references:ActionId"`
}

type AutoActionHistory struct {
	Id        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TriggerId string        `json:"trigger_id" gorm:"not null"`
	ActionId  string        `json:"action_id" gorm:"not null"`
	UserId    string        `json:"user_id" gorm:"not null"`
	MessageId string        `json:"message_id" gorm:"not null"`
	Total     int           `json:"total"`

	Action AutoAction `json:"auto_action" gorm:"foreignKey:ActionId;references:Id"`

	CreatedAt time.Time `json:"created_at"`
}
