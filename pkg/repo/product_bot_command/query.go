package productbotcommand

type ListQuery struct {
	Code  string
	Scope int64 `json:"scope,omitempty"`
}
