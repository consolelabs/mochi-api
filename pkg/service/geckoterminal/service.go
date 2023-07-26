package geckoterminal

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	Search(query string) (*Search, error)
	GetPool(network, pool string) (*response.GetCoinResponse, error)
	GetPoolInfo(network, pool string) (*Pool, error)
}
