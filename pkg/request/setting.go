package request

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/model"
	sliceutils "github.com/defipod/mochi/pkg/util/slice"
	uuidutils "github.com/defipod/mochi/pkg/util/uuid"
)

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
	DefaultMessageEnable    *bool                   `json:"default_message_enable" binding:"required"`
	DefaultMessageSettings  []DefaultMessageSetting `json:"default_message_settings" binding:"required"`
	TxLimitEnable           *bool                   `json:"tx_limit_enable" binding:"required"`
	TxLimitSettings         []TxLimitSetting        `json:"tx_limit_settings" binding:"required"`
}

type PrivacyCustomSetting struct {
	TargetGroup string `json:"target_group" binding:"required"`
	Platform    string `json:"platform" binding:"required"`
}

type BasePrivacySetting struct {
	Enable      bool   `json:"enable" binding:"required"`
	TargetGroup string `json:"target_group" binding:"required"`
}

type PrivacySetting struct {
	DestinationWalet *BasePrivacySetting `json:"destination_wallet" binding:"required"`
}

func (r *UpdateGeneralSettingsPayloadRequest) Bind(c *gin.Context) error {
	// binding payload
	if err := c.BindJSON(&r); err != nil {
		return errors.New("failed to bind payload")
	}

	// validate payment settings
	if err := r.Payment.validate(); err != nil {
		return err
	}

	// validate payment settings
	if err := r.Privacy.validate(); err != nil {
		return err
	}

	return nil
}

func (s *PaymentSetting) validate() error {
	// platforms
	platforms := []string{"discord", "telegram", "google"}
	if !sliceutils.Contains(platforms, s.DefaultReceiverPlatform) {
		return fmt.Errorf("default_receiver_platform: invalid value. Available values: %s", strings.Join(platforms, ","))
	}

	// token_priorities
	for _, tkn := range s.TokenPriorities {
		if !uuidutils.IsValid(tkn) {
			return errors.New("token_priorities: must contain valid token IDs")
		}
	}

	// default_message_settings
	if s.DefaultMessageEnable != nil && *s.DefaultMessageEnable {
		if len(s.DefaultMessageSettings) == 0 {
			return errors.New("default_message_settings: must specify at least one")
		}
	}

	// validate action of message settings
	msgActions := []string{"tip", "airdrop", "paylink", "payme"}
	hasInvalidMsgAction := sliceutils.Some(s.DefaultMessageSettings, func(s DefaultMessageSetting) bool {
		return !sliceutils.Contains(msgActions, s.Action)
	})
	if hasInvalidMsgAction {
		return fmt.Errorf("default_message_settings.action: invalid value. Available values: %s", strings.Join(msgActions, ","))
	}

	// check duplicated actions
	inputMsgActions := sliceutils.Map(s.DefaultMessageSettings, func(s DefaultMessageSetting) string {
		return s.Action
	})
	if duplications := sliceutils.FindDuplications(inputMsgActions); len(duplications) > 0 {
		return fmt.Errorf("default_message_settings.action: duplicated values (%s)", strings.Join(duplications, ","))
	}

	// tx_limit_settings
	if s.TxLimitEnable != nil && *s.TxLimitEnable {
		if len(s.TxLimitSettings) == 0 {
			return errors.New("tx_limit_settings: must specify at least one")
		}
	}

	// validate action of tx limit settings
	limitActions := []string{"tip", "airdrop", "paylink", "payme", "withdraw"}
	hasInvalidLimitAction := sliceutils.Some(s.TxLimitSettings, func(s TxLimitSetting) bool {
		return !sliceutils.Contains(limitActions, s.Action)
	})
	if hasInvalidLimitAction {
		return fmt.Errorf("tx_limit_settings.action: invalid value. Available values: %s", strings.Join(limitActions, ","))
	}

	hasInvalidLimitBoundary := sliceutils.Some(s.TxLimitSettings, func(s TxLimitSetting) bool {
		return s.Min < 0 || s.Min >= s.Max
	})
	if hasInvalidLimitBoundary {
		return fmt.Errorf("tx_limit_settings: min has to be a whole number and smaller than max")
	}

	// check duplicated actions
	inputLimitActions := sliceutils.Map(s.TxLimitSettings, func(s TxLimitSetting) string {
		return s.Action
	})
	if duplications := sliceutils.FindDuplications(inputLimitActions); len(duplications) > 0 {
		return fmt.Errorf("tx_limit_settings.action: duplicated values (%s)", strings.Join(duplications, ","))
	}

	return nil
}

func (s *PrivacySetting) validate() error {
	// destination wallet
	if err := s.DestinationWalet.validate(); err != nil {
		return fmt.Errorf("tx.%v", err)
	}

	return nil
}

func (s *BasePrivacySetting) validate() error {
	// target_group
	targetGroups := sliceutils.Map([]model.TargetGroup{model.TargetGroupAll, model.TargetGroupFriends, model.TargetGroupReceivers}, func(g model.TargetGroup) string {
		return string(g)
	})
	if !sliceutils.Contains(targetGroups, s.TargetGroup) {
		return fmt.Errorf("target_group: invalid value. Available values: %s", strings.Join(targetGroups, ","))
	}

	return nil
}

type UpdateNotificationSettingPayloadRequest struct {
	Enable    *bool           `json:"enable" binding:"required"`
	Platforms []string        `json:"platforms" binding:"required"`
	Flags     map[string]bool `json:"flags" binding:"required"`
}

func (r *UpdateNotificationSettingPayloadRequest) Bind(c *gin.Context) error {
	platforms := []string{"discord", "telegram", "web"}
	for _, p := range r.Platforms {
		if !sliceutils.Contains(platforms, p) {
			return fmt.Errorf("platforms: invalid value. Available values: %s", strings.Join(platforms, ","))
		}
	}

	return nil
}
