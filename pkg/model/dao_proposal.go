package model

import "time"

type DaoProposal struct {
	Id                       int64     `json:"id"`
	GuildId                  string    `json:"guild_id"`
	GuildConfigDaoProposalId int64     `json:"guild_config_dao_proposal_id"`
	VotingChannelId          string    `json:"voting_channel_id"`
	DiscussionChannelId      string    `json:"discussion_channel_id"`
	CreatorId                string    `json:"creator_id"`
	Title                    string    `json:"title"`
	Description              string    `json:"description"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	ClosedAt                 time.Time `json:"closed_at"`
}

func (DaoProposal) TableName() string {
	return "dao_proposal"
}
