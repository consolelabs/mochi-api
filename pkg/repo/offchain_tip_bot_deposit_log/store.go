package offchain_tip_bot_deposit_log

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	GetOne(chainID, txHash string) (*model.OffchainTipBotDepositLog, error)
	CreateMany([]model.OffchainTipBotDepositLog) error
}
