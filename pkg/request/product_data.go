package request

type ProductBotCommandRequest struct {
	Scope int64  `form:"scope,omitempty"`
	Code  string `form:"code,omitempty"`
}

type ProductChangelogsRequest struct {
	Product string `form:"product"`
	Size    string `form:"size"`
}
