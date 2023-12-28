package userprivacysetting

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	FirstOrCreate(model.UserPrivacySetting) (*model.UserPrivacySetting, error)
	Update(*model.UserPrivacySetting) error
}
