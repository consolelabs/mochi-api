package bluemove

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetCollections(chainId, page, pageSize string) (*response.BluemoveCollectionsResponse, error)
	SelectBluemoveCollection(collectionAddress, chainId string) (*model.NFTCollection, error)
	GetAllCollections(chainId string) ([]*model.NFTCollection, error)
	ChooseBluemoveChain(chainId string) string
}
