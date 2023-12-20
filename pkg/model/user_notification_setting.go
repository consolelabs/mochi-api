package model

import (
	"github.com/lib/pq"
)

// key = setting key, value = true/false = enbale/disable
type SettingFlags map[string]bool

type UserNotificationSetting struct {
	ProfileId            string             `json:"profile_id"`
	Enable               bool               `json:"enable"`
	Platforms            pq.StringArray     `json:"platforms" gorm:"type:json"`
	Flags                SettingFlags       `json:"flags" gorm:"type:json"`
	NotificationSettings []NotificationFlag `json:"notification_settings" gorm:"type:-"`
}
