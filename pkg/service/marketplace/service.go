package marketplace

import (
	res "github.com/defipod/mochi/pkg/response"
)

type Service interface {
	HandleMarketplaceLink(contractAddress, chain string) string
	GetCollectionFromOpensea(collectionSymbol string) (*res.OpenseaGetCollectionResponse, error)
	GetCollectionFromQuixotic(collectionSymbol string) (*res.QuixoticCollectionResponse, error)
	GetCollectionFromPaintswap(address string) (*response.PaintswapCollectionResponse, error)
	GetOpenseaAssetContract(address string) (*response.OpenseaAssetContractResponse, error)
}
