package coingecko

import (
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetHistoricalMarketData(req *request.GetMarketChartRequest) (*response.CoinPriceHistoryResponse, error, int)
	GetMarketData(coinID string, currency string) (*response.MarketDataResponse, error, int)
	GetCoin(coinID string) (*response.GetCoinResponse, error, int)
	GetCoinPrice(coinIDs []string, currency string) (map[string]float64, error)
	SearchCoins(query string) ([]response.SearchedCoin, error, int)
}
