package indexer

import (
	"github.com/defipod/mochi/pkg/response"
	res "github.com/defipod/mochi/pkg/response"
)

type Service interface {
	CreateERC721Contract(CreateERC721ContractRequest) error
	GetNFTCollectionTickers(address, rawQuery string) (*response.IndexerNFTCollectionTickersResponse, error)
	GetNFTTokenTickers(address, tokenID, rawQuery string) (*res.IndexerNFTTokenTickersData, error)
	GetNFTTradingVolume() ([]response.NFTTradingVolume, error)
	GetNFTCollections(query string) (*response.IndexerGetNFTCollectionsResponse, error)
	GetNFTTokens(address, query string) (*response.IndexerGetNFTTokensResponse, error)
	GetNFTDetail(collectionAddress, tokenID string) (*response.IndexerGetNFTTokenDetailResponse, error)
	GetNFTActivity(collectionAddress, tokenID, query string) (*response.IndexerGetNFTActivityResponse, error)
	GetNFTTokenTxHistory(collectionAddress, tokenID string) (*response.IndexerGetNFTTokenTxHistoryResponse, error)
	GetNftSales(addr string, platform string) (*response.NftSalesResponse, error)
	GetNFTContract(addr string) (*response.IndexerContract, error)
	GetNftMetadataAttrIcon() (*response.NftMetadataAttrIconResponse, error)
	GetNFTCollectionTickersForWl(address string) (*res.IndexerNFTCollectionTickersResponse, error)
	GetNftCollectionMetadata(collectionAddress, chainId string) (*res.IndexerNftCollectionMetadataResponse, error)
	GetSoulBoundNFT(collectionAddress string) (*res.IndexerSoulBoundNFTResponse, error)
}
