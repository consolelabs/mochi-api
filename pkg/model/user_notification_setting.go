package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

// key = setting key, value = true/false = enbale/disable
type SettingFlags map[string]bool

type UserNotificationSetting struct {
	ProfileId            string             `json:"profile_id"`
	Enable               bool               `json:"enable"`
	Platforms            pq.StringArray     `json:"platforms" gorm:"type:jsonb"`
	Flags                SettingFlags       `json:"flags" gorm:"type:jsonb"`
	NotificationSettings []NotificationFlag `json:"notification_settings" gorm:"-"`
}

// db explaination for select query
func (s *SettingFlags) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("flags has unsupported type: %v", v)
	}

	return json.Unmarshal(bytes, s)
}

// db explaination for insert/update query
func (s SettingFlags) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}
