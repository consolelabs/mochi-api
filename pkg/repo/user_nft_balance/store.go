package user_nft_balance

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(balance model.UserNFTBalance) error
	GetUserNFTBalancesByUserInGuild(guildID string) ([]model.UserAddressNFTBalancesByGuild, error)
	List(ListQuery) ([]model.UserNFTBalance, error)
	TotalBalance(collectionID string) (int, error)
	IsExists(collectionID, userAddress string) (bool, error)
}
