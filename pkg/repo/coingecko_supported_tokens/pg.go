package coingeckosupportedtokens

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

func (pg *pg) GetOne(id string) (*model.CoingeckoSupportedTokens, error) {
	token := &model.CoingeckoSupportedTokens{}
	return token, pg.db.Where("id = ?", id).Clauses(dbresolver.Write).First(token).Error
}

func (pg *pg) List(q ListQuery) ([]model.CoingeckoSupportedTokens, error) {
	var tokens []model.CoingeckoSupportedTokens
	db := pg.db.Table("coingecko_supported_tokens")
	if q.ID != "" {
		db = db.Where("id = ?", q.ID)
	}
	if q.Symbol != "" {
		db = db.Where("symbol ILIKE ?", q.Symbol)
	}
	return tokens, db.Find(&tokens).Error
}

func (pg *pg) Upsert(item *model.CoingeckoSupportedTokens) (int64, error) {
	tx := pg.db.Begin()
	tx = tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(item)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return tx.RowsAffected, tx.Commit().Error
}
