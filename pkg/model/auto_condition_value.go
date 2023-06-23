package model

import (
	"time"
)

type AutoConditionValue struct {
	Id          int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TypeId      int64     `json:"type_id"`
	ConditionId string    `json:"condition_id"`
	ChildId     string    `json:"child_id"`
	Index       int       `json:"index"`
	Operator    string    `json:"operator"`
	Matches     string    `json:"matches"`
	CreatedAt   time.Time `json:"created_at"`

	Type AutoType `json:"auto_type" gorm:"foreignKey:TypeId;references:Id"`
}
