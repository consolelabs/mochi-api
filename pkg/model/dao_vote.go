package model

import "time"

type DaoVote struct {
	Id         int64     `json:"id"`
	ProposalId int64     `json:"proposal_id"`
	UserId     string    `json:"user_id"`
	Choice     string    `json:"choice"`
	Point      float64   `json:"point"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (DaoVote) TableName() string {
	return "dao_vote"
}
