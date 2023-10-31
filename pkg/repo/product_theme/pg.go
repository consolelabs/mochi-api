package producttheme

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

func (pg *pg) Get() (themes []model.ProductTheme, err error) {
	return themes, pg.db.Table("product_themes").Find(&themes).Error
}
