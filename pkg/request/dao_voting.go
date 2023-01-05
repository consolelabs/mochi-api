package request

type DeleteGuildConfigDaoProposal struct {
	ID string `json:"id"`
}

type CreateDaoVoteRequest struct {
	UserID     string `json:"user_id" binding:"required"`
	ProposalID int64  `json:"proposal_id" binding:"required"`
	Choice     string `json:"choice" binding:"required"`
}

type CreateDaoProposalRequest struct {
	GuildId         string             `json:"guild_id"`
	VotingChannelId string             `json:"voting_channel_id"`
	CreatorId       string             `json:"creator_id"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	VoteOption      *VoteOptionRequest `json:"vote_option"`
}

type VoteOptionRequest struct {
	Id             int64  `json:"id"`
	Address        string `json:"address"`
	ChainId        int64  `json:"chain_id"`
	Symbol         string `json:"symbol"`
	RequiredAmount int64  `json:"required_amount"`
}
