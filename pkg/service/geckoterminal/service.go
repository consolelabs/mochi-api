package geckoterminal

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	Search(query string) (*Search, error)
	GetPool(network, pool string) (*response.GetCoinResponse, error)
	GetHistoricalMarketData(network, pool string, from, to int64) (*response.HistoricalMarketChartResponse, error)
}
