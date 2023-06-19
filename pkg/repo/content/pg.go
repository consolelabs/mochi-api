package content

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

func (pg *pg) GetContentByType(contentType string) (content *model.ProductMetadataCopy, err error) {
	return content, pg.db.Where("type = ?", contentType).First(&content).Error
}
