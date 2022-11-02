package entities

import (
	"strings"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) OffchainTipBotCreateAssignContract(ac *model.OffchainTipBotAssignContract) (userAssignedContract *model.OffchainTipBotAssignContract, err error) {
	return e.repo.OffchainTipBotContract.CreateAssignContract(ac)
}

func (e *Entity) OffchainTipBotDeleteExpiredAssignContract() (err error) {
	return e.repo.OffchainTipBotContract.DeleteExpiredAssignContract()
}

func (e *Entity) GetUserBalances(userID string) (bals []response.GetUserBalances, err error) {
	userBals, err := e.repo.OffchainTipBotUserBalances.GetUserBalances(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalances] - failed to get user balances")
		return []response.GetUserBalances{}, err
	}

	listCoinIDs := []string{}
	for _, userBal := range userBals {
		coinID := strings.Replace(strings.ToLower(userBal.Token.TokenName), " ", "-", -1)
		listCoinIDs = append(listCoinIDs, coinID)
		bals = append(bals, response.GetUserBalances{
			ID:       coinID,
			Name:     userBal.Token.TokenName,
			Symbol:   userBal.Token.TokenSymbol,
			Balances: userBal.Amount,
		})

	}

	tokenPrices, err := e.svc.CoinGecko.GetCoinPrice(listCoinIDs, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"listCoinIDs": listCoinIDs}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
		return []response.GetUserBalances{}, err
	}

	for i, bal := range bals {
		bals[i].BalancesInUSD = tokenPrices[bal.ID] * bal.Balances
	}

	return bals, nil
}

func (e *Entity) SweepNativeToken() ([]response.OffchainTipBotSweepToken, error) {
	contract, err := e.repo.OffchainTipBotContract.GetAll()
	if err != nil {
		e.log.Error(err, "[repo.OffchainTipBotContract.GetAll] - failed to get all contract")
		return nil, err
	}

	txs := make([]response.OffchainTipBotSweepToken, 0)
	for _, c := range contract {
		tx, err := e.svc.Abi.SweepNativeToken(c.Chain.RPCURL, c.ContractAddress)
		if err != nil {
			e.log.Fields(logger.Fields{"contract": c}).Error(err, "[svc.Abi.SweepNativeToken] - failed to sweep native token")
			return nil, err
		}

		_, err = e.repo.OffchainTipBotContract.UpdateContract(&c)
		if err != nil {
			e.log.Fields(logger.Fields{"contract": c}).Error(err, "[repo.OffchainTipBotContract.UpdateContract] - failed to update contract")
			return nil, err
		}

		txs = append(txs, response.OffchainTipBotSweepToken{
			ContractAddress: c.ContractAddress,
			Symbol:          c.Chain.ChainName,
			TxHash:          tx.Hash().Hex(),
		})
	}
	return txs, nil
}
