package response

type GetUserBalances struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Balances      float64 `json:"balances"`
	BalancesInUSD float64 `json:"balances_in_usd"`
}

type GetUserBalancesResponse struct {
	Data []GetUserBalances `json:"data"`
}
