package dao_vote

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id int64) (*model.DaoVote, error)
	CreateDaoVote(vote *model.DaoVote) error
	GetByUserId(userId string) (*[]model.DaoVote, error)
	GetByProposalId(proposalId int64) (*[]model.DaoVote, error)
}
