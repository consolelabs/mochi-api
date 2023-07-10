package usertokenwatchlistitem

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

func (pg *pg) List(q UserWatchlistQuery) ([]model.UserTokenWatchlistItem, int64, error) {
	var items []model.UserTokenWatchlistItem
	var total int64
	db := pg.db.Table("user_token_watchlist_items")
	if q.ProfileID != "" {
		db = db.Where("profile_id = ?", q.ProfileID)
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

func (pg *pg) Create(item *model.UserTokenWatchlistItem) error {
	return pg.db.Create(item).Error
}

func (pg *pg) Delete(profileID, symbol string) (int64, error) {
	tx := pg.db.Where("profile_id = ? AND symbol ILIKE ?", profileID, symbol).Delete(&model.UserTokenWatchlistItem{})
	return tx.RowsAffected, tx.Error
}

func (pg *pg) Count(q CountQuery) (count int64, err error) {
	db := pg.db.Model(&model.UserTokenWatchlistItem{})
	if q.CoingeckoId != "" {
		db = db.Where("coin_gecko_id like ? OR coin_gecko_id like ? OR coin_gecko_id like ?", q.CoingeckoId, q.CoingeckoId+"/%", "%/"+q.CoingeckoId)
	}
	if q.Symbol != "" {
		db = db.Where("symbol like ? OR symbol like ? OR symbol like ?", q.Symbol, q.Symbol+"/%", "%/"+q.Symbol)
	}
	if q.Distinct != "" {
		db = db.Distinct(q.Distinct)
	}
	return count, db.Count(&count).Error
}
