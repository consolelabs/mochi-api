package model

import "time"

type DaoProposalVoteOption struct {
	Id             int64          `json:"id"`
	ProposalId     int64          `json:"proposal_id"`
	VoteOptionId   *int64         `json:"vote_option_id"`
	Address        string         `json:"address"`
	ChainId        int64          `json:"chain_id"`
	Symbol         string         `json:"symbol"`
	RequiredAmount string         `json:"required_amount" gorm:"type:numeric"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	VoteOption     *DaoVoteOption `json:"vote_option,omitempty"`
}

func (DaoProposalVoteOption) TableName() string {
	return "dao_proposal_vote_option"
}
