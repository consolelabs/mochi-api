package nftsoulbound

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) CreateSoulBounds(soulbounds []model.NftSoulbound) error {
	return pg.db.Create(soulbounds).Error
}

func (pg *pg) GetSoulBoundsByCollectionAddress(collectionAddress string) (soulbounds []model.NftSoulbound, err error) {
	return soulbounds, pg.db.Where("collection_address = ?", collectionAddress).Find(&soulbounds).Error
}
