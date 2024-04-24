package binance

import (
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetExchangeInfo(symbol string) (*response.GetExchangeInfoResponse, error, int)
	GetTickerPrice(symbol string) (*response.GetTickerPriceResponse, error, int)
	GetKlinesBySymbol(symbol string) ([]response.GetKlinesDataResponse, error, int)
	GetAvgPriceBySymbol(symbol string) (*response.GetAvgPriceBySymbolResponse, error, int)
	GetApiKeyPermission(apiKey, apiSecret string) (*response.BinanceApiKeyPermissionResponse, error)
	GetUserAsset(apiKey, apiSecret string) ([]response.BinanceUserAssetResponse, error)
	GetFundingAsset(apiKey, apiSecret string) ([]response.BinanceUserAssetResponse, error)
	GetStakingProductPosition(apiKey, apiSecret string) ([]response.BinanceStakingProductPosition, error)
	GetLendingAccount(apiKey, apiSecret string) (*response.BinanceLendingAccount, error)
	GetSimpleEarn(apiKey, apiSecret string) (*response.BinanceSimpleEarnAccount, error)
	GetFutureAccountBalance(apiKey, apiSecret string) ([]response.BinanceFutureAccountBalance, error)
	GetFutureAccount(apiKey, apiSecret string) (*response.BinanceFutureAccount, error)
	GetFutureAccountInfo(apiKey, apiSecret string) ([]response.BinanceFuturePositionInfo, error)
	GetPrice(symbol string) (*response.BinanceApiTickerPriceResponse, error)
	GetSpotTransactions(apiKey, apiSecret, startTime, endTime string) ([]response.BinanceSpotTransaction, error)
}
