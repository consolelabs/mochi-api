package coingeckosupportedtokens

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Get(q GetQuery) ([]model.CoingeckoSupportedTokens, error) {
	var tokens []model.CoingeckoSupportedTokens
	db := pg.db.Table("coingecko_supported_tokens")
	if q.ID != "" {
		db = db.Where("id = ?", q.ID)
	}
	if q.Symbol != "" {
		db = db.Where("symbol = ?", q.Symbol)
	}
	return tokens, db.Find(&tokens).Error
}

func (pg *pg) Upsert(item *model.CoingeckoSupportedTokens) (int64, error) {
	tx := pg.db.Begin()
	tx = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(item)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return tx.RowsAffected, tx.Commit().Error
}
