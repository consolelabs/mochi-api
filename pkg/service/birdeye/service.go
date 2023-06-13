package birdeye

type Service interface {
	GetTokenPrice(address string) (*TokenPrice, error)
}
