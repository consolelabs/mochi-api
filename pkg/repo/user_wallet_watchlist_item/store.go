package userwalletwatchlistitem

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.UserWalletWatchlistItem, error)
	GetOne(GetOneQuery) (*model.UserWalletWatchlistItem, error)
	Upsert(*model.UserWalletWatchlistItem) error
	Remove(DeleteQuery) error
	UpdateOwnerFlag(profileID, address string, isOwner bool) error
	Update(*model.UserWalletWatchlistItem) error
}
