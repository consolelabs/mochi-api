package usertokenpricealert

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

func (pg *pg) Create(item *model.UserTokenPriceAlert) error {
	return pg.db.Create(item).Error
}

func (pg *pg) List(q UserTokenPriceAlertQuery) ([]model.UserTokenPriceAlert, int64, error) {
	var items []model.UserTokenPriceAlert
	var total int64
	db := pg.db.Table("user_token_price_alerts")
	if q.UserID != "" {
		db = db.Where("user_id = ?", q.UserID)
	}
	if q.CoincapID != "" {
		db = db.Where("coincap_id = ?", q.CoincapID)
	}
	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return items, total, db.Find(&items).Error
}

func (pg *pg) Delete(userID, CoincapID string) (int64, error) {
	tx := pg.db.Where("user_id = ? AND coincap_id ILIKE ?", userID, CoincapID).Delete(&model.UserTokenPriceAlert{})
	return tx.RowsAffected, tx.Error
}

func (pg *pg) Update(item *model.UserTokenPriceAlert) error {
	return pg.db.Where("user_id = ? AND coincap_id = ?", item.UserID, item.CoincapID).Save(item).Error
}
