package dao_proposal_vote_option

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id int64) (model *model.DaoProposalVoteOption, err error)
	GetOneByProposalID(proposalID int64) (model *model.DaoProposalVoteOption, err error)
	Create(model *model.DaoProposalVoteOption) (*model.DaoProposalVoteOption, error)
	DeleteAllByProposalID(proposalId int64) error
}
