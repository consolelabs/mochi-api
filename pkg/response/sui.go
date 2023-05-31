package response

type SuiAllBalance struct {
	Jsonrpc string                `json:"jsonrpc"`
	Result  []SuiAllBalanceResult `json:"result"`
	ID      int                   `json:"id"`
}

type SuiCoinMetadata struct {
	Jsonrpc string                `json:"jsonrpc"`
	Result  SuiCoinMetadataResult `json:"result"`
	ID      int                   `json:"id"`
}

type SuiAllBalanceResult struct {
	CoinType        string `json:"coinType"`
	CoinObjectCount int    `json:"coinObjectCount"`
	TotalBalance    string `json:"totalBalance"`
	LockedBalance   struct {
	} `json:"lockedBalance"`
}

type SuiCoinMetadataResult struct {
	Decimals    int     `json:"decimals"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Description string  `json:"description"`
	IconURL     *string `json:"iconUrl"`
	ID          string  `json:"id"`
}
