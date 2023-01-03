package dao_proposal

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id int64) (*model.DaoProposal, error)
	GetAllByCreatorId(userId string) (*[]model.DaoProposal, error)
	GetByCreatorIdAndProposalId(proposal int64, userId string) (models []model.DaoProposalWithView, err error)
}
