package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateDaoVote(req request.CreateDaoVoteRequest) error {
	return e.repo.DaoVote.CreateDaoVote(&model.DaoVote{
		UserId:     req.UserID,
		ProposalId: req.ProposalID,
		Choice:     req.Choice,
		Point:      1,
	})
}
