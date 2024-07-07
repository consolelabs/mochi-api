package job

import (
	"strings"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	sliceutils "github.com/defipod/mochi/pkg/util/slice"
)

type updateUserNFTBalances struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewUpdateUserNFTBalancesJob(e *entities.Entity) Job {
	return &updateUserNFTBalances{
		entity: e,
		log:    e.GetLogger(),
	}
}

// func (c *updateUserNFTBalances) Run() error {
// 	c.log.Info("start updating user nft balances")
// 	nftConfigs, err := c.entity.ListAllNFTCollectionConfigs()
// 	if err != nil {
// 		c.log.Error(err, "entity.ListAllNFTCollectionConfigs failed")
// 		return err
// 	}

// 	if len(nftConfigs) == 0 {
// 		c.log.Info("entity.ListAllNFTCollectionConfigs - no data found")
// 		return nil
// 	}

// 	was, err := c.entity.ListAllWalletAddresses()
// 	if err != nil {
// 		c.log.Error(err, "entity.ListAllWalletAddresses failed")
// 		return err
// 	}

// 	// // TODO: get from mochi profile
// 	for _, nftConfig := range nftConfigs {
// 		c.log.Infof("start nft config %s", nftConfig.Address)
// 		balanceOf, err := c.entity.GetNekoBalanceFunc(nftConfig)
// 		if err != nil {
// 			c.log.Errorf(err, "failed to get nft balance function of collection %s", nftConfig.Name)
// 			continue
// 		}

// 		for _, wa := range was {
// 			holding, staking, err := balanceOf(wa.Address)
// 			if err != nil {
// 				c.log.Errorf(err, "failed to get nft %s balance of address %s", nftConfig.Name, wa.Address)
// 				continue
// 			}

// 			balance := model.UserNFTBalance{
// 				UserAddress:     wa.Address,
// 				ChainType:       wa.ChainType,
// 				NFTCollectionID: nftConfig.ID,
// 				Balance:         int(holding.Int64()),
// 				ProfileID:       wa.ProfileId,
// 				StakingNekos:    int(staking.Int64()),
// 			}

// 			err = c.entity.NewUserNFTBalance(balance)
// 			if err != nil {
// 				c.log.Errorf(err, "failed to create user nft balance of address %s", wa.Address)
// 				continue
// 			}
// 		}
// 	}

// 	return nil
// }

func (c *updateUserNFTBalances) Run() error {
	c.log.Info("start updating user nft balances")
	collections, err := c.entity.ListAllNFTCollections()
	if err != nil {
		c.log.Error(err, "entity.ListAllNFTCollections failed")
		return err
	}

	nekoAddr := "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73"
	nekoCol := sliceutils.Find(collections, func(c model.NFTCollection) bool {
		return strings.EqualFold(nekoAddr, c.Address)
	})

	if nekoCol == nil {
		c.log.Info("Neko Collection not found")
		return nil
	}

	data, err := c.entity.GetNekoHolders(*nekoCol)
	if err != nil {
		c.log.Errorf(err, "entity.GetNekoHolders() failed")
		return err
	}

	for _, bal := range data {
		err = c.entity.NewUserNFTBalance(bal)
		if err != nil {
			c.log.Errorf(err, "NewUserNFTBalance() failed")
			continue
		}
	}

	return nil
}
