package geckoterminal

type Service interface {
	Search(query string) (*Search, error)
	GetPool(network, pool string) (*Pool, error)
	// ScrapePool(network, pool string) (*ScrapePool, error)
}
