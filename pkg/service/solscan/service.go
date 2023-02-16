package solscan

import "github.com/defipod/mochi/pkg/model"

type Service interface {
	GetSolanaCollection(address string) (*model.SolanaCollectionMetadata, error)
	GetTransactions(address string) ([]TransactionListItem, error)
	GetTokenBalances(address string) ([]TokenAmountItem, error)
	GetTokenMetadata(tokenAddress string) (*TokenMetadataResponse, error)
	GetTxDetails(signature string) (*TransactionDetailsResponse, error)
}
