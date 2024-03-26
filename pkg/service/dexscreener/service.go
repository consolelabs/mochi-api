package dexscreener

type Service interface {
	Search(query string) ([]Pair, error)
	GetByChainAndPairAddress(network, pairAddr string) (*Pair, error)
	GetByTokenAddress(tokenAddr string) ([]Pair, error)
}
