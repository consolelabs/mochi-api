package offchain_tip_bot_user_balances

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/google/uuid"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetUserBalances(userID string) ([]model.OffchainTipBotUserBalance, error) {
	var balance []model.OffchainTipBotUserBalance
	return balance, pg.db.Preload("Token").Find(&balance, "user_id = ?", userID).Error
}

func (pg *pg) GetUserBalanceByTokenID(userID string, tokenID uuid.UUID) (*model.OffchainTipBotUserBalance, error) {
	var balance model.OffchainTipBotUserBalance
	return &balance, pg.db.Preload("Token").First(&balance, "user_id = ? AND token_id = ?", userID, tokenID).Error
}

func (pg *pg) UpdateUserBalance(balance *model.OffchainTipBotUserBalance) error {
	return pg.db.Table("offchain_tip_bot_user_balances").Where("user_id = ? and token_id = ?", balance.UserID, balance.TokenID).Update("amount", balance.Amount).Error
}

func (pg *pg) UpdateListUserBalances(listUserID []string, tokenID uuid.UUID, amount float64) error {
	return pg.db.Table("offchain_tip_bot_user_balances").Where("user_id IN ? and token_id = ?", listUserID, tokenID).UpdateColumn("amount", gorm.Expr("amount + ?", amount)).Error
}

func (pg *pg) CreateIfNotExists(model *model.OffchainTipBotUserBalance) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "token_id"}},
		DoNothing: true,
	}).Create(model).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) SumAmountByTokenId() ([]response.TotalOffchainBalancesInDB, error) {
	rows, err := pg.db.Raw(
		`
		SELECT
			sum(bals.amount) AS total,
			bals.token_id,
			tokens.token_symbol,
			tokens.coin_gecko_id
		FROM
			offchain_tip_bot_user_balances AS bals
			JOIN offchain_tip_bot_tokens AS tokens ON bals.token_id = tokens.id
		GROUP BY
			bals.token_id,
			tokens.token_symbol,
			tokens.coin_gecko_id
		`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var totalOffchainBalances []response.TotalOffchainBalancesInDB
	for rows.Next() {
		t := response.TotalOffchainBalancesInDB{}
		rows.Scan(
			&t.Total,
			&t.TokenId,
			&t.TokenSymbol,
			&t.CoinGeckoId,
		)
		totalOffchainBalances = append(totalOffchainBalances, t)
	}
	return totalOffchainBalances, nil
}
