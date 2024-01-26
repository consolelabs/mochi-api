package request

type ProductBotCommandRequest struct {
	Scope int64  `form:"scope,omitempty"`
	Code  string `form:"code,omitempty"`
}

type ProductChangelogsRequest struct {
	Product string `form:"product"`
	Size    int64  `form:"size,default=10"`
	Page    int64  `form:"page,default=0"`
}

type CreateProductChangelogsViewRequest struct {
	Key           string `json:"key"`
	ChangelogName string `json:"changelog_name"`
}

type GetProductChangelogsViewRequest struct {
	Key           string `form:"key"`
	ChangelogName string `form:"changelog_name"`
}

type ProductChangelogSnapshotRequest struct {
	ChangelogName string `json:"changelog_name"`
	IsPublic      bool   `json:"is_public"`
}

type GetProductHashtagRequest struct {
	Alias string `form:"alias"`
}

type GetProductThemeRequest struct {
}

type GetListEmojiRequest struct {
	Size       int64  `form:"size" default:"10"`
	Page       int64  `form:"page" default:"0"`
	Codes      string `form:"codes"`
	IsQueryAll bool
	ListCode   []string
}
