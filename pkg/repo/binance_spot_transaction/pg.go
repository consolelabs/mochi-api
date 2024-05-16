package binancespottransaction

import (
	"strconv"

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
		profile_id,
		symbol,
		sum(
			CASE WHEN side = 'BUY' THEN
				executed_qty::decimal
			ELSE
				- executed_qty::decimal
			END) AS total_amount,
		sum(
			CASE WHEN side = 'BUY' THEN
				executed_qty::decimal * price::decimal
			ELSE
				- executed_qty::decimal * price::decimal
			END) /
		sum(
			CASE WHEN side = 'BUY' THEN
				executed_qty::decimal
			ELSE
				- executed_qty::decimal
			END) AS avg_cost
	FROM
		binance_spot_transactions
	WHERE
		profile_id = ?
		AND TYPE = 'LIMIT'
		AND status = 'FILLED'
	GROUP BY
		symbol, profile_id;
	`, profileId).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var avgCostItem model.BinanceAssetAverageCost
		if err := rows.Scan(&avgCostItem.ProfileId, &avgCostItem.Symbol, &avgCostItem.TotalAmount, &avgCostItem.AverageCost); err != nil {
			return nil, err
		}
		// Incase of invalid total amount, we just skip the pair
		// there is some case leading to invalid total amount
		// tx is not enough
		// user deposit or withdraw asset from cex
		amount, err := strconv.ParseFloat(avgCostItem.TotalAmount, 64)
		if err != nil {
			continue
		}
		if amount > 0 {
			avgCost = append(avgCost, avgCostItem)
		}
	}

	return avgCost, nil
}

func (pg *pg) Update(tx *model.BinanceSpotTransaction) error {
	return pg.db.Save(&tx).Error
}
