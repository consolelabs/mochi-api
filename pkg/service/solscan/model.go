package solscan

type TransactionListItem struct {
	BlockTime          int      `json:"blockTime"`
	Slot               int      `json:"slot"`
	TxHash             string   `json:"txHash"`
	Fee                int      `json:"fee"`
	Status             string   `json:"status"`
	Lamport            int      `json:"lamport"`
	Signer             []string `json:"signer"`
	IncludeSPLTransfer bool     `json:"includeSPLTransfer,omitempty"`
	ParsedInstruction  []struct {
		ProgramID string `json:"programId"`
		Type      string `json:"type"`
		Program   string `json:"program,omitempty"`
	} `json:"parsedInstruction"`
}

type TokenAmountItem struct {
	TokenAddress string `json:"tokenAddress"`
	TokenAmount  struct {
		Amount         string  `json:"amount"`
		Decimals       int     `json:"decimals"`
		UIAmount       float64 `json:"uiAmount"`
		UIAmountString string  `json:"uiAmountString"`
	} `json:"tokenAmount"`
	TokenAccount string `json:"tokenAccount"`
	TokenName    string `json:"tokenName"`
	TokenIcon    string `json:"tokenIcon"`
	RentEpoch    int    `json:"rentEpoch"`
	Lamports     int    `json:"lamports"`
	TokenSymbol  string `json:"tokenSymbol,omitempty"`
}

type TokenMetadataResponse struct {
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	Website string `json:"website"`
	Twitter string `json:"twitter"`
	Tag     []struct {
		Name        string      `json:"name"`
		Description interface{} `json:"description"`
	} `json:"tag"`
	Decimals       int     `json:"decimals"`
	CoingeckoID    string  `json:"coingeckoId"`
	Price          float64 `json:"price"`
	TokenAuthority string  `json:"tokenAuthority"`
	Supply         string  `json:"supply"`
	Type           string  `json:"type"`
}

type TransactionDetailsResponse struct {
	BlockTime    int      `json:"blockTime"`
	Slot         int      `json:"slot"`
	TxHash       string   `json:"txHash"`
	Fee          int      `json:"fee"`
	Status       string   `json:"status"`
	Lamport      int      `json:"lamport"`
	Signer       []string `json:"signer"`
	LogMessage   []string `json:"logMessage"`
	InputAccount []struct {
		Account     string `json:"account"`
		Signer      bool   `json:"signer"`
		Writable    bool   `json:"writable"`
		PreBalance  int    `json:"preBalance"`
		PostBalance int    `json:"postBalance"`
	} `json:"inputAccount"`
	RecentBlockhash string `json:"recentBlockhash"`
	Version         string `json:"version"`
	TokenTransfers  []struct {
		Source           string `json:"source"`
		Destination      string `json:"destination"`
		SourceOwner      string `json:"source_owner"`
		DestinationOwner string `json:"destination_owner"`
		Amount           string `json:"amount"`
		Token            struct {
			Address  string `json:"address"`
			Symbol   string `json:"symbol"`
			Icon     string `json:"icon"`
			Decimals int    `json:"decimals"`
		} `json:"token"`
		Type string `json:"type"`
	} `json:"tokenTransfers"`
	SolTransfers []struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
		Amount      int    `json:"amount"`
	} `json:"solTransfers"`
	// SerumTransactions   []interface{} `json:"serumTransactions"`
	// RaydiumTransactions []interface{} `json:"raydiumTransactions"`
	// UnknownTransfers    []interface{} `json:"unknownTransfers"`
}
