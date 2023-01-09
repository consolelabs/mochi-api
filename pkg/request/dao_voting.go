package request

import "github.com/defipod/mochi/pkg/model"

type DeleteGuildConfigDaoProposal struct {
	ID string `json:"id"`
}

type CreateDaoVoteRequest struct {
	UserID     string           `json:"user_id" binding:"required"`
	ProposalID int64            `json:"proposal_id" binding:"required"`
	Choice     model.VoteChoice `json:"choice" binding:"required"`
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

type UpdateDaoVoteRequest struct {
	UserID string           `json:"user_id" binding:"required"`
	Choice model.VoteChoice `json:"choice" binding:"required"`
}

type DAOAction string

const (
	CreateProposal DAOAction = "create_proposal"
	Vote           DAOAction = "vote"
)

type TokenHolderStatusRequest struct {
	Action             DAOAction `json:"action" form:"action" binding:"required,oneof=create_proposal vote"`
	UserID             string    `json:"user_id" form:"user_id" binding:"required"`
	ProposalID         string    `json:"proposal_id" form:"proposal_id"`
	GuildID            string    `json:"guild_id" form:"guild_id"`
	GuidelineChannelID string    `json:"guidline_channel_id" form:"guideline_channel_id"`
}
