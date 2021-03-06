package indexer

import (
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	CreateERC721Contract(CreateERC721ContractRequest) error
	GetNFTCollectionTickers(address, rawQuery string) (*response.IndexerNFTCollectionTickersResponse, error)
	GetNFTTradingVolume() ([]response.NFTTradingVolume, error)
	GetNFTCollections(query string) (*response.IndexerGetNFTCollectionsResponse, error)
	GetNFTTokens(address, query string) (*response.IndexerGetNFTTokensResponse, error)
	GetNFTDetail(collectionAddress, tokenID string) (*response.IndexerNFTToken, error)
	GetNftSales(addr string, platform string) (*response.NftSalesResponse, error)
	GetNFTContract(addr string) (*response.IndexerContract, error)
	GetNftMetadataAttrIcon() (*response.NftMetadataAttrIconResponse, error)
}
