package coingeckoinfo

import "github.com/defipod/mochi/pkg/model"

type ListQuery struct {
	ID     string
	Symbol string
}

type Store interface {
	GetOne(id string) (*model.CoingeckoInfo, error)
	Upsert(token *model.CoingeckoInfo) (rowsAffected int64, err error)
}
