package notificationflag

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

func (p *pg) List(ListQuery) (flags []model.NotificationFlag, err error) {
	db := p.db
	return flags, db.Find(&flags).Error
}
