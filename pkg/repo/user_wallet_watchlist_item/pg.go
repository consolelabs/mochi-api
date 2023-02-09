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

func (pg *pg) List(userID string) ([]model.UserWalletWatchlistItem, error) {
	var items []model.UserWalletWatchlistItem
	return items, pg.db.Where("user_id = ?", userID).Find(&items).Error
}

func (pg *pg) GetOne(q GetOneQuery) (*model.UserWalletWatchlistItem, error) {
	q.Query = strings.ToLower(q.Query)
	var item model.UserWalletWatchlistItem
	return &item, pg.db.Where("user_id = ?", q.UserID).Where("lower(address) = ? OR lower(alias) = ?", q.Query, q.Query).First(&item).Error
}

func (pg *pg) Create(item *model.UserWalletWatchlistItem) error {
	// return pg.db.Create(item).Error
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "address"}},
		DoUpdates: clause.AssignmentColumns([]string{"alias"}),
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
