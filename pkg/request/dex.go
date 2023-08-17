package request

type SearchDexPairRequest struct {
	Query string `form:"query"`
}

type GetDexPairRequest struct {
	Network string `form:"network"`
	Address string `form:"address"`
}
