package nftcollection

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByAddress(address string) (*model.NFTCollection, error)
	GetBySymbol(symbol string) (*model.NFTCollection, error)
	GetByID(id string) (*model.NFTCollection, error)
	Create(collection model.NFTCollection) (*model.NFTCollection, error)
	ListAll() ([]model.NFTCollection, error)
	ListAllNFTCollectionConfigs() ([]model.NFTCollectionConfig, error)
	ListByGuildID(guildID string) ([]model.NFTCollection, error)
}
