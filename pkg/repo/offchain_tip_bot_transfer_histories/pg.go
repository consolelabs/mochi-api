package offchain_tip_bot_transfer_histories

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) CreateTransferHistories(transferHistories []model.OffchainTipBotTransferHistory) ([]model.OffchainTipBotTransferHistory, error) {
	return transferHistories, pg.db.Create(transferHistories).Error
}

func (pg *pg) GetByUserDiscordId(userDiscordId string) (transferHistories []model.OffchainTipBotTransferHistory, err error) {
	return transferHistories, pg.db.Where("sender_id = ?", userDiscordId).Find(&transferHistories).Error
}

func (pg *pg) TotalFeeFromWithdraw() ([]response.TotalFeeWithdraw, error) {
	rows, err := pg.db.Raw(
		`
		SELECT
			sum(transf.fee_amount) as total_fee,
			transf.token
		FROM
			offchain_tip_bot_transfer_histories AS transf
		WHERE
			status = 'success'
			AND action = 'withdraw'
		GROUP BY
			transf.token;
		`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var totalFeeWithdraw []response.TotalFeeWithdraw
	for rows.Next() {
		t := response.TotalFeeWithdraw{}
		rows.Scan(
			&t.TotalFee,
			&t.Symbol,
		)
		totalFeeWithdraw = append(totalFeeWithdraw, t)
	}
	return totalFeeWithdraw, nil
}
