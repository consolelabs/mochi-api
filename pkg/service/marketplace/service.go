package marketplace

import (
	res "github.com/defipod/mochi/pkg/response"
)

type Service interface {
	HandleMarketplaceLink(contractAddress, chain string) string
	GetCollectionFromOpensea(collectionSymbol string) (*res.OpenseaGetCollectionResponse, error)
	GetCollectionFromQuixotic(collectionSymbol string) (*res.QuixoticCollectionResponse, error)
	GetCollectionFromPaintswap(address string) (*res.PaintswapCollectionResponse, error)
	GetOpenseaAssetContract(address string) (*res.OpenseaAssetContractResponse, error)
}
