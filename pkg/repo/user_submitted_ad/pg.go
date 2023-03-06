package usersubmittedad

import (
	"time"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) CreateOne(model model.UserSubmittedAd) (*model.UserSubmittedAd, error) {
	return &model, pg.db.Create(&model).Error
}
func (pg *pg) GetAll() (models []model.UserSubmittedAd, count int64, err error) {
	return models, count, pg.db.Find(&models).Count(&count).Error
}
func (pg *pg) GetById(id int) (*model.UserSubmittedAd, error) {
	var model model.UserSubmittedAd
	return &model, pg.db.Where("id = ?", id).First(&model).Error
}
func (pg *pg) UpdateStatus(id int, newStatus string) error {
	return pg.db.Table("user_submitted_ads").Where("id = ?", id).Updates(map[string]interface{}{"status": newStatus, "updated_at": time.Now().UTC()}).Error
}
func (pg *pg) DeleteOne(id int) error {
	model := model.UserSubmittedAd{}
	return pg.db.Table("user_submitted_ads").Where("id = ?", id).Delete(model).Error
}
