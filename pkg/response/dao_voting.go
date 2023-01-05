package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type GetGuildConfigDaoProposal struct {
	Data *model.GuildConfigDaoProposal `json:"data"`
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
