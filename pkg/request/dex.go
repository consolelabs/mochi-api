package request

type SearchDexPairRequest struct {
	Query string `form:"query" binding:"required"`
}

type GetDexPairRequest struct {
	Network string `form:"network"`
	Address string `form:"address"`
}

type SearchDexScreenerPairRequest struct {
	Symbol       string `form:"symbol"`
	TokenAddress string `form:"token_address"`
}
