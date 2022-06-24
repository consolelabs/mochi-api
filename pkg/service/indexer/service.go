package indexer

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	CreateERC721Contract(CreateERC721ContractRequest) error
	GetNFTCollectionTickers(address string) (*response.IndexerNFTCollectionTickersResponse, error)
	GetNFTTradingVolume() ([]response.NFTTradingVolume, error)
	GetNFTCollections(query string) (*response.IndexerGetNFTCollectionsResponse, error)
	GetNFTTokens(address, query string) (*response.IndexerGetNFTTokensResponse, error)
	GetNFTDetail(collectionAddress, tokenID string) (*response.IndexerNFTToken, error)
	GetNftSales() (*response.NftSalesResponse, error)
}
