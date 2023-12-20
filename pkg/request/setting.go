package request

type UserSettingBaseUriRequest struct {
	ProfileId string `uri:"profile_id" binding:"required"`
}

type MoneySource struct {
	Platform           string `json:"platform" binding:"required"`
	PlatformIdentifier string `json:"platform_identifier" binding:"required"`
}

type DefaultMessageSetting struct {
	Action  string `json:"action" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type TxLimitSetting struct {
	Action string  `json:"action" binding:"required"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
}

type UpdateGeneralSettingsPayloadRequest struct {
	Payment *PaymentSetting `json:"payment" binding:"required"`
	Privacy *PrivacySetting `json:"privacy" binding:"required"`
}

type PaymentSetting struct {
	DefaultMoneySource      MoneySource             `json:"default_money_source" binding:"required"`
	DefaultReceiverPlatform string                  `json:"default_receiver_platform" binding:"required"`
	TokenPriorities         []string                `json:"token_priorities" binding:"required"`
	DefaultToken            string                  `json:"default_token" binding:"required"`
	DefaultMessageEnable    *bool                   `json:"default_message_enable" binding:"required"`
	DefaultMessageSettings  []DefaultMessageSetting `json:"default_message_settings" binding:"required"`
	TxLimitEnable           *bool                   `json:"tx_limit_enable" binding:"required"`
	TxLimitSettings         []TxLimitSetting        `json:"tx_limit_settings" binding:"required"`
}

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

type PrivacySetting struct {
	Tx             *BasePrivacySetting `json:"tx"`
	SocialAccounts *BasePrivacySetting `json:"social_accounts"`
	Wallets        *BasePrivacySetting `json:"wallets"`
}

type UpdateNotificationSettingPayloadRequest struct {
	Enable    bool            `json:"enable" binding:"required"`
	Platforms []string        `json:"platforms" binding:"required"`
	Flags     map[string]bool `json:"flag" binding:"required"`
}
