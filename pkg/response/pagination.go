package response

import "github.com/defipod/mochi/pkg/model"

type PaginationResponse struct {
	model.Pagination
	Total int64 `json:"total"`
}
