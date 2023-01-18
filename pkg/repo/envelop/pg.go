package envelop

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

func (pg *pg) Create(envelop *model.Envelop) error {
	return pg.db.Create(envelop).Error
}

func (pg *pg) GetUserStreak(userID string) (model *model.UserEnvelopStreak, err error) {
	return model, pg.db.Raw("SELECT user_id, total_envelop FROM user_envelops WHERE user_id = ?", userID).First(&model).Error
}
