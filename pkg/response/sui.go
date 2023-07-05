package response

type SuiAllBalance struct {
	Result []SuiAllBalanceResult `json:"result"`
}

type SuiCoinMetadata struct {
	Result SuiCoinMetadataResult `json:"result"`
}

type SuiTransactionBlock struct {
	Result SuiTransactionBlockResult `json:"result"`
}

type SuiAllBalanceResult struct {
	CoinType     string `json:"coinType"`
	TotalBalance string `json:"totalBalance"`
	// CoinObjectCount int    `json:"coinObjectCount"`
	// LockedBalance   struct {
	// 	} `json:"lockedBalance"`
}

type SuiCoinMetadataResult struct {
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	// Description string  `json:"description"`
	// IconURL     *string `json:"iconUrl"`
	// ID          string  `json:"id"`
}

type SuiTransactionBlockResult struct {
	Data []struct {
		Digest         string              `json:"digest"`
		ObjectChanges  []SuiObjectChanges  `json:"objectChanges"`
		BalanceChanges []SuiBalanceChanges `json:"balanceChanges"`
		TimestampMs    string              `json:"timestampMs"`
		// Checkpoint     string              `json:"checkpoint"`
	} `json:"data"`
	// NextCursor  string `json:"nextCursor"`
	// HasNextPage bool   `json:"hasNextPage"`
}

type SuiObjectChanges struct {
	Type   string `json:"type"`
	Sender string `json:"sender"`
	Owner  struct {
		AddressOwner string `json:"AddressOwner"`
	} `json:"owner"`
	// ObjectType      string `json:"objectType"`
	// ObjectID        string `json:"objectId"`
	// Version         string `json:"version"`
	// PreviousVersion string `json:"previousVersion,omitempty"`
	// Digest          string `json:"digest"`
}

type SuiBalanceChanges struct {
	Owner struct {
		AddressOwner string `json:"AddressOwner"`
	} `json:"owner"`
	CoinType string `json:"coinType"`
	Amount   string `json:"amount"`
}
