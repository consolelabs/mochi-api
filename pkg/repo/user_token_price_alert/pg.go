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

func (pg *pg) GetOne(q UserTokenPriceAlertQuery) (model.UserTokenPriceAlert, error) {
	var item model.UserTokenPriceAlert
	db := pg.db.Table("user_token_price_alerts")
	return item, db.Where("user_discord_id = ? AND symbol = ? AND price = ?", q.UserDiscordID, q.Symbol, q.Price).First(&item).Error
}

func (pg *pg) List(q UserTokenPriceAlertQuery) ([]model.UserTokenPriceAlert, int64, error) {
	var items []model.UserTokenPriceAlert
	var total int64
	db := pg.db.Table("user_token_price_alerts")
	if q.UserDiscordID != "" {
		db = db.Where("user_discord_id = ?", q.UserDiscordID)
	}
	if q.Symbol != "" {
		db = db.Where("symbol = ?", q.Symbol)
	}
	if q.Price != 0 {
		db = db.Where("price = ?", q.Price)
	}
	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return items, total, db.Find(&items).Error
}

func (pg *pg) Delete(userID, symbol string, price float64) (int64, error) {
	tx := pg.db.Where("user_discord_id = ? AND symbol ILIKE ? AND price = ?", userID, symbol, price).Delete(&model.UserTokenPriceAlert{})
	return tx.RowsAffected, tx.Error
}

func (pg *pg) Update(item *model.UserTokenPriceAlert) error {
	return pg.db.Where("user_discord_id = ? AND symbol = ?", item.UserDiscordID, item.Symbol).Save(item).Error
}

func (pg *pg) FetchListSymbol() ([]string, error) {
	var symbols []string
	return symbols, pg.db.Table("user_token_price_alerts").Distinct().Pluck("symbol", &symbols).Error
}
