package geckoterminal

type Service interface {
	Search(query string) (*GeckoTerminalSearch, error)
	GetPool(network, pool string) (*GeckoTerminalGetPool, error)
	ScrapePool(network, pool string) (*ScrapePool, error)
}
