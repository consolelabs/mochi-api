package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type GetGuildConfigDaoProposalData struct {
	Id                 int64                       `json:"id"`
	GuildId            string                      `json:"guild_id"`
	ProposalChannelId  string                      `json:"proposal_channel_id"`
	GuidelineChannelId string                      `json:"guideline_channel_id"`
	Authority          model.ProposalAuthorityType `json:"authority"`
	Type               *model.ProposalVotingType   `json:"type"`
	RequiredAmount     string                      `json:"required_amount" gorm:"type:numeric"`
	ChainID            int64                       `json:"chain_id"`
	Chain              string                      `json:"chain"`
	Symbol             string                      `json:"symbol"`
	Address            string                      `json:"address"`
	CreatedAt          time.Time                   `json:"created_at"`
	UpdatedAt          time.Time                   `json:"updated_at"`
}

type GetGuildConfigDaoProposal struct {
	Data *GetGuildConfigDaoProposalData `json:"data"`
}

type GetAllDaoProposals struct {
	Data *[]model.DaoProposal `json:"data"`
}
type GetAllDaoProposalVotes struct {
	Proposal *GetDaoProposalData `json:"proposal"`
	Votes    *[]model.DaoVote    `json:"votes"`
}

type GetDaoProposalData struct {
	Id                       int64                         `json:"id"`
	GuildId                  string                        `json:"guild_id"`
	GuildConfigDaoProposalId int64                         `json:"guild_config_dao_proposal_id"`
	VotingChannelId          string                        `json:"voting_channel_id"`
	DiscussionChannelId      string                        `json:"discussion_channel_id"`
	CreatorId                string                        `json:"creator_id"`
	Title                    string                        `json:"title"`
	Description              string                        `json:"description"`
	Points                   *[]model.DaoProposalVoteCount `json:"points"`
	CreatedAt                *time.Time                    `json:"created_at"`
	UpdatedAt                *time.Time                    `json:"updated_at"`
	ClosedAt                 *time.Time                    `json:"closed_at"`
}

type GetVote struct {
	Data *model.DaoVote `json:"data"`
}

type UpdateVote struct {
	Data *model.DaoVote `json:"data"`
}

type TokenHolderStatusData struct {
	IsWalletConnected bool                          `json:"is_wallet_connected"`
	UserHoldingAmount *string                       `json:"user_holding_amount,omitempty"`
	IsQualified       *bool                         `json:"is_qualified,omitempty"`
	GuildConfig       *model.GuildConfigDaoProposal `json:"guild_config,omitempty"`
	VoteConfig        *model.DaoProposalVoteOption  `json:"vote_config,omitempty"`
}

type TokenHolderStatus struct {
	Data *TokenHolderStatusData `json:"data"`
}
