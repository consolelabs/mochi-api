package covalent

import "time"

type pagination struct {
	HasMore    bool `json:"has_more"`
	PageNumber int  `json:"page_number"`
	PageSize   int  `json:"page_size"`
	TotalCount *int `json:"total_count"`
}

type decoded struct {
	Name      string          `json:"name"`
	Signature string          `json:"signature"`
	Params    []decodedParams `json:"params"`
}

type decodedParams struct {
	Value interface{} `json:"value"`
	// Name    string      `json:"name"`
	// Type    string      `json:"type"`
	// Indexed bool        `json:"indexed"`
	// Decoded bool        `json:"decoded"`
}

type LogEvent struct {
	SenderContractDecimals     int     `json:"sender_contract_decimals"`
	SenderName                 string  `json:"sender_name"`
	SenderContractTickerSymbol string  `json:"sender_contract_ticker_symbol"`
	SenderAddress              string  `json:"sender_address"`
	Decoded                    decoded `json:"decoded"`
	// BlockSignedAt              time.Time   `json:"block_signed_at"`
	// BlockHeight                int         `json:"block_height"`
	// TxOffset                   int         `json:"tx_offset"`
	// LogOffset                  int         `json:"log_offset"`
	// TxHash                     string      `json:"tx_hash"`
	// RawLogTopics               []string    `json:"raw_log_topics"`
	// SenderAddressLabel         interface{} `json:"sender_address_label"`
	// SenderLogoURL              string      `json:"sender_logo_url"`
	// RawLogData                 string      `json:"raw_log_data"`
}

type TransactionItemData struct {
	BlockSignedAt time.Time  `json:"block_signed_at"`
	TxHash        string     `json:"tx_hash"`
	Successful    bool       `json:"successful"`
	FromAddress   string     `json:"from_address"`
	ToAddress     string     `json:"to_address"`
	Value         string     `json:"value"`
	ValueQuote    float64    `json:"value_quote"`
	LogEvents     []LogEvent `json:"log_events"`
	// BlockHeight      int         `json:"block_height"`
	// TxOffset         int         `json:"tx_offset"`
	// FromAddressLabel interface{} `json:"from_address_label"`
	// ToAddressLabel   interface{} `json:"to_address_label"`
	// GasOffered       int         `json:"gas_offered"`
	// GasSpent         int         `json:"gas_spent"`
	// GasPrice         int64       `json:"gas_price"`
	// FeesPaid         string      `json:"fees_paid"`
	// GasQuote         float64     `json:"gas_quote"`
	// GasQuoteRate     float64     `json:"gas_quote_rate"`
	// TokenSymbol      string      `json:"-"`
	// TokenContract    string      `json:"-"`
	// Amount           float64     `json:"-"`
}

type GetTransactionsByAddressData struct {
	// Address       string                `json:"address"`
	// UpdatedAt     time.Time             `json:"updated_at"`
	// NextUpdateAt  time.Time             `json:"next_update_at"`
	// QuoteCurrency string                `json:"quote_currency"`
	// ChainID       int                   `json:"chain_id"`
	Items []TransactionItemData `json:"items"`
	// Pagination    pagination            `json:"pagination"`
}

type GetTransactionsResponse struct {
	Data         GetTransactionsByAddressData `json:"data"`
	Error        bool                         `json:"error"`
	ErrorMessage string                       `json:"error_message"`
	ErrorCode    int                          `json:"error_code"`
}

type portfolioItem struct {
	ContractDecimals     int         `json:"contract_decimals"`
	ContractName         string      `json:"contract_name"`
	ContractTickerSymbol string      `json:"contract_ticker_symbol"`
	ContractAddress      string      `json:"contract_address"`
	SupportsErc          interface{} `json:"supports_erc"`
	LogoURL              string      `json:"logo_url"`
	Holdings             []struct {
		Timestamp time.Time `json:"timestamp"`
		QuoteRate float64   `json:"quote_rate"`
		Open      struct {
			Balance string  `json:"balance"`
			Quote   float64 `json:"quote"`
		} `json:"open"`
		High struct {
			Balance string  `json:"balance"`
			Quote   float64 `json:"quote"`
		} `json:"high"`
		Low struct {
			Balance string  `json:"balance"`
			Quote   float64 `json:"quote"`
		} `json:"low"`
		Close struct {
			Balance string  `json:"balance"`
			Quote   float64 `json:"quote"`
		} `json:"close"`
	} `json:"holdings"`
}

type GetHistoricalPortfolioData struct {
	Address       string          `json:"address"`
	UpdatedAt     time.Time       `json:"updated_at"`
	NextUpdateAt  time.Time       `json:"next_update_at"`
	QuoteCurrency string          `json:"quote_currency"`
	ChainID       int             `json:"chain_id"`
	ChainName     string          `json:"chain_name"`
	Items         []portfolioItem `json:"items"`
	Pagination    interface{}     `json:"pagination"`
}

type GetHistoricalPortfolioResponse struct {
	Data         *GetHistoricalPortfolioData `json:"data"`
	Error        bool                        `json:"error"`
	ErrorMessage string                      `json:"error_message"`
	ErrorCode    int                         `json:"error_code"`
}

type TokenBalanceItem struct {
	ContractDecimals     int     `json:"contract_decimals"`
	ContractName         string  `json:"contract_name"`
	ContractTickerSymbol string  `json:"contract_ticker_symbol"`
	ContractAddress      string  `json:"contract_address"`
	NativeToken          bool    `json:"native_token"`
	Type                 string  `json:"type"`
	Balance              string  `json:"balance"`
	Quote                float64 `json:"quote"`
	QuoteRate            float64 `json:"quote_rate"`
	// LogoURL              string      `json:"logo_url"`
	// NftData              interface{} `json:"nft_data"`
}

type GetTokenBalancesData struct {
	ChainName  string             `json:"chain_name"`
	Items      []TokenBalanceItem `json:"items"`
	Pagination interface{}        `json:"pagination"`
	// Address       string             `json:"address"`
	// UpdatedAt     time.Time          `json:"updated_at"`
	// NextUpdateAt  time.Time          `json:"next_update_at"`
	// QuoteCurrency string             `json:"quote_currency"`
	// ChainID       int                `json:"chain_id"`
}

type GetTokenBalancesResponse struct {
	Data         *GetTokenBalancesData `json:"data"`
	Error        bool                  `json:"error"`
	ErrorMessage string                `json:"error_message"`
	ErrorCode    int                   `json:"error_code"`
}
