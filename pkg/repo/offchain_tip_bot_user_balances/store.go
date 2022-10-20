package offchain_tip_bot_user_balances

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetUserBalances(userID string) ([]model.OffchainTipBotUserBalance, error)
}
