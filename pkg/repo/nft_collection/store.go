package nftcollection

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByAddress(address string) (*model.NFTCollection, error)
	GetBySymbol(symbol string) (*model.NFTCollection, error)
}
