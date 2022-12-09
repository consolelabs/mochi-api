package nftaddrequesthistory

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(GetOneQuery) (*model.NftAddRequestHistory, error)
	UpsertOne(model.NftAddRequestHistory) error
}
