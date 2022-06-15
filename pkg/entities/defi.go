package entities

import (
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (e *Entity) GetHistoricalMarketChart(c *gin.Context) (*response.CoinPriceHistoryResponse, error, int) {
	req, err := request.ValidateRequest(c)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	data, err, statusCode := e.svc.CoinGecko.GetHistoricalMarketData(req)
	if err != nil {
		return nil, err, statusCode
	}

	return data, nil, http.StatusOK
}

func (e *Entity) generateInDiscordWallet(user *model.User) error {
	if !user.InDiscordWalletAddress.Valid || user.InDiscordWalletAddress.String == "" {
		inDiscordWalletNumber := e.repo.Users.GetLatestWalletNumber() + 1
		inDiscordAddress, err := e.dcwallet.GetAccountByWalletNumber(inDiscordWalletNumber)
		if err != nil {
			err = fmt.Errorf("error getting wallet address: %v", err)
			return err
		}

		user.InDiscordWalletNumber = model.JSONNullInt64{NullInt64: sql.NullInt64{Int64: int64(inDiscordWalletNumber), Valid: true}}
		user.InDiscordWalletAddress = model.JSONNullString{NullString: sql.NullString{String: inDiscordAddress.Address.Hex(), Valid: true}}

		if err := e.repo.Users.Create(user); err != nil {
			err = fmt.Errorf("error upsert user: %v", err)
			return err
		}
	}

	return nil
}

func (e *Entity) InDiscordWalletTransfer(req request.TransferRequest) ([]response.InDiscordWalletTransferResponse, []string) {
	res := []response.InDiscordWalletTransferResponse{}
	errs := []string{}

	fromUser, err := e.repo.Users.GetOne(req.Sender)
	if err != nil {
		errs = append(errs, fmt.Sprintf("user not found: %v", err))
		return nil, errs
	}
	if err = e.generateInDiscordWallet(fromUser); err != nil {
		errs = append(errs, fmt.Sprintf("cannot generate in-discord wallet: %v", err))
		return nil, errs
	}

	toUsers, err := e.repo.Users.GetByDiscordIDs(req.Recipients)
	if err != nil || len(toUsers) == 0 {
		errs = append(errs, fmt.Sprintf("recipients not found: %v", err))
		return nil, errs
	}
	amountEach := req.Amount / float64(len(toUsers))
	if req.Each {
		amountEach = req.Amount
	}

	fromAcc, err := e.dcwallet.GetAccountByWalletNumber(int(fromUser.InDiscordWalletNumber.Int64))
	if err != nil {
		errs = append(errs, fmt.Sprintf("error getting user address: %v", err))
		return nil, errs
	}

	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency), true)
	if err != nil {
		errs = append(errs, fmt.Sprintf("error getting token info: %v", err))
		return nil, errs
	}

	nonce := -1
	for _, toUser := range toUsers {
		if err = e.generateInDiscordWallet(&toUser); err != nil {
			errs = append(errs, fmt.Sprintf("cannot generate in-discord wallet: %v", err))
			continue
		}

		toAcc, err := e.dcwallet.GetAccountByWalletNumber(int(toUser.InDiscordWalletNumber.Int64))
		if err != nil {
			errs = append(errs, fmt.Sprintf("error getting user address: %v", err))
			continue
		}

		signedTx, transferredAmount, err := e.transfer(fromAcc, toAcc, amountEach, token, nonce, req.All)
		if err != nil {
			errs = append(errs, fmt.Sprintf("error transfer: %v", err))
			continue
		}
		nonce = int(signedTx.Nonce()) + 1
		transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()

		gActivityConfig, err := e.GetGuildActivityConfig(req.GuildID, string(req.TransferType))
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}

		if err := e.repo.GuildUserActivityLog.CreateOne(model.GuildUserActivityLog{
			GuildID:      req.GuildID,
			UserID:       req.Sender,
			ActivityName: gActivityConfig.Activity.Name,
			EarnedXP:     gActivityConfig.Activity.XP,
			CreatedAt:    time.Now(),
		}); err != nil {
			errs = append(errs, fmt.Sprintf("error create activity log: %v", err))
			continue
		}

		res = append(res, response.InDiscordWalletTransferResponse{
			FromDiscordID:  req.Sender,
			ToDiscordID:    toUser.ID,
			Amount:         transferredAmount,
			Cryptocurrency: req.Cryptocurrency,
			TxHash:         signedTx.Hash().Hex(),
			TxUrl:          fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
			TransactionFee: transactionFee,
		})
	}
	if len(errs) == 0 {
		errs = nil
	}

	return res, errs
}

