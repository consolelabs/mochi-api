package offchain_tip_bot_user_balances

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/google/uuid"
)

type Store interface {
	GetUserBalances(userID string) ([]model.OffchainTipBotUserBalance, error)
	GetUserBalanceByTokenID(userID string, tokenID uuid.UUID) (*model.OffchainTipBotUserBalance, error)
	UpdateUserBalance(balance *model.OffchainTipBotUserBalance) error
	UpdateListUserBalances(listUserID []string, tokenID uuid.UUID, amount float64) error
	CreateIfNotExists(model *model.OffchainTipBotUserBalance) error
	SumAmountByTokenId() ([]response.TotalOffchainBalancesInDB, error)
}
