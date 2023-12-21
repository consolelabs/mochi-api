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

type PlatformGroup string

const (
	PlatformGroupAll    PlatformGroup = "all"
	PlatformGroupCustom PlatformGroup = "custom"
)

type PrivacyCustomSetting struct {
	TargetGroup TargetGroup `json:"target_group"`
	Platform    string      `json:"platform"`
}

type PrivacyCustomSettings []PrivacyCustomSetting

// db explaination for select query
func (s *PrivacyCustomSettings) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("privacy_custom_settings has unsupported type: %v", v)
	}

	return json.Unmarshal(bytes, s)
}

// db explaination for insert/update query
func (s PrivacyCustomSettings) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

type BasePrivacySetting struct {
	GeneralTargetGroup   TargetGroup           `json:"general_target_group"`
	GeneralPlatformGroup PlatformGroup         `json:"general_platform_group"`
	CustomSettings       PrivacyCustomSettings `json:"custom_settings"`
}

type UserPrivacySetting struct {
	ProfileId      string              `json:"profile_id"`
	Tx             *BasePrivacySetting `json:"tx" gorm:"type:jsonb"`
	SocialAccounts *BasePrivacySetting `json:"social_accounts" gorm:"type:jsonb"`
	Wallets        *BasePrivacySetting `json:"wallets" gorm:"type:jsonb"`
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
