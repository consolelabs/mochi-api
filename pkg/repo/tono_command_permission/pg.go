package tonocommandpermission

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

func (pg *pg) List(q ListQuery) (cmds []model.TonoCommandPermission, err error) {
	db := pg.db
	if q.Code != "" {
		db = db.Where("lower(code) = lower(?)", q.Code)
	}

	return cmds, db.Find(&cmds).Error
}
