package entities

import (
	"github.com/consolelabs/mochi-toolkit/formatter"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	notificationflag "github.com/defipod/mochi/pkg/repo/notification_flag"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) initUserPaymentSetting(profileId string) model.UserPaymentSetting {
	return model.UserPaymentSetting{
		ProfileId: profileId,
		DefaultMoneySource: model.MoneySource{
			Platform:           "mochi",
			PlatformIdentifier: "mochi-balance",
		},
		DefaultReceiverPlatform: "discord",
		PrioritizedTokenIds:     []string{},
		DefaultMessageEnable:    false,
		DefaultMessageSettings:  []model.DefaultMessageSetting{},
		TxLimitEnable:           false,
		TxLimitSettings:         []model.TxLimitSetting{},
	}
}

func (e *Entity) initUserPrivacySetting(profileId string) model.UserPrivacySetting {
	defaultCustomSettings := []model.PrivacyCustomSetting{
		{TargetGroup: model.TargetGroupAll, Platform: formatter.PlatformDiscord},
		{TargetGroup: model.TargetGroupAll, Platform: formatter.PlatformTelegram},
		{TargetGroup: model.TargetGroupAll, Platform: formatter.PlatformWeb},
	}

	return model.UserPrivacySetting{
		ProfileId: profileId,
		Tx: &model.BasePrivacySetting{
			GeneralTargetGroup:   model.TargetGroupAll,
			GeneralPlatformGroup: model.PlatformGroupAll,
			CustomSettings:       defaultCustomSettings,
		},
		SocialAccounts: &model.BasePrivacySetting{
			GeneralTargetGroup:   model.TargetGroupAll,
			GeneralPlatformGroup: model.PlatformGroupAll,
			CustomSettings:       defaultCustomSettings,
		},
		Wallets: &model.BasePrivacySetting{
			GeneralTargetGroup:   model.TargetGroupAll,
			GeneralPlatformGroup: model.PlatformGroupAll,
			CustomSettings:       defaultCustomSettings,
		},
	}
}

func (e *Entity) GetUserGeneralSettings(uri request.UserSettingBaseUriRequest) (*model.UserPaymentSetting, *model.UserPrivacySetting, error) {
	logger := e.log.Fields(logger.Fields{
		"component":  "entity.setting.GetUserGeneralSettings",
		"profile_id": uri.ProfileId,
	})

	// compose init settings for user
	initPayment := e.initUserPaymentSetting(uri.ProfileId)
	initPrivacy := e.initUserPrivacySetting(uri.ProfileId)

	/////// start working with db layer
	tx, fn := e.repo.Store.NewTransaction()

	// get user's payment setting or create if not exists
	payment, err := tx.UserPaymentSetting.FirstOrCreate(initPayment)
	if err != nil {
		logger.Error(err, "tx.UserPaymentSetting.FirstOrCreate() failed")
		fn.Rollback(err)
		return nil, nil, err
	}

	// update privacy setting or create if not exists
	privacy, err := tx.UserPrivacySetting.FirstOrCreate(initPrivacy)
	if err != nil {
		logger.Error(err, "tx.UserPrivacySetting.FirstOrCreate() failed")
		fn.Rollback(err)
		return nil, nil, err
	}

	// commit db tx
	err = fn.Commit()
	if err != nil {
		logger.Error(err, "fn.Commit() failed")
		fn.Rollback(err)
		return nil, nil, err
	}

	return payment, privacy, nil
}

