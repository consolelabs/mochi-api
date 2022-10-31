package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/k0kubun/pp"
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

		balances, err := e.balances(user.InDiscordWalletAddress.String, tokens)
		if err != nil {
			pp.Println(err)
		}
		pp.Println(balances)
		// for symbol, balance := range balances {
		// 	if balance == 0 {
		// 		continue
		// 	}

		// 	// transfer money to centralized wallet
		// 	tokenTransfer := model.Token{}
		// 	for _, t := range tokens {
		// 		if t.Symbol == symbol {
		// 			tokenTransfer = t
		// 			break
		// 		}
		// 	}

		// 	if tokenTransfer.IsNative {
		// 		fromAccount, _ := e.dcwallet.GetAccountByWalletNumber(int(user.InDiscordWalletNumber.Int64))

		// 		fmt.Println("Transfer for symbol", symbol)
		// 		signedTx, transferredAmount, err := e.transfer(fromAccount,
		// 			accounts.Account{Address: common.HexToAddress("0x4ec16127E879464bEF6ab310084FAcEC1E4Fe465")},
		// 			balance,
		// 			tokenTransfer, -1, true)

		// 		fmt.Println("signedTx: ", signedTx)
		// 		fmt.Println("transferredAmount: ", transferredAmount)
		// 		fmt.Println(err)
		// 		if signedTx != nil {
		// 			fmt.Println(fmt.Sprintf("%s/%s", tokenTransfer.Chain.TxBaseURL, signedTx.Hash().Hex()))
		// 		}
		// 		if err != nil && !strings.Contains(err.Error(), "error transfer: not found") {
		// 			if strings.Contains(err.Error(), "insufficient funds for gas") {
		// 				continue
		// 			}
		// 			if strings.Contains(err.Error(), "This token dose not have native token") {
		// 				continue
		// 			}
		// 			return err
		// 		}

		// 		// process after transfer
		// 		offchainToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(symbol)
		// 		if err != nil {
		// 			e.log.Error(err, "cannot get offchain token: ")
		// 			return err
		// 		}

		// 		// create user balance
		// 		err = e.repo.OffchainTipBotUserBalances.CreateIfNotExists(&model.OffchainTipBotUserBalance{
		// 			UserID:  user.ID,
		// 			TokenID: offchainToken.ID,
		// 			Amount:  balance,
		// 		})
		// 		if err != nil {
		// 			e.log.Error(err, "cannot create offchain user balance: ")
		// 			return err
		// 		}

		// 		// create migrate balance
		// 		signedTxHash := ""
		// 		if signedTx != nil {
		// 			signedTxHash = signedTx.Hash().Hex()
		// 		}
		// 		e.repo.MigrateBalances.StoreMigrateBalances(&model.MigrateBalance{
		// 			Symbol:            symbol,
		// 			CreatedAt:         time.Now(),
		// 			Username:          user.Username,
		// 			UserDiscordID:     user.ID,
		// 			Txhash:            signedTxHash,
		// 			Txurl:             fmt.Sprintf("%s/%s", tokenTransfer.Chain.TxBaseURL, signedTxHash),
		// 			Transferredamount: transferredAmount,
		// 		})
		// 		if err != nil {
		// 			e.log.Error(err, "cannot store migrate balance: ")
		// 			return err
		// 		}
		// 	} else {
		// 		pp.Println("This is non-native token, so need to wait: ", symbol)
		// 		pp.Println("balance: ", balance)
		// 	}

		// }
	}

	return nil
}
