package chain

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

func (pg *pg) List(discussionID *int64) ([]model.CommonwealthDiscussionSubscription, error) {
	var subs []model.CommonwealthDiscussionSubscription
	query := pg.db
	if discussionID != nil {
		query = query.Where("discussion_id = ?", *discussionID)
	}
	return subs, query.Find(&subs).Error
}

func (pg *pg) Create(sub *model.CommonwealthDiscussionSubscription) error {
	return pg.db.Create(&sub).Error
}
