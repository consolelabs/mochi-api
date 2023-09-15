package request

type ProductBotCommandRequest struct {
	Scope int64  `form:"scope,omitempty"`
	Code  string `form:"code,omitempty"`
}

type ProductChangelogsRequest struct {
	Product string `form:"product"`
	Size    string `form:"size"`
}

type CreateProductChangelogsViewRequest struct {
	Key           string `json:"key"`
	ChangelogName string `json:"changelog_name"`
}

type GetProductChangelogsViewRequest struct {
	Key           string `form:"key"`
	ChangelogName string `form:"changelog_name"`
}

type GetProductHashtagRequest struct {
	Alias string `form:"alias"`
}

type GetListEmojiRequest struct {
	Size     int64  `form:"size" default:"10"`
	Page     int64  `form:"page" default:"0"`
	Codes    string `form:"codes"`
	ListCode []string
}
