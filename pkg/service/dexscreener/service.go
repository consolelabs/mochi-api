package dexscreener

type Service interface {
	Search(query string) ([]Pair, error)
	Get(network, address string) (*Pair, error)
}
