package user_nft_balance

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(balance model.UserNFTBalance) error
	GetUserNFTBalancesByUserInGuild(nftCollectionIds []string, guildID string, userDiscordId string) (*model.UserNFTBalancesByGuild, error)
}
