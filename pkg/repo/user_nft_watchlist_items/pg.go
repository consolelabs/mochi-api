package usernftwatchlistitem

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

func (pg *pg) List(q UserNftWatchlistQuery) ([]model.UserNftWatchlistItem, int64, error) {
	var items []model.UserNftWatchlistItem
	var total int64
	db := pg.db.Table("user_nft_watchlist_items")
	if q.ProfileID != "" {
		db = db.Where("profile_id = ?", q.ProfileID)
	}
	if q.Symbol != "" {
		db = db.Where("symbol ILIKE ?", q.Symbol)
	}
	if q.CollectionAddress != "" {
		db = db.Where("collection_address = ?", q.CollectionAddress)
	}
	if q.ChainID != "" {
		db = db.Where("chain_id = ?", q.ChainID)
	}

	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return items, total, db.Find(&items).Error
}

func (pg *pg) Create(item *model.UserNftWatchlistItem) error {
	return pg.db.Create(item).Error
}

func (pg *pg) Delete(profileID, symbol string) (int64, error) {
	tx := pg.db.Where("profile_id = ? AND symbol ILIKE ?", profileID, symbol).Delete(&model.UserNftWatchlistItem{})
	return tx.RowsAffected, tx.Error
}
