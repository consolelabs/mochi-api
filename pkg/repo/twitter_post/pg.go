package twitterpost

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}
func (pg *pg) CreateOne(post *model.TwitterPost) error {
	return pg.db.Table("twitter_posts").Create(&post).Error
}
