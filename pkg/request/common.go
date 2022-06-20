package request

type PaginationRequest struct {
	Page int64 `json:"page" form:"page,default=0"`
	Size int64 `json:"size" form:"size,default=10"`
}
