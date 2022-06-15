package response

type NFTCollectionResponse struct {
	Tickers         TokenTickers `json:"tickers"`
	FloorPrice      float64      `json:"floor_price"`
	Name            string       `json:"name"`
	ContractAddress string       `json:"contract_address"`
	Chain           string       `json:"chain"`
	Platforms       []string     `json:"platforms"`
}
