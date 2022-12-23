package offchain_tip_bot_deposit_log

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type Store interface {
	GetLatestByChainIDAndContract(chainID, contractAddress string) (*model.OffchainTipBotDepositLog, error)
	GetByID(chainID uuid.UUID, txHash string) (*model.OffchainTipBotDepositLog, error)
	CreateMany([]model.OffchainTipBotDepositLog) error
}
