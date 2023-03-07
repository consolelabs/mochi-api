package coingeckotokenalias

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(alias string) (*model.CoingeckoTokenAlias, error)
	SetAlias(alias string) error
}
