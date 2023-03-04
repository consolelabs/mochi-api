package offchaintipbotuserbalancesnapshot

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateBatch([]model.OffchainTipBotUserBalanceSnapshot) error
}
