package covalent

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetHistoricalTokenPrices(chainID int, currency string, address string) (*response.HistoricalTokenPricesResponse, error, int)
	GetTransactionsByAddress(chainID int, address string) (res *GetTransactionsResponse, err error)
}
