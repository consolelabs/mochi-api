package covalent

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetHistoricalTokenPrices(chainID int, currency string, address string) (*response.HistoricalTokenPricesResponse, error, int)
}
