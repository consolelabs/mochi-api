package tokeninfo

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/plugin/dbresolver"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetOne(id string) (*model.TokenInfo, error) {
	info := &model.TokenInfo{}
	return info, pg.db.Where("id = ?", id).Clauses(dbresolver.Write).First(info).Error
}

func (pg *pg) Upsert(item *model.TokenInfo) (int64, error) {
	tx := pg.db.Begin()
	tx = tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(item)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return tx.RowsAffected, tx.Commit().Error
}
