package coingecko

import (
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetHistoricalMarketData(coinID, currency string, days int) (res *response.HistoricalMarketChartResponse, err error, statusCode int)
	GetCoin(coinID string) (res *response.GetCoinResponse, err error, statusCode int)
	GetCoinPrice(coinIDs []string, currency string) (map[string]float64, error)
	GetHistoryCoinInfo(sourceSymbol string, interval string) (res [][]float64, err error, statusCode int)
	GetCoinsMarketData(ids []string, sparkline bool, page, pageSize string) ([]response.CoinMarketItemData, error, int)
	GetSupportedCoins() (res []response.CoingeckoSupportedTokenResponse, err error, statusCode int)
	GetAssetPlatforms() ([]*response.AssetPlatformResponseData, error)
	GetCoinByContract(platformId, contractAddress string, retry int) (*response.GetCoinByContractResponseData, error)
	GetTrendingSearch() (*response.GetTrendingSearch, error)
	GetTopLoserGainer(req request.TopGainerLoserRequest) (*response.GetTopGainerLoser, error)
	GetHistoricalGlobalMarketChart(days int) (*response.GetHistoricalGlobalMarketResponse, error)
	GetGlobalData() (*response.GetGlobalDataResponse, error)
	SearchCoin(query string) (res *response.SearchCoinResponse, err error, statusCode int)
}
