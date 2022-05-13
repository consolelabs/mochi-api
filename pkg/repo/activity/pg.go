package activity

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetOne(id int) (*model.Activity, error) {
	activity := &model.Activity{}
	return activity, pg.db.First(&activity, id).Error
}

func (pg *pg) GetDefaultActivities() ([]model.Activity, error) {
	var activities []model.Activity
	return activities, pg.db.Where("guild_default = TRUE").Find(&activities).Error
}
