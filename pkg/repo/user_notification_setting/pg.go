package usernotificationsetting

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (p *pg) FirstOrCreate(s model.UserNotificationSetting) (*model.UserNotificationSetting, error) {
	return &s, p.db.Where("profile_id = ?", s.ProfileId).FirstOrCreate(&s).Error
}

func (p *pg) Update(s *model.UserNotificationSetting) error {
	return p.db.Where("profile_id = ?", s.ProfileId).Save(s).Error
}