func (e *Entity) InDiscordWalletWithdraw(req request.TransferRequest) (res response.InDiscordWalletWithdrawResponse, err error) {
	fromUser, err := e.repo.Users.GetOne(req.Sender)
	if err != nil {
		err = fmt.Errorf("user not found: %v", err)
		return
	}
	if err = e.generateInDiscordWallet(fromUser); err != nil {
		err = fmt.Errorf("cannot generate in-discord wallet: %v", err)
		return
	}

	fromAccount, err := e.dcwallet.GetAccountByWalletNumber(int(fromUser.InDiscordWalletNumber.Int64))
	if err != nil {
		err = fmt.Errorf("error getting user address: %v", err)
		return
	}

	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency), true)
	if err != nil {
		err = fmt.Errorf("error getting token info: %v", err)
		return
	}

	signedTx, transferredAmount, err := e.transfer(fromAccount,
		accounts.Account{Address: common.HexToAddress(req.Recipients[0])},
		req.Amount,
		token, -1, req.All)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return
	}

	gActivityConfig, err := e.GetGuildActivityConfig(req.GuildID, string(req.TransferType))
	if err != nil {
		return
	}

	if err = e.repo.GuildUserActivityLog.CreateOne(model.GuildUserActivityLog{
		GuildID:      req.GuildID,
		UserID:       req.Sender,
		ActivityName: gActivityConfig.Activity.Name,
		EarnedXP:     gActivityConfig.Activity.XP,
		CreatedAt:    time.Now(),
	}); err != nil {
		err = fmt.Errorf("error create activity log: %v", err)
		return
	}

	withdrawalAmount := util.WeiToEther(signedTx.Value())
	transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()

	res = response.InDiscordWalletWithdrawResponse{
		FromDiscordId:    req.Sender,
		ToAddress:        req.Recipients[0],
		Amount:           transferredAmount,
		Cryptocurrency:   req.Cryptocurrency,
		TxHash:           signedTx.Hash().Hex(),
		TxURL:            fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
		WithdrawalAmount: withdrawalAmount,
		TransactionFee:   transactionFee,
	}
	return
}

func (e *Entity) balances(address string, tokens []model.Token) (map[string]float64, error) {
	balances := make(map[string]float64, 0)
	for _, token := range tokens {
		chain := e.dcwallet.Chain(token.ChainID)
		if chain == nil {
			return nil, errors.New("cryptocurrency not supported")
		}

		bals, err := chain.Balances(
			address, []model.Token{token},
		)
		if err != nil {
			err = fmt.Errorf("error getting balances: %v", err)
			return nil, err
		}
		for k, v := range bals {
			balances[k] = v
		}
	}

	return balances, nil
}

func (e *Entity) transfer(fromAccount accounts.Account, toAccount accounts.Account, amount float64, token model.Token, nonce int, all bool) (*types.Transaction, float64, error) {
	chain := e.dcwallet.Chain(token.ChainID)
	if chain == nil {
		return nil, 0, errors.New("cryptocurrency not supported")
	}
	signedTx, amount, err := chain.Transfer(
		fromAccount,
		toAccount,
		amount,
		token,
		nonce,
		all,
	)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return nil, 0, err
	}

	return signedTx, amount, nil
}

func (e *Entity) InDiscordWalletBalances(guildID, discordID string) (*response.UserBalancesResponse, error) {
	response := &response.UserBalancesResponse{}

	user, err := e.repo.Users.GetOne(discordID)
	if err != nil {
		err = fmt.Errorf("failed to get user %s: %v", discordID, err)
		return nil, err
	}

	gTokens, err := e.repo.GuildConfigToken.GetByGuildID(guildID)
	if err != nil {
		err = fmt.Errorf("failed to get guild %s tokens - err: %v", guildID, err)
		return nil, err
	}
	if len(gTokens) == 0 {
		if err = e.InitGuildDefaultTokenConfigs(guildID); err != nil {
			return nil, err
		}
		return e.InDiscordWalletBalances(guildID, discordID)
	}

	var tokens []model.Token
	for _, gToken := range gTokens {
		tokens = append(tokens, *gToken.Token)
	}

	if user.InDiscordWalletAddress.String == "" {
		if err = e.generateInDiscordWallet(user); err != nil {
			err = fmt.Errorf("cannot generate in-discord wallet: %v", err)
			return nil, err
		}
	}

	balances, err := e.balances(user.InDiscordWalletAddress.String, tokens)
	if err != nil {
		err = fmt.Errorf("cannot get user balances: %v", err)
		return nil, err
	}
	response.Balances = balances

	var coinIDs []string
	for _, token := range tokens {
		coinIDs = append(coinIDs, token.CoinGeckoID)
	}

	tokenPrices, err := e.svc.CoinGecko.GetCoinPrice(coinIDs, "usd")
	if err != nil {
		err = fmt.Errorf("cannot get user balances: %v", err)
		return nil, err
	}

	response.BalancesInUSD = make(map[string]float64)
	for _, token := range tokens {
		response.BalancesInUSD[token.Symbol] = response.Balances[token.Symbol] * tokenPrices[token.CoinGeckoID]
	}

	return response, nil
}

