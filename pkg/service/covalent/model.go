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
	Name    string `json:"name"`
	Type    string `json:"type"`
	Indexed bool   `json:"indexed"`
	Decoded bool   `json:"decoded"`
	Value   string `json:"value"`
}

type logEvent struct {
	BlockSignedAt              time.Time   `json:"block_signed_at"`
	BlockHeight                int         `json:"block_height"`
	TxOffset                   int         `json:"tx_offset"`
	LogOffset                  int         `json:"log_offset"`
	TxHash                     string      `json:"tx_hash"`
	RawLogTopics               []string    `json:"raw_log_topics"`
	SenderContractDecimals     int         `json:"sender_contract_decimals"`
	SenderName                 string      `json:"sender_name"`
	SenderContractTickerSymbol string      `json:"sender_contract_ticker_symbol"`
	SenderAddress              string      `json:"sender_address"`
	SenderAddressLabel         interface{} `json:"sender_address_label"`
	SenderLogoURL              string      `json:"sender_logo_url"`
	RawLogData                 string      `json:"raw_log_data"`
	Decoded                    decoded     `json:"decoded"`
}

type TransactionItemData struct {
	BlockSignedAt    time.Time   `json:"block_signed_at"`
	BlockHeight      int         `json:"block_height"`
	TxHash           string      `json:"tx_hash"`
	TxOffset         int         `json:"tx_offset"`
	Successful       bool        `json:"successful"`
	FromAddress      string      `json:"from_address"`
	FromAddressLabel interface{} `json:"from_address_label"`
	ToAddress        string      `json:"to_address"`
	ToAddressLabel   interface{} `json:"to_address_label"`
	Value            string      `json:"value"`
	ValueQuote       float64     `json:"value_quote"`
	GasOffered       int         `json:"gas_offered"`
	GasSpent         int         `json:"gas_spent"`
	GasPrice         int64       `json:"gas_price"`
	FeesPaid         string      `json:"fees_paid"`
	GasQuote         float64     `json:"gas_quote"`
	GasQuoteRate     float64     `json:"gas_quote_rate"`
	LogEvents        []logEvent  `json:"log_events"`
	TokenSymbol      string      `json:"-"`
	TokenContract    string      `json:"-"`
	Amount           float64     `json:"-"`
}

type GetTransactionsByAddressData struct {
	Address       string                `json:"address"`
	UpdatedAt     time.Time             `json:"updated_at"`
	NextUpdateAt  time.Time             `json:"next_update_at"`
	QuoteCurrency string                `json:"quote_currency"`
	ChainID       int                   `json:"chain_id"`
	Items         []TransactionItemData `json:"items"`
	Pagination    pagination            `json:"pagination"`
}

type GetTransactionsResponse struct {
	Data         GetTransactionsByAddressData `json:"data"`
	Error        bool                         `json:"error"`
	ErrorMessage *string                      `json:"error_message"`
	ErrorCode    *int                         `json:"error_code"`
}
