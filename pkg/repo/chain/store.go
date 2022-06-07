package chain

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetAll() ([]model.Chain, error)
}
