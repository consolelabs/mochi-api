package nftsoulbound

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateSoulBounds(soulbounds []model.NftSoulbound) error
	GetSoulBoundsByCollectionAddress(collectionAddress string) (soulbounds []model.NftSoulbound, err error)
}
