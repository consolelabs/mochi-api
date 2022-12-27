package offchain_tip_bot_deposit_log

import (
	"github.com/google/uuid"

	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	GetLatestByChainIDAndContract(chainID, contractAddress string) (*model.OffchainTipBotDepositLog, error)
	GetByID(chainID uuid.UUID, txHash string) (*model.OffchainTipBotDepositLog, error)
	CreateMany([]model.OffchainTipBotDepositLog) error
}
