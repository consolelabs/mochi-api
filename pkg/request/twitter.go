package request

type TwitterSalesMessage struct {
	TokenName         string `json:"token_name"`
	CollectionName    string `json:"collection_name"`
	Price             string `json:"price"`
	SellerAddress     string `json:"seller_address"`
	BuyerAddress      string `json:"buyer_address"`
	Marketplace       string `json:"marketplace"`
	MarketplaceURL    string `json:"marketplace_url"`
	Image             string `json:"image"`
	TxURL             string `json:"tx_url"`
	CollectionAddress string `json:"collection_address"`
	TokenID           string `json:"token_id"`
}

type TwitterPost struct {
	TwitterID     string `json:"twitter_id"`
	TwitterHandle string `json:"twitter_handle"`
	TweetID       string `json:"tweet_id"`
	GuildID       string `json:"guild_id"`
}
