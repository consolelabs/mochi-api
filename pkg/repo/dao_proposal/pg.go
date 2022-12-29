package dao_proposal

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id int64) (*model.DaoProposal, error)
}
