package bluemove

import "github.com/defipod/mochi/pkg/model"

type Service interface {
	GetCollection(collectionAddress, chainId string) (*model.NFTCollection, error)
}
