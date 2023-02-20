package response

import "github.com/defipod/mochi/pkg/model"

type CreateDaoProposalResponse struct {
	Data model.DaoProposal `json:"data"`
}

type ProposalCount struct {
	GuildId string `json:"guild_id"`
	Count   int64  `json:"count"`
}
type GuildProposalUsageResponse struct {
	Pagination PaginationResponse        `json:"metadata"`
	Data       *[]GuildProposalUsageData `json:"data"`
}

type GuildProposalUsageData struct {
	GuildId       string `json:"guild_id"`
	ProposalCount int64  `json:"proposal_count"`
	IsActive      bool   `json:"is_active"`
}

type DaoTrackerSpaceCountResponse struct {
	Pagination PaginationResponse          `json:"metadata"`
	Data       *[]DaoTrackerSpaceCountData `json:"data"`
}
type DaoTrackerSpaceCountData struct {
	Space  string `json:"space"`
	Count  int64  `json:"count"`
	Source string `json:"source"`
}
