package krystal

type Service interface {
	GetBalanceTokenByAddress(address string) (*BalanceTokenResponse, error)
}
