package entities

import (
	"errors"
	"strings"

	"github.com/consolelabs/mochi-toolkit/formatter"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	notificationflag "github.com/defipod/mochi/pkg/repo/notification_flag"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	sliceutils "github.com/defipod/mochi/pkg/util/slice"
)

func (e *Entity) initUserPaymentSetting(profileId string) model.UserPaymentSetting {
	balances, err := e.svc.MochiPay.GetListBalances(profileId)
	if err != nil {
		e.log.Error(err, "svc.MochiPay.GetListBalances() failed")
	}

	prioritizedTokens := make([]string, 0)
	if balances != nil && balances.Data != nil {
		for _, b := range balances.Data {
			prioritizedTokens = append(prioritizedTokens, b.TokenId)
		}
	}

	return model.UserPaymentSetting{
		ProfileId: profileId,
		DefaultMoneySource: model.MoneySource{
			Platform:           "mochi",
			PlatformIdentifier: "mochi-balance",
		},
		DefaultReceiverPlatform: "discord",
		PrioritizedTokenIds:     prioritizedTokens,
		DefaultMessageEnable:    false,
		DefaultMessageSettings:  []model.DefaultMessageSetting{},
		TxLimitEnable:           false,
		TxLimitSettings:         []model.TxLimitSetting{},
	}
}

func (e *Entity) initUserPrivacySetting(profileId string) model.UserPrivacySetting {
	return model.UserPrivacySetting{
		ProfileId:             profileId,
		ShowDestinationWallet: true,
		TxTargetGroup:         model.TargetGroupAll,
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
			Enable:  *s.Enable,
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

	privacy := model.UserPrivacySetting{
		ProfileId:             uri.ProfileId,
		ShowDestinationWallet: *payload.Privacy.ShowDestinationWallet,
		TxTargetGroup:         model.TargetGroup(payload.Privacy.TxTargetGroup),
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

func (e *Entity) ValidateMoneySourceSetting(profileId string, s request.MoneySource) error {
	logger := e.log.Fields(logger.Fields{
		"component":  "entity.setting.ValidateMoneySourceSetting",
		"profile_id": profileId,
	})

	mochiSource := mochiprofile.AssociatedAccount{
		Platform:           "mochi",
		PlatformIdentifier: "mochi-balance",
	}
	if string(mochiSource.Platform) == s.Platform && mochiSource.PlatformIdentifier == s.PlatformIdentifier {
		return nil
	}

	profile, err := e.svc.MochiProfile.GetByID(profileId, e.cfg.MochiBotSecret)
	if err != nil {
		logger.Error(err, "svc.MochiProfile.GetByID() failed")
		return errors.New("failed to validate profile accounts")
	}

	if profile == nil || profile.AssociatedAccounts == nil || len(profile.AssociatedAccounts) == 0 {
		return errors.New("invalid money source")
	}

	// a money source is considered as valid when it's either mochi wallet or connected wallet (evm,sol,etc.)
	// other social platforms such as telegram, discord, email are invalid money source
	validMoneySource := sliceutils.Some(profile.AssociatedAccounts, func(acc mochiprofile.AssociatedAccount) bool {
		existingSource := string(acc.Platform) == s.Platform && strings.EqualFold(acc.PlatformIdentifier, s.PlatformIdentifier)
		isConnectedWallet := strings.Contains(s.Platform, "chain") && existingSource
		return isConnectedWallet
	})
	if !validMoneySource {
		return errors.New("invalid money source")
	}

	return nil
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
