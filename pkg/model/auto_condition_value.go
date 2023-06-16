package model

import (
	"time"

	"github.com/google/uuid"
)

type AutoConditionValue struct {
	Id          uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TypeId      string        `json:"type_id"`
	ConditionId string        `json:"condition_id"`
	ChildId     string        `json:"child_id"`
	Index       int           `json:"index"`
	Operator    string        `json:"operator"`
	Matches     string        `json:"matches"`
	CreatedAt   time.Time     `json:"created_at"`

	Type AutoType `json:"auto_type" gorm:"foreignKey:TypeId;references:Id"`
}
