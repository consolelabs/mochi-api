package request

type HandleNftWebhookRequest struct {
	Event             string `json:"event"`
	TokenId           string `json:"token_id"`
	CollectionAddress string `json:"collection_address"`
	Marketplace       string `json:"marketplace"`
	Transaction       string `json:"transaction"`
	From              string `json:"from"`
	To                string `json:"to"`
	Price             Token  `json:"price"`
	Hodl              int    `json:"hold"`
	LastPrice         Token  `json:"last_price"`

	Chain string `json:"chain"`
}
type Token struct {
	Token  TokenInfo `json:"token"`
	Amount string    `json:"amount"`
}
type TokenInfo struct {
	Symbol   string `json:"symbol"`
	IsNative bool   `json:"is_native"`
	Address  string `json:"address"`
	Decimal  int    `json:"decimal"`
}
