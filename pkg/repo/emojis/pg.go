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

func (pg *pg) ListEmojis(q Query) (model []*model.ProductMetadataEmojis, total int64, err error) {
	db := pg.db.Model(model).Order("id ASC")
	if len(q.Codes) > 0 {
		db = db.Where("code in (?)", q.Codes)
	}
	db.Count(&total)
	if q.Size > 0 {
		db = db.Limit(q.Size)
	}
	if q.Page > 0 {
		db = db.Offset(q.Page * q.Size)
	}
	return model, total, db.Find(&model).Error
}
