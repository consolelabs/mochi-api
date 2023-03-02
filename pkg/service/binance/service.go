package binance

import (
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetExchangeInfo(symbol string) (*response.GetExchangeInfoResponse, error, int)
	GetTickerPrice(symbol string) (*response.GetTickerPriceResponse, error, int)
	GetKlinesBySymbol(symbol string) ([]response.GetKlinesDataResponse, error, int)
	GetAvgPriceBySymbol(symbol string) (*response.GetAvgPriceBySymbolResponse, error, int)
}
