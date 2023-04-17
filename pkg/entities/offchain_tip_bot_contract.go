package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) OffchainTipBotDeleteExpiredAssignContract() (err error) {
	return e.repo.OffchainTipBotContract.DeleteExpiredAssignContract()
}

func (e *Entity) GetUserBalances(userID string) ([]response.GetUserBalances, error) {
	userBals, err := e.repo.OffchainTipBotUserBalances.GetUserBalances(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalances] - failed to get user balances")
		return nil, err
	}

	listCoinIDs := []string{}
	res := make([]response.GetUserBalances, 0, len(userBals))
	for _, userBal := range userBals {
		coinID := userBal.Token.CoinGeckoID
		listCoinIDs = append(listCoinIDs, coinID)
		res = append(res, response.GetUserBalances{
			ID:       coinID,
			Name:     userBal.Token.TokenName,
			Symbol:   userBal.Token.TokenSymbol,
			Balances: userBal.Amount,
		})

	}

	tokenPrices, err := e.svc.CoinGecko.GetCoinPrice(listCoinIDs, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"listCoinIDs": listCoinIDs}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
		return nil, err
	}

	for i, bal := range res {
		res[i].RateInUSD = tokenPrices[bal.ID]
		res[i].BalancesInUSD = tokenPrices[bal.ID] * bal.Balances
	}

	return res, nil
}
