package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
)

type updateUserNFTBalances struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewUpdateUserNFTBalancesJob(e *entities.Entity, l logger.Logger) Job {
	return &updateUserNFTBalances{
		entity: e,
		log:    l,
	}
}

func (c *updateUserNFTBalances) Run() error {
	c.log.Info("start updating user nft balances")
	nftConfigs, err := c.entity.ListAllNFTCollectionConfigs()
	if err != nil {
		c.log.Error(err, "entity.ListAllNFTCollectionConfigs failed")
		return err
	}

	if len(nftConfigs) == 0 {
		c.log.Info("entity.ListAllNFTCollectionConfigs - no data found")
		return nil
	}

	was, err := c.entity.ListAllWalletAddresses()
	if err != nil {
		c.log.Error(err, "entity.ListAllWalletAddresses failed")
		return err
	}

	for _, nftConfig := range nftConfigs {
		c.log.Infof("start nft config %s", nftConfig.Address)
		balanceOf, err := c.entity.GetNFTBalanceFunc(nftConfig)
		if err != nil {
			c.log.Errorf(err, "failed to get nft balance function of collection %s", nftConfig.Name)
			continue
		}

		for _, wa := range was {
			n, err := balanceOf(wa.Address)
			if err != nil {
				c.log.Errorf(err, "failed to get nft %s balance of address %s", nftConfig.Name, wa.Address)
				continue
			}

			balance := model.UserNFTBalance{
				UserAddress:     wa.Address,
				ChainType:       wa.ChainType,
				NFTCollectionID: nftConfig.ID,
				Balance:         int(n.Int64()),
			}

			err = c.entity.NewUserNFTBalance(balance)
			if err != nil {
				c.log.Errorf(err, "failed to create user nft balance of address %s", wa.Address)
				continue
			}
		}
	}

	return nil
}
