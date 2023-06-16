package model

import (
	"time"

	"github.com/google/uuid"
)

type AutoCondition struct {
	Id        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TriggerId string        `json:"trigger_id" gorm:"unique_index:idx_guild_user_guild_id_user_id"`
	TypeId    string        `json:"type_id"`
	ChannelId string        `json:"channel_id"`
	UserIds   string        `json:"user_id"`
	ChildId   string        `json:"child_id"`
	Index     int           `json:"index"`
	Platform  string        `json:"platform"`
	IsPrimary bool          `json:"is_primary"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`

	ConditionValues []AutoConditionValue `json:"auto_condition_value" gorm:"foreignKey:ConditionId;references:Id"`
	ChildConditions []AutoCondition      `json:"child_condition" gorm:"foreignKey:ChildId;references:Id"`
	Type            AutoType             `json:"auto_type" gorm:"foreignKey:TypeId;references:Id"`
}

type AutoType struct {
	Id        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	IconUrl   string        `json:"icon_url"`
	CreatedAt time.Time     `json:"created_at"`
}

type AutoTypePrefix struct {
	Id        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TypeId    string        `json:"type_id"`
	Value     string        `json:"value"`
	CreatedAt time.Time     `json:"created_at"`
}
