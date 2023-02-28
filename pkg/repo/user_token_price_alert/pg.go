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

func (pg *pg) Create(item *model.UserTokenPriceAlert) (int, error) {
	return item.ID, pg.db.Create(&item).Error
}

func (pg *pg) GetById(ID int) (model.UserTokenPriceAlert, error) {
	var user model.UserTokenPriceAlert
	return user, pg.db.Table("user_token_price_alerts").First(&user, ID).Error
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
	if q.Value != 0 {
		db = db.Where("value = ?", q.Value)
	}
	db = db.Count(&total).Offset(q.Offset)
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	return items, total, db.Find(&items).Error
}

func (pg *pg) DeleteByID(alertID int) error {
	tx := pg.db.Delete(&model.UserTokenPriceAlert{}, "id = ?", alertID)
	return tx.Error
}

func (pg *pg) Update(item *model.UserTokenPriceAlert) error {
	return pg.db.Where("user_discord_id = ? AND symbol = ? AND value = ? AND price_by_percent = ? ", item.UserDiscordID, item.Symbol, item.Value, item.PriceByPercent).Save(item).Error
}
