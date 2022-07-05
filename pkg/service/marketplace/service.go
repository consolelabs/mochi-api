package marketplace

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	HandleMarketplaceLink(contractAddress, chain string) string
	GetCollectionFromOpensea(collectionSymbol string) (*response.OpenseaGetCollectionResponse, error)
}
