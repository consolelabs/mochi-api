package nftcollection

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"strings"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetByAddress(address string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	err := pg.db.Table("nft_collections").Where("lower(address) = ?", strings.ToLower(address)).First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (pg *pg) GetBySymbol(symbol string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	err := pg.db.Table("nft_collections").Where("lower(symbol) = ?", strings.ToLower(symbol)).First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}
