package request

type DeleteGuildConfigDaoProposal struct {
	ID string `json:"id"`
}

type CreateDaoVoteRequest struct {
	UserID     string  `json:"user_id" binding:"required"`
	ProposalID int64   `json:"proposal_id" binding:"required"`
	Choice     string  `json:"choice" binding:"required"`
	Point      float64 `json:"point" binding:"required"`
}
