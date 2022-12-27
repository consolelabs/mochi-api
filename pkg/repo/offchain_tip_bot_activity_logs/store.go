package offchain_tip_bot_activity_logs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateActivityLog(al *model.OffchainTipBotActivityLog) (*model.OffchainTipBotActivityLog, error)
	List(ListQuery) ([]model.OffchainTipBotActivityLog, error)
}
