package indexer

import "github.com/defipod/mochi/pkg/response"

type Indexer interface {
	CreateERC721Contract(CreateERC721ContractRequest) error
	GetNFTCollection(address string) (*response.NFTCollectionResponse, error)
}
