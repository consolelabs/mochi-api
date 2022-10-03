package request

type AddNftWatchlistRequest struct {
	UserID            string `json:"user_id"`
	GuildID           string `json:"guild_id"`
	CollectionSymbol  string `json:"collection_symbol"`
	CollectionAddress string `json:"collection_address"`
	Chain             string `json:"chain"`
}

type DeleteNftWatchlistRequest struct {
	UserID string `json:"user_id" form:"user_id" binding:"required"`
	Symbol string `json:"symbol" form:"symbol" binding:"required"`
}

type GetNftWatchlistRequest struct {
	UserID string `json:"user_id" form:"user_id" binding:"required"`
	Page   int    `json:"page" form:"page"`
	Size   int    `json:"size" form:"size"`
}

type GetNFTCollectionTickersRequest struct {
	CollectionAddress string `json:"collection_address" form:"collection_address" binding:"required"`
}
