package response

import "github.com/defipod/mochi/pkg/model"

type GetGuildConfigDaoProposal struct {
	Data *model.GuildConfigDaoProposal `json:"data"`
}
