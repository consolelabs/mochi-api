package userwalletwatchlistitem

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) ([]model.UserWalletWatchlistItem, error) {
	var items []model.UserWalletWatchlistItem
	db := pg.db
	if q.UserID != "" {
		db = db.Where("user_id = ?", q.UserID)
	}
	if q.IsOwner != nil {
		db = db.Where("is_owner = ?", *q.IsOwner)
	}
	return items, db.Find(&items).Error
}

func (pg *pg) GetOne(q GetOneQuery) (*model.UserWalletWatchlistItem, error) {
	var item model.UserWalletWatchlistItem

	q.Query = strings.ToLower(q.Query)

	query := pg.db.Where("user_id = ?", q.UserID).Where("lower(address) = ? OR lower(alias) = ?", q.Query, q.Query)

	if q.ForUpdate {
		query = query.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	if err := query.First(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (pg *pg) Upsert(item *model.UserWalletWatchlistItem) error {
	// return pg.db.Create(item).Error
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "address"}},
		DoUpdates: clause.AssignmentColumns([]string{"alias", "type"}),
	}).Create(item).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) Remove(q DeleteQuery) error {
	db := pg.db.Where("user_id = ?", q.UserID)
	if q.Address != "" {
		db = db.Where("address = ?", q.Address)
	}
	if q.Alias != "" {
		db = db.Where("alias = ?", q.Alias)
	}
	return db.Delete(&model.UserWalletWatchlistItem{}).Error
}

func (pg *pg) UpdateOwnerFlag(userID, address string, isOwner bool) error {
	return pg.db.Model(&model.UserWalletWatchlistItem{}).Where("user_id = ? AND address = ?", userID, address).Update("is_owner", isOwner).Error
}

func (pg *pg) Update(item *model.UserWalletWatchlistItem) error {
	return pg.db.
		Model(&model.UserWalletWatchlistItem{}).
		Select("alias"). // This is just a trick to update with zero values
		Where("user_id = ? AND address = ?", item.UserID, item.Address).
		Updates(item).Error
}
