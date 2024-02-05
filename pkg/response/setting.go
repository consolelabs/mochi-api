package response

import "github.com/defipod/mochi/pkg/model"

type GeneralSettingData struct {
	Payment *model.UserPaymentSetting `json:"payment"`
	Privacy *model.UserPrivacySetting `json:"privacy"`
}

type UserGeneralSettingResponse struct {
	Data GeneralSettingData `json:"data"`
}

func ToUserGeneralSettingResponse(payment *model.UserPaymentSetting, privacy *model.UserPrivacySetting) *UserGeneralSettingResponse {
	return &UserGeneralSettingResponse{
		Data: GeneralSettingData{
			Payment: payment,
			Privacy: privacy,
		},
	}
}

type UserPaymentSettingResponse struct {
	Data *model.UserPaymentSetting `json:"data"`
}

type UserPrivacySettingResponse struct {
	Data *model.UserPrivacySetting `json:"data"`
}

type UserNotificationSettingResponse struct {
	Data *model.UserNotificationSetting `json:"data"`
}

type GetUserTipMessageResponse struct {
	Data *UserTipMessageData `json:"data"`
}

type UserTipMessageData struct {
	Message string `json:"message"`
}

func ToUserTipMessageResponse(msg string) *GetUserTipMessageResponse {
	return &GetUserTipMessageResponse{
		Data: &UserTipMessageData{
			Message: msg,
		},
	}
}
