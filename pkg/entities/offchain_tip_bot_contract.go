package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

func (e *Entity) OffchainTipBotCreateAssignContract(ac *model.OffchainTipBotAssignContract) (userAssignedContract *model.OffchainTipBotAssignContract, err error) {
	return e.repo.OffchainTipBotContract.CreateAssignContract(ac)
}

func (e *Entity) OffchainTipBotDeleteExpiredAssignContract() (err error) {
	return e.repo.OffchainTipBotContract.DeleteExpiredAssignContract()
}

func (e *Entity) GetUserBalances(userID string) (bals []response.GetUserBalances, err error) {
	// userBals, err := e.repo.OffchainTipBotUserBalances.GetUserBalances(userID)
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalances] - failed to get user balances")
	// 	return []response.GetUserBalances{}, err
	// }

	// listCoinIDs := []string{}
	// for _, userBal := range userBals {
	// 	coinID := strings.Replace(strings.ToLower(userBal.Token.TokenName), " ", "-", -1)
	// 	listCoinIDs = append(listCoinIDs, coinID)
	// 	bals = append(bals, response.GetUserBalances{
	// 		ID:       coinID,
	// 		Name:     userBal.Token.TokenName,
	// 		Symbol:   userBal.Token.TokenSymbol,
	// 		Balances: userBal.Amount,
	// 	})

	// }

	// tokenPrices, err := e.svc.CoinGecko.GetCoinPrice(listCoinIDs, "usd")
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"listCoinIDs": listCoinIDs}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
	// 	return []response.GetUserBalances{}, err
	// }

	// for i, bal := range bals {
	// 	bals[i].BalancesInUSD = tokenPrices[bal.ID] * bal.Balances
	// }

	e.MigrateBalance()
	return nil, nil
}

func (e *Entity) MigrateBalance() error {
	tokens, err := e.repo.Token.GetAll()
	if err != nil {
		e.log.Error(err, "[entities.migrateBalance] - failed to get supported tokens")
	}
	// get user check migrate bals
	users, _ := e.repo.Users.Get50Users()
	fmt.Println(users)
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
		for symbol, balance := range balances {
			if balance == 0 {
				continue
			}

			// transfer money to centralized wallet
			tokenTransfer := model.Token{}
			for _, t := range tokens {
				if t.Symbol == symbol {
					tokenTransfer = t
					break
				}
			}

			fromAccount, _ := e.dcwallet.GetAccountByWalletNumber(int(user.InDiscordWalletNumber.Int64))

			fmt.Println("Transfer for symbol", symbol)
			signedTx, transferredAmount, _ := e.transfer(fromAccount,
				accounts.Account{Address: common.HexToAddress("0x4ec16127E879464bEF6ab310084FAcEC1E4Fe465")},
				balance,
				tokenTransfer, -1, true)

			fmt.Println("signedTx: ", signedTx)
			fmt.Println("transferredAmount: ", transferredAmount)
			fmt.Println(err)
			if signedTx != nil {
				fmt.Println(fmt.Sprintf("%s/%s", tokenTransfer.Chain.TxBaseURL, signedTx.Hash().Hex()))
			}

		}
	}

	return nil
}
