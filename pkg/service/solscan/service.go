package solscan

import (
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetCollectionBySolscanId(id string) (*response.CollectionDataResponse, error)
	GetNftTokenFromCollection(id, page string) (*response.NftTokenDataResponse, error)
	GetTransactions(address string) ([]TransactionListItem, error)
	GetTokenBalances(address string) ([]TokenAmountItem, error)
	GetTokenMetadata(tokenAddress string) (*TokenMetadataResponse, error)
	GetTxDetails(signature string) (*TransactionDetailsResponse, error)
}
