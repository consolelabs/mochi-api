package userpaymentsetting

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

func (p *pg) FirstOrCreate(s model.UserPaymentSetting) (*model.UserPaymentSetting, error) {
	return &s, p.db.Where("profile_id = ?", s.ProfileId).FirstOrCreate(&s).Error
}

func (p *pg) Update(s *model.UserPaymentSetting) error {
	return p.db.Where("profile_id = ?", s.ProfileId).Save(s).Error
}
