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
	nftConfigs, err := c.entity.ListAllNFTCollectionConfigs()
	if err != nil {
		return err
	}

	if len(nftConfigs) == 0 {
		c.log.Infof("no nft collection configs found")
		return nil
	}

	was, err := c.entity.ListAllWalletAddresses()
	if err != nil {
		return err
	}

	for _, nftConfig := range nftConfigs {
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
				TokenID:         nftConfig.TokenID,
				Balance:         n,
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
