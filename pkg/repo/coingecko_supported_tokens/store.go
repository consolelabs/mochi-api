package coingeckosupportedtokens

import "github.com/defipod/mochi/pkg/model"

type GetQuery struct {
	ID     string
	Symbol string
}

type Store interface {
	Get(q GetQuery) ([]model.CoingeckoSupportedTokens, error)
	Upsert(token *model.CoingeckoSupportedTokens) (rowsAffected int64, err error)
}
