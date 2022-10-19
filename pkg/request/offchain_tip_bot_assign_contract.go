package request

type CreateAssignContract struct {
	UserID      string `json:"user_id"`
	TokenSymbol string `json:"token_symbol"`
}