func (e *Entity) UpdateUserGeneralSettings(uri request.UserSettingBaseUriRequest, payload request.UpdateGeneralSettingsPayloadRequest) (*model.UserPaymentSetting, *model.UserPrivacySetting, error) {
	logger := e.log.Fields(logger.Fields{
		"component":  "entity.setting.UpdateUserGeneralSettings",
		"profile_id": uri.ProfileId,
	})

	/////// prepare data
	// payment settings
	defaultMessages := make([]model.DefaultMessageSetting, len(payload.Payment.DefaultMessageSettings))
	for i, s := range payload.Payment.DefaultMessageSettings {
		defaultMessages[i] = model.DefaultMessageSetting{
			Action:  s.Action,
			Message: s.Message,
		}
	}

	txLimits := make([]model.TxLimitSetting, len(payload.Payment.TxLimitSettings))
	for i, s := range payload.Payment.TxLimitSettings {
		txLimits[i] = model.TxLimitSetting{
			Action: s.Action,
			Min:    s.Min,
			Max:    s.Max,
		}
	}

	// transform custom privacy settings
	txPrivacyCustom := make([]model.PrivacyCustomSetting, len(payload.Privacy.Tx.CustomSettings))
	for i, s := range payload.Privacy.Tx.CustomSettings {
		txPrivacyCustom[i] = model.PrivacyCustomSetting{
			TargetGroup: model.TargetGroup(s.TargetGroup),
			Platform:    s.Platform,
		}
	}
	socialAccsPrivacyCustom := make([]model.PrivacyCustomSetting, len(payload.Privacy.Tx.CustomSettings))
	for i, s := range payload.Privacy.SocialAccounts.CustomSettings {
		socialAccsPrivacyCustom[i] = model.PrivacyCustomSetting{
			TargetGroup: model.TargetGroup(s.TargetGroup),
			Platform:    s.Platform,
		}
	}
	walletsPrivacyCustom := make([]model.PrivacyCustomSetting, len(payload.Privacy.Tx.CustomSettings))
	for i, s := range payload.Privacy.Wallets.CustomSettings {
		walletsPrivacyCustom[i] = model.PrivacyCustomSetting{
			TargetGroup: model.TargetGroup(s.TargetGroup),
			Platform:    s.Platform,
		}
	}

	privacy := model.UserPrivacySetting{
		ProfileId: uri.ProfileId,
		Tx: &model.BasePrivacySetting{
			GeneralTargetGroup:   model.TargetGroup(payload.Privacy.Tx.GeneralTargetGroup),
			GeneralPlatformGroup: model.PlatformGroup(payload.Privacy.Tx.GeneralPlatformGroup),
			CustomSettings:       txPrivacyCustom,
		},
		SocialAccounts: &model.BasePrivacySetting{
			GeneralTargetGroup:   model.TargetGroup(payload.Privacy.SocialAccounts.GeneralTargetGroup),
			GeneralPlatformGroup: model.PlatformGroup(payload.Privacy.SocialAccounts.GeneralPlatformGroup),
			CustomSettings:       socialAccsPrivacyCustom,
		},
		Wallets: &model.BasePrivacySetting{
			GeneralTargetGroup:   model.TargetGroup(payload.Privacy.Wallets.GeneralTargetGroup),
			GeneralPlatformGroup: model.PlatformGroup(payload.Privacy.Wallets.GeneralPlatformGroup),
			CustomSettings:       walletsPrivacyCustom,
		},
	}

	payment := model.UserPaymentSetting{
		ProfileId:               uri.ProfileId,
		DefaultMoneySource:      model.MoneySource(payload.Payment.DefaultMoneySource),
		DefaultReceiverPlatform: payload.Payment.DefaultReceiverPlatform,
		PrioritizedTokenIds:     payload.Payment.TokenPriorities,
		DefaultMessageEnable:    *payload.Payment.DefaultMessageEnable,
		DefaultMessageSettings:  defaultMessages,
		TxLimitEnable:           *payload.Payment.TxLimitEnable,
		TxLimitSettings:         txLimits,
	}

	/////// start working with db layer
	tx, fn := e.repo.Store.NewTransaction()

	// update payment setting
	err := tx.UserPaymentSetting.Update(&payment)
	if err != nil {
		logger.Error(err, "tx.UserPaymentSetting.Update() failed")
		fn.Rollback(err)
		return nil, nil, err
	}

	// update privacy setting
	err = tx.UserPrivacySetting.Update(&privacy)
	if err != nil {
		logger.Error(err, "tx.UserPrivacySetting.Update() failed")
		fn.Rollback(err)
		return nil, nil, err
	}

	// commit db tx
	err = fn.Commit()
	if err != nil {
		logger.Error(err, "fn.Commit() failed")
		fn.Rollback(err)
		return nil, nil, err
	}

	return &payment, &privacy, nil
}

