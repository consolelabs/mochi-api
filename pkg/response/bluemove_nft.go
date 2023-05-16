package response

type BluemoveAptosMeta struct {
	Pagination BluemoveAptosPagination `json:"pagination"`
}

type BluemoveAptosPagination struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}

// bluemove collection detail data structure
type BluemoveCollectionsResponse struct {
	Data []BluemoveCollectionDetail `json:"data"`
	Meta BluemoveAptosMeta          `json:"meta"`
}

type BluemoveCollectionDetailResponse struct {
	Data BluemoveCollectionDetail `json:"data"`
}

type BluemoveCollectionDetail struct {
	Id         int64                      `json:"id"`
	Attributes BluemoveCollectionProperty `json:"attributes"`
}

type BluemoveCollectionProperty struct {
	Name        string                    `json:"name"`
	Slug        string                    `json:"slug"`
	Description string                    `json:"description"`
	Twitter     string                    `json:"twitter"`
	Discord     string                    `json:"discord"`
	Telegram    string                    `json:"telegram"`
	Website     string                    `json:"website"`
	Creator     string                    `json:"creator"`
	Items       string                    `json:"items"`
	Owners      string                    `json:"owners"`
	Uri         string                    `json:"uri"`
	Type        string                    `json:"type"`
	Metadata    map[string]ValueTraitType `json:"metadata"`
}

type ValueTraitType struct {
	Values map[string]int64 `json:"values"`
}
