package covalent

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetHistoricalTokenPrices(chainID int, currency string, address string) (*response.HistoricalTokenPricesResponse, error, int)
	GetTransactionsByAddress(chainID int, address string, size int, retry int) (*GetTransactionsResponse, error)
	GetTokenBalances(chainID int, address string, retry int) (*GetTokenBalancesResponse, error)
	GetSolanaTokenBalances(chainName string, address string, retry int) (*GetTokenBalancesResponse, error)
}
