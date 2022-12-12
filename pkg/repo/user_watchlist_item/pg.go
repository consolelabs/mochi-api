package userwatchlistitem

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

func (pg *pg) List(q UserWatchlistQuery) ([]model.UserWatchlistItem, int64, error) {
	var items []model.UserWatchlistItem
	var total int64
	db := pg.db.Table("user_watchlist_items")
	if q.UserID != "" {
		db = db.Where("user_id = ?", q.UserID)
	}
	if q.CoinGeckoID != "" {
		db = db.Where("coin_gecko_id = ?", q.CoinGeckoID)
	}
	if q.Symbol != "" {
		db = db.Where("symbol ILIKE ?", q.Symbol)
	}
	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return items, total, db.Find(&items).Error
}

func (pg *pg) Create(item *model.UserWatchlistItem) error {
	return pg.db.Create(item).Error
}

func (pg *pg) Delete(userID, symbol string) (int64, error) {
	tx := pg.db.Where("user_id = ? AND symbol ILIKE ?", userID, symbol).Delete(&model.UserWatchlistItem{})
	return tx.RowsAffected, tx.Error
}
