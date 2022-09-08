package userwatchlistitem

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q UserWatchlistQuery) ([]model.UserWatchlistItem, error) {
	var items []model.UserWatchlistItem
	db := pg.db
	if q.UserID != "" {
		db = db.Where("user_id = ?", q.UserID)
	}
	if q.CoinGeckoID != "" {
		db = db.Where("coin_gecko_id = ?", q.CoinGeckoID)
	}
	if q.Symbol != "" {
		db = db.Where("symbol ILIKE ?", q.Symbol)
	}
	db = db.Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return items, db.Find(&items).Error
}

func (pg *pg) Create(item *model.UserWatchlistItem) error {
	return pg.db.Create(item).Error
}

func (pg *pg) Delete(userID, symbol string) error {
	return pg.db.Where("user_id = ? AND symbol ILIKE ?", userID, symbol).Delete(&model.UserWatchlistItem{}).Error
}