func (e *Entity) GetSupportedTokens() (tokens []model.Token, err error) {
	tokens, err = e.repo.Token.GetAllSupported()
	if err != nil {
		err = fmt.Errorf("failed to get supported tokens - err: %v", err)
		return
	}
	return
}

func (e *Entity) GetCoinData(c *gin.Context) (*response.GetCoinResponse, error, int) {
	coinID := c.Param("id")
	if coinID == "" {
		return nil, fmt.Errorf("id is required"), http.StatusBadRequest
	}

	data, err, statusCode := e.svc.CoinGecko.GetCoin(coinID)
	if err != nil {
		return nil, err, statusCode
	}

	return data, nil, http.StatusOK
}

func (e *Entity) SearchCoins(c *gin.Context) ([]response.SearchedCoin, error, int) {
	query := c.Query("query")
	if query == "" {
		return nil, fmt.Errorf("query is required"), http.StatusBadRequest
	}

	data, err, statusCode := e.svc.CoinGecko.SearchCoins(query)
	if err != nil {
		return nil, err, statusCode
	}

	return data, nil, http.StatusOK
}

func (e *Entity) InitGuildDefaultTokenConfigs(guildID string) error {
	tokens, err := e.repo.Token.GetDefaultTokens()
	if err != nil {
		return err
	}
	if len(tokens) == 0 {
		return fmt.Errorf("No default tokens found")
	}

	var configs []model.GuildConfigToken
	for _, token := range tokens {
		configs = append(configs, model.GuildConfigToken{
			TokenID: token.ID,
			GuildID: guildID,
			Active:  true,
		})
	}

	return e.repo.GuildConfigToken.UpsertMany(configs)
}

func (e *Entity) GetHighestTicker(symbol string, interval int) ([]string, error) {
	var coinData []string
	coinRequest := request.GetMarketChartRequest{CoinID: symbol, Currency: "usd", Days: interval}
	data, err, _ := e.svc.CoinGecko.GetHistoricalMarketData(&coinRequest)
	if err != nil {
		return coinData, err
	}
	highestPrice := util.GetMaxFloat64(data.Prices)
	coinData = append(coinData, symbol, fmt.Sprintf("%v", interval), fmt.Sprintf("%v", highestPrice))
	return coinData, nil
}

func (e *Entity) InDiscordWalletBalancesVer2(guildID, discordID string) (*response.UserBalancesResponse, error) {
	response := &response.UserBalancesResponse{}

	gTokens, err := e.repo.GuildConfigToken.GetByGuildID(guildID)
	if err != nil {
		err = fmt.Errorf("failed to get guild %s tokens - err: %v", guildID, err)
		return nil, err
	}
	if len(gTokens) == 0 {
		if err = e.InitGuildDefaultTokenConfigs(guildID); err != nil {
			return nil, err
		}
		return e.InDiscordWalletBalancesVer2(guildID, discordID)
	}

	balances, err := e.repo.UserBalance.GetUserBalances(discordID)
	if err != nil {
		return nil, err
	}

	gTokenMap := make(map[int]*model.Token)
	for _, gToken := range gTokens {
		gTokenMap[gToken.TokenID] = gToken.Token
	}

	response.Balances = make(map[string]float64)
	for _, b := range balances {
		if token, ok := gTokenMap[b.TokenID]; ok {
			response.Balances[token.Symbol] = b.Balance
		}
	}

	var coinIDs []string
	for _, b := range balances {
		coinIDs = append(coinIDs, b.Token.CoinGeckoID)
	}

	tokenPrices, err := e.svc.CoinGecko.GetCoinPrice(coinIDs, "usd")
	if err != nil {
		err = fmt.Errorf("cannot get user balances: %v", err)
		return nil, err
	}

	response.BalancesInUSD = make(map[string]float64)
	for _, b := range balances {
		token := b.Token
		response.BalancesInUSD[token.Symbol] = response.Balances[token.Symbol] * tokenPrices[token.CoinGeckoID]
	}

	return response, nil
}

