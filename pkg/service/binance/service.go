package binance

import (
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetExchangeInfo(symbol string) (*response.GetExchangeInfoResponse, error, int)
	GetKlinesBySymbol(symbol string) ([]response.GetKlinesDataResponse, error, int)
}