func (e *Entity) initUserNotiSetting(profileId string, settings []model.NotificationFlag) model.UserNotificationSetting {
	userFlags := make(map[string]bool)
	for _, f := range settings {
		userFlags[f.Key] = f.Key != "disable_all"
	}

	return model.UserNotificationSetting{
		ProfileId:            profileId,
		Enable:               true,
		Platforms:            []string{formatter.PlatformDiscord, formatter.PlatformTelegram, formatter.PlatformWeb},
		Flags:                userFlags,
		NotificationSettings: settings,
	}
}

func (e *Entity) GetUserNotificationSettings(uri request.UserSettingBaseUriRequest) (*model.UserNotificationSetting, error) {
	logger := e.log.Fields(logger.Fields{
		"component":  "entity.setting.GetUserNotificationSettings",
		"profile_id": uri.ProfileId,
	})

	// get all notification flags
	var listQ notificationflag.ListQuery
	flags, err := e.repo.NotificationFlag.List(listQ)
	if err != nil {
		logger.Error(err, "repo.NotificationFlag.List() failed")
		return nil, err
	}

	// compose init noti setting for user
	initNotiSetting := e.initUserNotiSetting(uri.ProfileId, flags)

	/////// start working with db layer
	tx, fn := e.repo.Store.NewTransaction()

	// update payment setting
	userNotiSettings, err := tx.UserNotificationSetting.FirstOrCreate(initNotiSetting)
	if err != nil {
		logger.Error(err, "tx.UserNotificationSetting.FirstOrCreate() failed")
		fn.Rollback(err)
		return nil, err
	}
	userNotiSettings.NotificationSettings = flags

	// commit db tx
	err = fn.Commit()
	if err != nil {
		logger.Error(err, "fn.Commit() failed")
		fn.Rollback(err)
		return nil, err
	}

	return userNotiSettings, nil
}

func (e *Entity) ListAllNotificationFlags() ([]model.NotificationFlag, error) {
	logger := e.log.Fields(logger.Fields{
		"component": "entity.setting.ListAllNotificationFlags",
	})
	// get all notification flags
	var listQ notificationflag.ListQuery
	notificationFlags, err := e.repo.NotificationFlag.List(listQ)
	if err != nil {
		logger.Error(err, "repo.NotificationFlag.List() failed")
		return nil, err
	}

	return notificationFlags, nil
}

func (e *Entity) UpdateUserNotificationSettings(uri request.UserSettingBaseUriRequest, payload request.UpdateNotificationSettingPayloadRequest, notiSettings []model.NotificationFlag) (*model.UserNotificationSetting, error) {
	logger := e.log.Fields(logger.Fields{
		"component":  "entity.setting.UpdateUserNotificationSettings",
		"profile_id": uri.ProfileId,
	})

	/////// start working with db layer
	tx, fn := e.repo.Store.NewTransaction()

	// update payment setting
	userNotiSettings := model.UserNotificationSetting{
		ProfileId:            uri.ProfileId,
		Enable:               *payload.Enable,
		Platforms:            payload.Platforms,
		Flags:                payload.Flags,
		NotificationSettings: notiSettings,
	}
	err := tx.UserNotificationSetting.Update(&userNotiSettings)
	if err != nil {
		logger.Error(err, "tx.UserNotificationSetting.Update() failed")
		fn.Rollback(err)
		return nil, err
	}

	// commit db tx
	err = fn.Commit()
	if err != nil {
		logger.Error(err, "fn.Commit() failed")
		fn.Rollback(err)
		return nil, err
	}

	return &userNotiSettings, nil
}
