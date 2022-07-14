package request

type NftSalesRequest struct {
	TokenId           string  `json:"token_id"`
	CollectionAddress string  `json:"collection_address"`
	CollectionName    string  `json:"collection_name"`
	CollectionImage   string  `json:"collection_image"`
	TokenName         string  `json:"token_name"`
	TokenImage        string  `json:"token_image"`
	Rarity            string  `json:"rarity"`
	Rank              uint64  `json:"rank"`
	Marketplace       string  `json:"marketplace"`
	Transaction       string  `json:"transaction"`
	From              string  `json:"from"`
	To                string  `json:"to"`
	Price             Token   `json:"price"`
	Sold              string  `json:"sold"`
	Hodl              int     `json:"hold"`
	Gain              Token   `json:"gain"`
	Pnl               float64 `json:"pnl"`
	SubPnl            float64 `json:"sub_pnl"`
	PaymentToken      string  `json:"payment_token"`
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
