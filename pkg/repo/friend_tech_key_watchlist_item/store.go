package friend_tech_key_watchlist_item

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(item model.FriendTechKeyWatchlistItem) (*model.FriendTechKeyWatchlistItem, error)
	Update(item model.FriendTechKeyWatchlistItem) error
	Delete(id int) error
	DeleteByAddressAndProfileId(address string, profileId string) error
	ListByProfileId(profileId string) ([]model.FriendTechKeyWatchlistItem, error)
	Exist(id int, address string, profileId string) (bool, error)
	Get(id int) (*model.FriendTechKeyWatchlistItem, error)
}
