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

func (e *Entity) MigrateBalance() error {
	tokens, err := e.repo.Token.GetAll()
	if err != nil {
		e.log.Error(err, "[entities.migrateBalance] - failed to get supported tokens")
	}
	// get user check migrate bals
	users, _ := e.repo.Users.Get50Users()
	for _, user := range users {
		if user.IsMigrateBal {
			continue
		}
		// get old balances
		if user.InDiscordWalletAddress.String == "" {
			e.log.Info("[entities.migrateBalance] - user has no valid wallet")
			continue
		}

		balances, _ := e.balances(user.InDiscordWalletAddress.String, tokens)
		// migrate to tip_bot_offchain_balances
		for symbol, balance := range balances {
			if balance == 0 {
				continue
			}
			token, err := e.repo.OffchainTipBotTokens.GetBySymbol(symbol)
			if err != nil {
				e.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[entities.migrateBalance] - failed to get offchain tip bot token")
				continue
			}
			err = e.repo.OffchainTipBotUserBalances.CreateIfNotExists(&model.OffchainTipBotUserBalance{
				UserID:  user.ID,
				TokenID: token.ID,
				Amount:  balance,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"userID": user.ID, "tokenID": token.ID, "balance": balance}).Error(err, "[entities.migrateBalance] - failed to add offchain tip bot balance")
			}
		}
		e.repo.Users.UpdateUserIsMigrateBals(user.ID)
	}

	return nil
}
