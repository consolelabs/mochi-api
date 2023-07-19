package dao_vote

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id int64) (*model.DaoVote, error)
	GetByUserAndProposalID(proposalId int64, userId string) (model *model.DaoVote, err error)
	CreateDaoVote(vote *model.DaoVote) error
	GetByUserId(userId string) (*[]model.DaoVote, error)
	GetByProposalId(proposalId int64) (*[]model.DaoVote, error)
	Update(vote *model.DaoVote) (*model.DaoVote, error)
}
