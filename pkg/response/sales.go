package response

type NftSales struct {
	Platform             string  `json:"platform"`
	NftName              string  `json:"nft_name"`
	NftStatus            string  `json:"nft_status"`
	NftCollectionAddress string  `json:"nft_collection_address"`
	NftPrice             float64 `json:"nft_price"`
	NftPriceToken        string  `json:"nft_price_token"`
	Buyer                string  `json:"buyer"`
	Seller               string  `json:"seller"`
}

type NftSalesResponse struct {
	Data []NftSales `json:"data"`
}
