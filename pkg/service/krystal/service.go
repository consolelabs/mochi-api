package krystal

type Service interface {
	GetBalanceTokenByAddress(address string) (*BalanceTokenResponse, error)
	GetEarningOptions(platforms, chainIds, types, statuses, address string) (*GetEarningOptionsResponse, error)
	BuildStakeTx(req BuildStakeTxReq) (*BuildTxResp, error)
	BuildUnstakeTx(req BuildUnstakeTxReq) (*BuildTxResp, error)
}
