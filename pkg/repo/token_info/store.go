package tokeninfo

import "github.com/defipod/mochi/pkg/model"

type ListQuery struct {
	ID     string
	Symbol string
}

type Store interface {
	GetOne(token string) (*model.TokenInfo, error)
	Upsert(token model.TokenInfo) (rowsAffected int64, err error)
}
