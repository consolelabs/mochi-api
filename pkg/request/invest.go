package request

type GetInvestListRequest struct {
	ChainIds  string `form:"chainIds"`
	Platforms string `form:"platforms"`
	Types     string `form:"types"`
	Address   string `form:"address"`
	Status    string `form:"status"`
}

type OnchainInvestStakeDataRequest struct {
	ChainID      int    `form:"chainID" binding:"required"`
	Type         string `form:"type" binding:"required"`
	Platform     string `form:"platform" binding:"required"`
	TokenAddress string `form:"tokenAddress" binding:"required"`
	TokenAmount  string `form:"tokenAmount" binding:"required"`
	UserAddress  string `form:"userAddress" binding:"required"`
}

type OnchainInvestUnstakeDataRequest struct {
	ChainID      int    `form:"chainID" binding:"required"`
	Type         string `form:"type" binding:"required"`
	Platform     string `form:"platform" binding:"required"`
	TokenAddress string `form:"tokenAddress" binding:"required"`
	TokenAmount  string `form:"tokenAmount" binding:"required"`
	UserAddress  string `form:"userAddress" binding:"required"`
}
