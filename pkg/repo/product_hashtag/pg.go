package producthashtag

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

func (pg *pg) GetByAlias(alias string) (p *model.ProductHashtagAlias, err error) {
	db := pg.db.Table("product_hashtag_alias")

	return p, db.Where("alias = ?", alias).Preload("ProductHashtag").First(&p).Error
}
