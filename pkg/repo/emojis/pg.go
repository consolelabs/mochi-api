package emojis

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

func (pg *pg) ListEmojis(codes []string) (model []*model.Emojis, err error) {
	db := pg.db
	if len(codes) > 0 {
		db = db.Where("code in (?)", codes)
	}
	return model, db.Find(&model).Error
}
