package response

import "github.com/defipod/mochi/pkg/model"

type ProductBotCommand struct {
	Data []model.ProductBotCommand `json:"data"`
}

type ProductChangelogs struct {
	Data []model.ProductChangelogs `json:"data"`
}
