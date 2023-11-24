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

func (pg *pg) GetByID(id int64) (theme *model.ProductTheme, err error) {
	return theme, pg.db.Table("product_themes").Where("id = ?", id).First(&theme).Error
}
