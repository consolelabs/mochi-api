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

func (pg *pg) GetByAlias(alias string) (p *model.ProductHashtag, err error) {
	db := pg.db.Table("product_hashtags")

	return p, db.Where("lower(?)=ANY(alias)", alias).First(&p).Error
}

func (pg *pg) GetBySlug(slug string) (p *model.ProductHashtag, err error) {
	db := pg.db.Table("product_hashtags")

	return p, db.Where("lower(slug)=?", slug).First(&p).Error
}
