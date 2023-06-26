package model

import (
	"time"
)

type AutoCondition struct {
	Id        int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TriggerId int64     `json:"trigger_id" gorm:"unique_index:idx_guild_user_guild_id_user_id"`
	TypeId    int64     `json:"type_id"`
	ChannelId string    `json:"channel_id"`
	UserIds   string    `json:"user_id"`
	ChildId   int64     `json:"child_id"`
	Index     int       `json:"index"`
	Platform  string    `json:"platform"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	ConditionValues []AutoConditionValue `json:"auto_condition_value" gorm:"foreignKey:ConditionId;references:Id"`
	ChildConditions []AutoCondition      `json:"child_condition" gorm:"foreignKey:ChildId;references:Id"`
	Type            AutoType             `json:"auto_type" gorm:"foreignKey:TypeId;references:Id"`
}

type AutoType struct {
	Id        int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	IconUrl   string    `json:"icon_url"`
	CreatedAt time.Time `json:"created_at"`
}

type AutoTypePrefix struct {
	Id        int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TypeId    string    `json:"type_id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}
