package request

type GetInvestListRequest struct {
	ChainIds  string `form:"chainIds"`
	Platforms string `form:"platforms"`
	Types     string `form:"types"`
	Address   string `form:"address"`
	Status    string `form:"status"`
}
