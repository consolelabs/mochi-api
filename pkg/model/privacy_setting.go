package model

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

type BasePrivacySetting struct {
	GeneralTargetGroup   TargetGroup            `json:"general_target_group"`
	GeneralPlatformGroup PlatformGroup          `json:"general_platform_group"`
	CustomSettings       []PrivacyCustomSetting `json:"custom_settings"`
}

type UserPrivacySetting struct {
	ProfileId      string              `json:"profile_id"`
	Tx             *BasePrivacySetting `json:"tx" gorm:"type:json"`
	SocialAccounts *BasePrivacySetting `json:"social_accounts" gorm:"type:json"`
	Wallets        *BasePrivacySetting `json:"wallets" gorm:"type:json"`
}
