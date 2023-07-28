package response

type GetInvestListResponse struct {
	Data []InvestItem `json:"data"`
}

type InvestItem struct {
	Apy       float64           `json:"apy"`
	Chain     InvestChain       `json:"chain"`
	Platforms []InvestPlatforms `json:"platforms"`
	Token     InvestToken       `json:"token"`
	Tvl       float64           `json:"tvl"`
}

type InvestToken struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
}

type InvestChain struct {
	ID   int    `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

type InvestStatus struct {
	Detail string `json:"detail"`
	Value  string `json:"value"`
}

type InvestPlatforms struct {
	Apy       float64      `json:"apy"`
	Desc      string       `json:"desc"`
	Logo      string       `json:"logo"`
	Name      string       `json:"name"`
	RewardAPY float64      `json:"reward_apy"`
	Status    InvestStatus `json:"status"`
	Tvl       float64      `json:"tvl"`
	Type      string       `json:"type"`
}

type OnchainInvestDataResponse struct {
	Data OnchainInvestData `json:"data"`
}

type OnchainInvestData struct {
	TxObject TxObject `json:"tx_object"`
}

type TxObject struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Data     string `json:"data"`
	GasPrice string `json:"gas_price"`
	Nonce    string `json:"nonce"`
	GasLimit string `json:"gas_limit"`
}
