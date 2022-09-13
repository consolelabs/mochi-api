package coingeckosupportedtokens

import "github.com/defipod/mochi/pkg/model"

type ListQuery struct {
	ID     string
	Symbol string
}

type Store interface {
	GetOne(id string) (*model.CoingeckoSupportedTokens, error)
	List(q ListQuery) ([]model.CoingeckoSupportedTokens, error)
	Upsert(token *model.CoingeckoSupportedTokens) (rowsAffected int64, err error)
}
