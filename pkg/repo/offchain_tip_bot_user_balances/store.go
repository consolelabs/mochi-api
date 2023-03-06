package offchain_tip_bot_user_balances

import (
	"github.com/google/uuid"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Store interface {
	GetUserBalances(userID string) ([]model.OffchainTipBotUserBalance, error)
	GetUserBalanceByTokenID(userID string, tokenID uuid.UUID) (*model.OffchainTipBotUserBalance, error)
	SumAmountByTokenId() ([]response.TotalOffchainBalancesInDB, error)

	// Update or create if not exists batch of users' balances
	//
	//	ChangedAmount float64 (for update): amount of increment/decrement to balance after transaction
	// 	e.g. send airdrop of 2 FTM -> ChangedAmount = -2 FTM
	//	Amount float64 (for creation): total amount of balance after transaction
	// 	e.g. current balance = 1 FTM, receive tip of 3 FTM -> Amount = 4 FTM
	UpsertBatch(list []model.OffchainTipBotUserBalance) error
}
