package usernotificationsetting

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	FirstOrCreate(model.UserNotificationSetting) (*model.UserNotificationSetting, error)
	Update(*model.UserNotificationSetting) error
}
