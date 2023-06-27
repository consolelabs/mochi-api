package krystal

type Service interface {
	GetBalanceTokenByAddress(address string, chainIDs []int) (*BalanceTokenResponse, error)
}
