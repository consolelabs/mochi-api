package offchain_tip_bot_tokens

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetBySymbol(symbol string) (*model.OffchainTipBotToken, error)
	Create(*model.OffchainTipBotToken) error
}
