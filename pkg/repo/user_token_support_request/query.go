package user_token_support_request

type ListQuery struct {
	TokenAddress string
	TokenChainID *int
	Status       string
	Offset       int
	Limit        int
}
