package entities

import (
	"strings"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) OffchainTipBotCreateAssignContract(ac *model.OffchainTipBotAssignContract) (userAssignedContract *model.OffchainTipBotAssignContract, err error) {
	return e.repo.OffchainTipBotContract.CreateAssignContract(ac)
}

func (e *Entity) OffchainTipBotDeleteExpiredAssignContract() (err error) {
	return e.repo.OffchainTipBotContract.DeleteExpiredAssignContract()
}

func (e *Entity) GetUserBalances(userID string) (bals []response.GetUserBalances, err error) {
	err = e.migrateBalance(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[e.migrateBalance] - failed to migrate user balances")
		return []response.GetUserBalances{}, err
	}
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

func (e *Entity) migrateBalance(userID string) error {
	// get user check migrate bals
	user, err := e.repo.Users.GetOne(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[entities.migrateBalance] - failed to get user")
		return err
	}
	if user.IsMigrateBal {
		return nil
	}
	// get old balances
	if user.InDiscordWalletAddress.String == "" {
		e.log.Fields(logger.Fields{"userID": userID}).Info("[entities.migrateBalance] - user has no valid wallet")
		return nil
	}
	tokens, err := e.repo.Token.GetAllSupported()
	if err != nil {
		e.log.Error(err, "[entities.migrateBalance] - failed to get supported tokens")
		return err
	}
	balances, err := e.balances(user.InDiscordWalletAddress.String, tokens)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[entities.migrateBalance] - failed to get user balances")
		return err
	}
	// migrate to tip_bot_offchain_balances
	for symbol, balance := range balances {
		token, err := e.repo.OffchainTipBotTokens.GetBySymbol(symbol)
		if err != nil {
			e.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[entities.migrateBalance] - failed to get offchain tip bot token")
			return err
		}
		// check if user has amount token before migrate
		newBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(userID, token.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{"userID": userID, "tokenID": token.ID}).Error(err, "[entities.migrateBalance] - failed to get user offchain tip bot balances")
			return err
		}
		// case no bal found
		if err == gorm.ErrRecordNotFound {
			err = e.repo.OffchainTipBotUserBalances.CreateIfNotExists(&model.OffchainTipBotUserBalance{
				UserID:  userID,
				TokenID: token.ID,
				Amount:  balance,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"userID": userID, "tokenID": token.ID, "balance": balance}).Error(err, "[entities.migrateBalance] - failed to add offchain tip bot balance")
				return err
			}
			continue
		}
		// case there is a transaction in that token before migrate
		newBal.Amount = newBal.Amount + balance
		err = e.repo.OffchainTipBotUserBalances.UpdateUserBalance(newBal)
		if err != nil {
			e.log.Fields(logger.Fields{"userID": userID, "tokenID": token.ID, "balance": balance}).Error(err, "[entities.migrateBalance] - failed to update offchain tip bot balance")
			return err
		}
	}
	err = e.repo.Users.UpdateUserIsMigrateBals(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[entities.migrateBalance] - failed to get update user is_migrate_bal")
		return err
	}
	return nil
}
