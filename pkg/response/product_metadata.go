package response

import "github.com/defipod/mochi/pkg/model"

type ProductBotCommand struct {
	Data []model.ProductBotCommand `json:"data"`
}
