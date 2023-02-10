package userwalletwatchlistitem

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(userID string) ([]model.UserWalletWatchlistItem, error)
	GetOne(GetOneQuery) (*model.UserWalletWatchlistItem, error)
	Create(*model.UserWalletWatchlistItem) error
	Remove(DeleteQuery) error
}
