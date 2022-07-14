package request

type NftSale struct {
	TokenId           string `json:"token_id"`
	CollectionAddress string `json:"collection_address"`
	CollectionName    string `json:"collection_name"`
	CollectionImage   string `json:"collection_image"`
	TokenName         string `json:"token_name"`
	TokenImage        string `json:"token_image"`
	Rarity            string `json:"rarity"`
	Rank              uint64 `json:"rank"`
	Marketplace       string `json:"marketplace"`
	Transaction       string `json:"transaction"`
	From              string `json:"from"`
	To                string `json:"to"`
	Price             string `json:"price"`
	Bought            string `json:"bought"`
	Sold              string `json:"sold"`
	Hodl              string `json:"hodl"`
	Gain              string `json:"gain"`
	Pnl               string `json:"pnl"`
	SubPnl            string `json:"sub_pnl"`
	PaymentToken      string `json:"payment_token"`
}

type NftSalesRequest struct {
	NftSales []NftSale `json:"nft_sales"`
}
