package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

var payment model.UserPaymentSetting = model.UserPaymentSetting{
	ProfileId: "50453",
	DefaultMoneySource: model.MoneySource{
		Platform:           "evm",
		PlatformIdentifier: "0xa1bc23fabefbaabfaebb446",
	},
	DefaultReceiverPlatform: "discord",
	PrioritizedTokenIds:     []string{"941f0fb1-00da-49dc-a538-5e81fc874cb4", "61388b7c-5505-4fdf-8084-077422369a93"},
	DefaultTokenId:          "941f0fb1-00da-49dc-a538-5e81fc874cb4",
	DefaultMessageEnable:    true,
	DefaultMessageSettings: []model.DefaultMessageSetting{
		{Action: "tip", Message: "take my money"},
	},
	TxLimitEnable: true,
	TxLimitSettings: []model.TxLimitSetting{
		{Action: "tip", Max: 20},
		{Action: "airdrop", Min: 2},
	},
	DefaultToken: &model.PayToken{Id: "941f0fb1-00da-49dc-a538-5e81fc874cb4", Symbol: "ICY"},
	PrioritizedTokens: []model.PayToken{
		{Id: "941f0fb1-00da-49dc-a538-5e81fc874cb4", Symbol: "ICY"},
		{Id: "61388b7c-5505-4fdf-8084-077422369a93", Symbol: "FTM"},
	},
}

var privacy model.UserPrivacySetting = model.UserPrivacySetting{
	ProfileId: "50453",
	Tx: &model.BasePrivacySetting{
		GeneralTargetGroup:   model.TargetGroupAll,
		GeneralPlatformGroup: model.PlatformGroupCustom,
		CustomSettings: []model.PrivacyCustomSetting{
			{TargetGroup: model.TargetGroupFriends, Platform: "discord"},
			{TargetGroup: model.TargetGroupReceivers, Platform: "telegram"},
			{TargetGroup: model.TargetGroupAll, Platform: "website"},
		},
	},
	SocialAccounts: &model.BasePrivacySetting{
		GeneralTargetGroup:   model.TargetGroupAll,
		GeneralPlatformGroup: model.PlatformGroupCustom,
		CustomSettings: []model.PrivacyCustomSetting{
			{TargetGroup: model.TargetGroupFriends, Platform: "discord"},
			{TargetGroup: model.TargetGroupReceivers, Platform: "telegram"},
			{TargetGroup: model.TargetGroupAll, Platform: "website"},
		},
	},
	Wallets: &model.BasePrivacySetting{
		GeneralTargetGroup:   model.TargetGroupAll,
		GeneralPlatformGroup: model.PlatformGroupCustom,
		CustomSettings: []model.PrivacyCustomSetting{
			{TargetGroup: model.TargetGroupFriends, Platform: "discord"},
			{TargetGroup: model.TargetGroupReceivers, Platform: "telegram"},
			{TargetGroup: model.TargetGroupAll, Platform: "website"},
		},
	},
}

func (e *Entity) GetUserGeneralSettings(uri request.UserSettingBaseUriRequest) (*model.UserPaymentSetting, *model.UserPrivacySetting, error) {
	return &payment, &privacy, nil
}

func (e *Entity) UpdateUserGeneralSettings(uri request.UserSettingBaseUriRequest, payload request.UpdateGeneralSettingsPayloadRequest) (*model.UserPaymentSetting, *model.UserPrivacySetting, error) {
	return &payment, &privacy, nil
}

var notificationSettings []model.NotificationFlag = []model.NotificationFlag{
	{Group: model.NotificationGroupWallet, Key: "disable_all", Description: "Disable all notification wallets"},
	{Group: model.NotificationGroupWallet, Key: "receive_airdrop_success", Description: "Receive airdrops"},
	{Group: model.NotificationGroupWallet, Key: "receive_deposit_success", Description: "Deposit completed"},
	{Group: model.NotificationGroupWallet, Key: "send_withdraw_success", Description: "Withdrawal completed"},
	{Group: model.NotificationGroupWallet, Key: "receive_payme_success", Description: "Payment request completed"},
	{Group: model.NotificationGroupWallet, Key: "*_payme_expired", Description: "Payment request expired"},
	{Group: model.NotificationGroupWallet, Key: "*_paylink_expired", Description: "Pay link has expired"},
	{Group: model.NotificationGroupWallet, Key: "send_paylink_success", Description: "Pay link claimed by another"},
	{Group: model.NotificationGroupWallet, Key: "receive_paylink_success", Description: "Claim a pay link"},
	{Group: model.NotificationGroupCommunity, Key: "new_configuration", Description: "New configuration"},
	{Group: model.NotificationGroupApp, Key: "new_vault_tx", Description: "New vault transactions"},
	{Group: model.NotificationGroupApp, Key: "new_api_call", Description: "New API calls"},
	{Group: model.NotificationGroupApp, Key: "info_updated", Description: "Information changes"},
	{Group: model.NotificationGroupApp, Key: "new_member", Description: "New members"},
}

var userNotiSettings model.UserNotificationSetting = model.UserNotificationSetting{
	ProfileId: "50453",
	Enable:    true,
	Platforms: []string{"discord", "telegram"},
	Flags: map[string]bool{
		"disable_all":             false,
		"receive_tip_success":     true,
		"receive_airdrop_success": false,
		"receive_deposit_success": true,
		"send_withdraw_success":   false,
		"receive_payme_success":   true,
		"*_payme_expired":         false,
		"*_paylink_expired":       true,
		"send_paylink_success":    false,
		"receive_paylink_success": true,
		"new_configuration":       false,
		"new_vault_tx":            true,
		"new_api_call":            true,
		"info_updated":            true,
		"new_member":              false,
	},
	NotificationSettings: notificationSettings,
}

func (e *Entity) GetUserNotificationSettings(uri request.UserSettingBaseUriRequest) (*model.UserNotificationSetting, error) {
	return &userNotiSettings, nil
}

func (e *Entity) UpdateUserNotificationSettings(uri request.UserSettingBaseUriRequest, payload request.UpdateGeneralNotificationSettingPayloadRequest) (*model.UserNotificationSetting, error) {
	return &userNotiSettings, nil
}

func (e *Entity) UpdateUserActivityNotificationSettings(uri request.UpdateActivityNotificationSettingUriRequest, payload request.UpdateActivityNotificationSettingPayloadRequest) (*model.UserNotificationSetting, error) {
	return &userNotiSettings, nil
}
