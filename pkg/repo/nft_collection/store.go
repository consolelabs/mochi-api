package nftcollection

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	GetByAddress(address string) (*model.NFTCollection, error)
	GetBySymbol(symbol string) (*model.NFTCollection, error)
	GetByID(id string) (*model.NFTCollection, error)
	GetNewListed(interval int, page int, size int) ([]model.NewListedNFTCollection, int64, error)
	Create(collection model.NFTCollection) (*model.NFTCollection, error)
	ListAll() ([]model.NFTCollection, error)
	ListAllWithPaging(page int, size int) ([]model.NFTCollection, int64, error)
	ListAllNFTCollectionConfigs() ([]model.NFTCollectionConfig, error)
	ListByGuildID(guildID string) ([]model.NFTCollection, error)
}
