package binancespottransaction

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

func (pg *pg) Create(tx *model.BinanceSpotTransaction) error {
	return pg.db.Create(&tx).Error
}

func (pg *pg) List(q ListQuery) ([]model.BinanceSpotTransaction, error) {
	var txs []model.BinanceSpotTransaction
	db := pg.db
	if q.ProfileId != "" {
		db = db.Where("profile_id = ?", q.ProfileId)
	}
	if q.Status != "" {
		db = db.Where("status = ?", q.Status)
	}
	if q.Offset > 0 {
		db = db.Offset(q.Offset)
	}
	if q.Limit > 0 {
		db = db.Limit(q.Limit)
	}
	return txs, db.Order("created_at DESC").Find(&txs).Error
}

func (pg *pg) GetUserAverageCost(profileId string) ([]model.BinanceAssetAverageCost, error) {
	avgCost := make([]model.BinanceAssetAverageCost, 0)
	db := pg.db
	rows, err := db.Raw(`
	SELECT
		profile_id, symbol, sum(executed_qty::decimal * price_in_usd::decimal) / sum(executed_qty::decimal) as average_cost
	FROM
		binance_spot_transactions
	WHERE
		profile_id = ?
		AND TYPE = 'LIMIT'
		AND side = 'BUY'
		AND status = 'FILLED'
		GROUP BY symbol, profile_id;
	`, profileId).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var avgCostItem model.BinanceAssetAverageCost
		if err := rows.Scan(&avgCostItem.ProfileId, &avgCostItem.Symbol, &avgCostItem.AverageCost); err != nil {
			return nil, err
		}
		avgCost = append(avgCost, avgCostItem)
	}

	return avgCost, nil
}
