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
	Hodl              string `json:"hodl"`
	Pnl               string `json:"pnl"`
	SubPnl            string `json:"sub_pnl"`
}

type TwitterPost struct {
	TwitterID     string `json:"twitter_id"`
	TwitterHandle string `json:"twitter_handle"`
	TweetID       string `json:"tweet_id"`
	GuildID       string `json:"guild_id"`
	Content       string `json:"content"`
}

type AddToTwitterBlackListRequest struct {
	GuildID         string `json:"guild_id"`
	TwitterID       string `json:"twitter_id"`
	TwitterUsername string `json:"twitter_username"`
	CreatedBy       string `json:"created_by"`
}

type DeleteFromTwitterBlackListRequest struct {
	GuildID   string `json:"guild_id" form:"guild_id" binding:"required"`
	TwitterID string `json:"twitter_id" form:"twitter_id" binding:"required"`
}

type GetTwitterLeaderboardRequest struct {
	GuildID string `json:"guild_id" form:"guild_id" binding:"required"`
	PaginationRequest
}
