package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TargetGroup string

const (
	TargetGroupAll       TargetGroup = "all"
	TargetGroupReceivers TargetGroup = "receivers"
	TargetGroupFriends   TargetGroup = "friends"
)

type BasePrivacySetting struct {
	Enable      bool        `json:"enable"`
	TargetGroup TargetGroup `json:"target_group"`
}

type UserPrivacySetting struct {
	ProfileId         string              `json:"profile_id"`
	DestinationWallet *BasePrivacySetting `json:"destination_wallet" gorm:"type:jsonb"`
}

// db explaination for select query
func (s *BasePrivacySetting) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("money source has unsupported type: %v", v)
	}

	return json.Unmarshal(bytes, s)
}

// db explaination for insert/update query
func (s BasePrivacySetting) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}
