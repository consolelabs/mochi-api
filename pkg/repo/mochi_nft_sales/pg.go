package mochinftsales

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) CreateOne(message *request.TwitterSalesMessage) error {
	return pg.db.Table("mochi_nft_sales").Create(&message).Error
}
func (pg *pg) GetAllUnnotified() ([]model.TwitterSalesMessage, error) {
	messages := []model.TwitterSalesMessage{}
	err := pg.db.Table("mochi_nft_sales").Where("is_notified_twitter=false").Find(&messages)
	if err.Error != nil {
		return nil, err.Error
	}
	return messages, nil
}

func (pg *pg) DeleteOne(message *model.TwitterSalesMessage) error {
	return pg.db.Table("mochi_nft_sales").Delete(message).Error
}
