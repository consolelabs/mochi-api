package response

import "github.com/defipod/mochi/pkg/model"

type ProductBotCommand struct {
	Data []model.ProductBotCommand `json:"data"`
}

type ProductChangelogs struct {
	Data       []model.ProductChangelogs `json:"data"`
	Pagination PaginationResponse        `json:"pagination"`
}

type CreateProductChangelogsViewed struct {
	Data model.ProductChangelogView `json:"data"`
}

type GetProductChangelogsViewed struct {
	Data []model.ProductChangelogView `json:"data"`
}
