package response

import "github.com/defipod/mochi/pkg/model"

type GetSaleTwitterConfigResponse struct {
	Data []model.SaleBotTwitterConfig `json:"data,omitempty"`
}

type CreateTwitterSaleConfigResponse struct {
	Data *model.SaleBotTwitterConfig `json:"data,omitempty"`
}
