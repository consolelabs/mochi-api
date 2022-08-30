package response

import "github.com/defipod/mochi/pkg/model"

type GetListAllChainsResponse struct {
	Data []model.Chain `json:"data"`
}
