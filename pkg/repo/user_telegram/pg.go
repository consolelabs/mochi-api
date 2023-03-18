package usertelegram

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
func (pg *pg) GetByUsername(username string) (model *model.UserTelegram, err error) {
	return model, pg.db.Where("username = ?", username).First(&model).Error
}
