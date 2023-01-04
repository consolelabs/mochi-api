package response

import "github.com/defipod/mochi/pkg/model"

type CreateDaoProposalResponse struct {
	Data model.DaoProposal `json:"data"`
}
