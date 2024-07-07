package request

type AddNftWatchlistRequest struct {
	WatchlistBaseRequest
	GuildID           string `json:"guild_id"`
	CollectionSymbol  string `json:"collection_symbol"`
	CollectionAddress string `json:"collection_address"`
	Chain             string `json:"chain"`
}

type DeleteNftWatchlistRequest struct {
	ProfileID string
	Symbol    string `json:"symbol" binding:"required"`
}

type ListTrackingNftsRequest struct {
	ProfileID string
	Page      int `form:"page,default=0"`
	Size      int `form:"size,default=16"`
}

type GetNFTCollectionTickersRequest struct {
	CollectionAddress string `json:"collection_address" form:"collection_address" binding:"required"`
}

type GetNFTTokenTickersRequest struct {
	CollectionAddress string `json:"collection_address" form:"collection_address" binding:"required"`
	TokenID           string `json:"token_id" form:"token_id" binding:"required"`
}

type GetProfileNFTsRequest struct {
	ProfileID         string `json:"profile_id" form:"profile_id"`
	CollectionAddress string `json:"address" uri:"address" binding:"required"`
}
