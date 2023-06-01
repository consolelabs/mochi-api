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

type SuiTransactionBlock struct {
	Jsonrpc string                    `json:"jsonrpc"`
	Result  SuiTransactionBlockResult `json:"result"`
	ID      string                    `json:"id"`
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

type SuiTransactionBlockResult struct {
	Data []struct {
		Digest         string              `json:"digest"`
		ObjectChanges  []SuiObjectChanges  `json:"objectChanges"`
		BalanceChanges []SuiBalanceChanges `json:"balanceChanges"`
		TimestampMs    string              `json:"timestampMs"`
		Checkpoint     string              `json:"checkpoint"`
	} `json:"data"`
	NextCursor  string `json:"nextCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type SuiObjectChanges struct {
	Type   string `json:"type"`
	Sender string `json:"sender"`
	Owner  struct {
		AddressOwner string `json:"AddressOwner"`
	} `json:"owner"`
	ObjectType      string `json:"objectType"`
	ObjectID        string `json:"objectId"`
	Version         string `json:"version"`
	PreviousVersion string `json:"previousVersion,omitempty"`
	Digest          string `json:"digest"`
}

type SuiBalanceChanges struct {
	Owner struct {
		AddressOwner string `json:"AddressOwner"`
	} `json:"owner"`
	CoinType string `json:"coinType"`
	Amount   string `json:"amount"`
}