func (e *Entity) InDiscordWalletTransferVer2(req request.TransferRequest) ([]response.InDiscordWalletTransferResponse, error) {
	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency), true)
	if err != nil {
		return nil, err
	}

	userBal, err := e.repo.UserBalance.GetOne(req.Sender, token.ID)
	if err != nil {
		return nil, err
	}
	if req.All {
		req.Amount = userBal.Balance
		req.Each = false
	}

	amountEach := req.Amount / float64(len(req.Recipients))
	if req.Each {
		amountEach = req.Amount
	}
	totalAmount := amountEach * float64(len(req.Recipients))

	if userBal.Balance < totalAmount || userBal.Balance == 0 {
		return nil, errors.New("insufficient balance")
	}

	if err := e.repo.UserDefiActivityLog.CreateTransferLogs(req, token.ID, amountEach, totalAmount); err != nil {
		return nil, err
	}

	var res []response.InDiscordWalletTransferResponse
	for _, recipient := range req.Recipients {
		res = append(res, response.InDiscordWalletTransferResponse{
			FromDiscordID:  req.Sender,
			ToDiscordID:    recipient,
			Amount:         amountEach,
			Cryptocurrency: req.Cryptocurrency,
		})

		if req.GuildID == "" {
			continue
		}

		gActivityConfig, err := e.GetGuildActivityConfig(req.GuildID, string(req.TransferType))
		if err != nil {
			return nil, err
		}

		if err := e.repo.GuildUserActivityLog.CreateOne(model.GuildUserActivityLog{
			GuildID:      req.GuildID,
			UserID:       req.Sender,
			ActivityName: gActivityConfig.Activity.Name,
			EarnedXP:     gActivityConfig.Activity.XP,
			CreatedAt:    time.Now(),
		}); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (e *Entity) GetGuildActivityConfig(guildID, transferType string) (*model.GuildConfigActivity, error) {
	gActivityConfig, err := e.repo.GuildConfigActivity.GetOneByActivityName(guildID, transferType)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		if err = e.repo.GuildConfigActivity.ForkDefaulActivityConfigs(guildID); err != nil {
			return nil, err
		}
		if gActivityConfig, err = e.repo.GuildConfigActivity.GetOneByActivityName(guildID, transferType); err != nil {
			return nil, err
		}
	}
	return gActivityConfig, nil
}

func (e *Entity) InDiscordWalletWithdrawVer2(req request.TransferRequest) (*response.InDiscordWalletWithdrawResponse, error) {
	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency), true)
	if err != nil {
		return nil, fmt.Errorf("error getting token info: %v", err)
	}

	userBal, err := e.repo.UserBalance.GetOne(req.Sender, token.ID)
	if err != nil {
		return nil, err
	}
	if !req.All && userBal.Balance < req.Amount {
		return nil, errors.New("insufficient balance")
	}
	if req.All {
		req.Amount = userBal.Balance
	}

	fromAccount, err := e.dcwallet.GetAccountByWalletNumber(0)
	if err != nil {
		return nil, fmt.Errorf("error getting Mochi fund account: %v", err)
	}

	signedTx, transferredAmount, err := e.transferVer2(
		fromAccount, accounts.Account{Address: common.HexToAddress(req.Recipients[0])},
		req.Amount, token, req.All,
	)
	if err != nil {
		return nil, fmt.Errorf("error transfer: %v", err)
	}

	if err := e.repo.UserDefiActivityLog.CreateTransferLogs(req, token.ID, 0, req.Amount); err != nil {
		return nil, fmt.Errorf("error creating transfer logs: %v", err)
	}

	gActivityConfig, err := e.GetGuildActivityConfig(req.GuildID, string(req.TransferType))
	if err != nil {
		return nil, fmt.Errorf("error getting guild activity config: %v", err)
	}

	if err = e.repo.GuildUserActivityLog.CreateOne(model.GuildUserActivityLog{
		GuildID:      req.GuildID,
		UserID:       req.Sender,
		ActivityName: gActivityConfig.Activity.Name,
		EarnedXP:     gActivityConfig.Activity.XP,
		CreatedAt:    time.Now(),
	}); err != nil {
		return nil, fmt.Errorf("error create withdraw activity log: %v", err)
	}

	withdrawalAmount := util.WeiToEther(signedTx.Value())
	transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()

	return &response.InDiscordWalletWithdrawResponse{
		FromDiscordId:    req.Sender,
		ToAddress:        req.Recipients[0],
		Amount:           transferredAmount,
		Cryptocurrency:   req.Cryptocurrency,
		TxHash:           signedTx.Hash().Hex(),
		TxURL:            fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
		WithdrawalAmount: withdrawalAmount,
		TransactionFee:   transactionFee,
	}, nil
}

func (e *Entity) transferVer2(fromAccount accounts.Account, toAccount accounts.Account, amount float64, token model.Token, all bool) (*types.Transaction, float64, error) {
	chain := e.dcwallet.Chain(token.ChainID)
	if chain == nil {
		return nil, 0, errors.New("cryptocurrency not supported")
	}
	signedTx, amount, err := chain.TransferVer2(
		fromAccount,
		toAccount,
		amount,
		token,
		all,
	)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return nil, 0, err
	}

	return signedTx, amount, nil
}
