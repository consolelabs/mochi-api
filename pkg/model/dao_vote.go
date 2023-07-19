package model

import (
	"time"

	"github.com/defipod/mochi/pkg/model/errors"
)

type VoteChoice string

const (
	Yes     VoteChoice = "Yes"
	No      VoteChoice = "No"
	Abstain VoteChoice = "Abstain"
)

func (c VoteChoice) IsValid() error {
	switch c {
	case Yes, No, Abstain:
		return nil
	}
	return errors.ErrInvalidVoteChoice
}

type DaoVote struct {
	Id         int64      `json:"id"`
	ProposalId int64      `json:"proposal_id"`
	UserId     string     `json:"user_id"`
	Choice     VoteChoice `json:"choice"`
	Point      float64    `json:"point"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func (DaoVote) TableName() string {
	return "dao_vote"
}
