package chain

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetAll() ([]model.Chain, error)
	GetByID(id int) (model.Chain, error)
	GetByShortName(shortName string) (*model.Chain, error)
}
