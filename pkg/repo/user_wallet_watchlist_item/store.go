package userwalletwatchlistitem

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.UserWalletWatchlistItem, error)
	GetOne(GetOneQuery) (*model.UserWalletWatchlistItem, error)
	Create(*model.UserWalletWatchlistItem) error
	Remove(DeleteQuery) error
	UpdateOwnerFlag(userID, address string, isOwner bool) error
}
