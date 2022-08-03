package request

type TwitterSalesMessage struct {
	TokenName      string `json:"token_name"`
	CollectionName string `json:"collection_name"`
	Price          string `json:"price"`
	SellerAddress  string `json:"seller_address"`
	BuyerAddress   string `json:"buyer_address"`
	Marketplace    string `json:"marketplace"`
	MarketplaceURL string `json:"marketplace_url"`
	Image          string `json:"image"`
}
