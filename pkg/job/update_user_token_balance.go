package job

import (
	"math"
	"math/big"
	"strconv"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/service"
)

type updateUserTokenBalances struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewUpdateUserTokenBalancesJob(e *entities.Entity, l logger.Logger, svc service.Service) Job {
	return &updateUserTokenBalances{
		entity: e,
		log:    l,
	}
}

func (c *updateUserTokenBalances) Run() error {
	c.log.Info("start updating user token balances")
	tokens, err := c.entity.ListAllConfigTokens()
	if err != nil {
		c.log.Error(err, "entity.ListAllConfigTokens failed")
		return err
	}

	if len(tokens) == 0 {
		c.log.Info("entity.ListAllConfigTokens - no data found")
		return nil
	}

	was, err := c.entity.ListAllWalletAddresses()
	if err != nil {
		c.log.Error(err, "entity.ListAllWalletAddresses failed")
		return err
	}

	for _, token := range tokens {
		c.log.Infof("start token config %s", token.Address)
		chainIdStr := strconv.Itoa(token.ChainID)
		balanceOf, err := c.entity.GetTokenBalanceFunc(chainIdStr, token)
		if err != nil {
			c.log.Errorf(err, "failed to get token balance function of token %s", token.Name)
			continue
		}

		for _, wa := range was {
			n, err := balanceOf(wa.Address)
			if err != nil {
				c.log.Errorf(err, "failed to get token %s balance of address %s", token.Name, wa.Address)
				continue
			}

			balance := new(big.Float)
			balance.SetString(n.String())
			value := new(big.Float).Quo(balance, big.NewFloat(math.Pow10(token.Decimals)))
			v, _ := value.Float64()

			userTokenBalance := model.UserTokenBalance{
				UserAddress: wa.Address,
				ChainType:   wa.ChainType,
				TokenID:     token.ID,
				Balance:     v,
			}

			err = c.entity.NewUserTokenBalance(userTokenBalance)
			if err != nil {
				c.log.Errorf(err, "failed to create user token balance of address %s", wa.Address)
				continue
			}
		}
	}

	return nil
}
