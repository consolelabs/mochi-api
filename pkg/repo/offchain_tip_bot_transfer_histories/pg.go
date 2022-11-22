package offchain_tip_bot_transfer_histories

import (
	"reflect"

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

func (pg *pg) GetTransactionsByQuery(senderId, receiverId, token string) ([]response.Transactions, error) {
	var transferHistories []response.Transactions
	rows, err := pg.db.Raw(
		`
		SELECT
			offchain_tip_bot_transfer_histories.id,
			sender_id,
			receiver_id,
			offchain_tip_bot_transfer_histories.guild_id,
			log_id,
			offchain_tip_bot_transfer_histories.status,
			offchain_tip_bot_transfer_histories.amount,
			offchain_tip_bot_transfer_histories.token,
			offchain_tip_bot_transfer_histories.action,
			offchain_tip_bot_transfer_histories.service_fee,
			offchain_tip_bot_transfer_histories.fee_amount,
			offchain_tip_bot_transfer_histories.created_at,
			offchain_tip_bot_transfer_histories.updated_at,
			full_command
		FROM
			offchain_tip_bot_transfer_histories
			LEFT JOIN offchain_tip_bot_activity_logs ON offchain_tip_bot_transfer_histories.log_id = offchain_tip_bot_activity_logs.id
		WHERE (sender_id = ?
			OR receiver_id = ?)
		AND token = ? AND offchain_tip_bot_transfer_histories.status = 'success'
		`, senderId, receiverId, token).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		transaction := response.Transactions{}
		s := reflect.ValueOf(&transaction).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		err := rows.Scan(columns...)
		if err != nil {
			return nil, err
		}
		transferHistories = append(transferHistories, transaction)
	}
	return transferHistories, nil
}
