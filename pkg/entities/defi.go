package entities

import (
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
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

		if err := e.repo.Users.Upsert(user); err != nil {
			e.log.Fields(logger.Fields{"user": user}).Error(err, "[entity.generateInDiscordWallet] repo.Users.Create() failed")
			return err
		}
	}

	return nil
}

func (e *Entity) InDiscordWalletTransfer(req request.TransferRequest) ([]response.InDiscordWalletTransferResponse, []string) {
	res := []response.InDiscordWalletTransferResponse{}
	errs := []string{}

	fromUser, err := e.GetOneOrUpsertUser(req.Sender)
	if err != nil {
		e.log.Fields(logger.Fields{"sender": req.Sender}).Error(err, "[entity.InDiscordWalletTransfer] GetOneOrUpsertUser() failed")
		errs = append(errs, err.Error())
		return nil, errs
	}

	toUsers, err := e.repo.Users.GetByDiscordIDs(req.Recipients)
	if err != nil {
		e.log.Fields(logger.Fields{"recipients": req.Recipients}).Error(err, "[entity.InDiscordWalletTransfer] repo.Users.GetByDiscordIDs() failed")
		errs = append(errs, err.Error())
		return nil, errs
	}
	// create + generate wallet if not exists
	if len(toUsers) == 0 {
		for _, r := range req.Recipients {
			u, err := e.GetOneOrUpsertUser(r)
			if err != nil {
				e.log.Fields(logger.Fields{"discord_id": r}).Error(err, "[entity.InDiscordWalletTransfer] GetOneOrUpsertUser() failed")
				errs = append(errs, err.Error())
				return nil, errs
			}
			toUsers = append(toUsers, *u)
		}
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

	var token model.Token
	if req.Cryptocurrency == "" {
		token, err = e.repo.Token.GetDefaultTokenByGuildID(req.GuildID)
		if err != nil {
			errs = append(errs, fmt.Sprintf("error getting default token: %v", err))
			return nil, errs
		}
	} else {
		token, err = e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency), true)
		if err != nil {
			errs = append(errs, fmt.Sprintf("error getting token info: %v", err))
			return nil, errs
		}
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

		res = append(res, response.InDiscordWalletTransferResponse{
			FromDiscordID:  req.Sender,
			ToDiscordID:    toUser.ID,
			Amount:         transferredAmount,
			Cryptocurrency: token.Symbol,
			TxHash:         signedTx.Hash().Hex(),
			TxUrl:          fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
			TransactionFee: transactionFee,
		})
	}
	if len(errs) == 0 {
		errs = nil
	}
	if len(res) > 0 {
		if _, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
			GuildID:   req.GuildID,
			ChannelID: req.ChannelID,
			UserID:    req.Sender,
			Timestamp: time.Now(),
			Action:    req.TransferType,
		}); err != nil {
			err = fmt.Errorf("error create activity log: %v", err)
			errs = append(errs, err.Error())
		}
	}

	if err := e.sendTransferLogs(req, res); err != nil {
		e.log.Errorf(err, "[entity.InDiscordWalletTransfer] failed")
	}
	return res, errs
}

func (e *Entity) sendTransferLogs(req request.TransferRequest, res []response.InDiscordWalletTransferResponse) error {
	if req.GuildID == "" || len(res) == 0 {
		return nil
	}
	guild, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		return err
	}

	token := strings.ToUpper(res[0].Cryptocurrency)
	var description string
	if req.TransferType == "withdraw" {
		description = fmt.Sprintf("<@%s> has made a withdrawal of **%g %s** to address `%s`", res[0].FromDiscordID, res[0].Amount, token, req.Recipients[0])
	} else {
		var recipients []string
		for _, tx := range res {
			recipients = append(recipients, fmt.Sprintf("<@%s>", tx.ToDiscordID))
		}
		recipientsStr := strings.Join(recipients, ", ")
		description = fmt.Sprintf("<@%s> has sent %s **%g %s** each at <#%s>", res[0].FromDiscordID, recipientsStr, res[0].Amount, token, req.ChannelID)
	}
	return e.svc.Discord.SendGuildActivityLogs(guild.LogChannel, req.Sender, strings.ToUpper(req.TransferType), description)
}

func (e *Entity) InDiscordWalletWithdraw(req request.TransferRequest) (*response.InDiscordWalletWithdrawResponse, error) {
	fromUser, err := e.GetOneOrUpsertUser(req.Sender)
	if err != nil {
		e.log.Fields(logger.Fields{"sender": req.Sender}).Error(err, "[entity.InDiscordWalletWithdraw] GetOneOrUpsertUser() failed")
		return nil, err
	}

	fromAccount, err := e.dcwallet.GetAccountByWalletNumber(int(fromUser.InDiscordWalletNumber.Int64))
	if err != nil {
		err = fmt.Errorf("error getting user address: %v", err)
		return nil, err
	}

	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency), true)
	if err != nil {
		err = fmt.Errorf("error getting token info: %v", err)
		return nil, err
	}

	signedTx, transferredAmount, err := e.transfer(fromAccount,
		accounts.Account{Address: common.HexToAddress(req.Recipients[0])},
		req.Amount,
		token, -1, req.All)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return nil, err
	}

	if req.GuildID == "" {
		// log activity
		if err := e.repo.GuildUserActivityLog.CreateOneNoGuild(model.GuildUserActivityLog{
			UserID:       req.Sender,
			ActivityName: "withdraw",
		}); err != nil {
			err = fmt.Errorf("error create activity log: %v", err)
			return nil, err
		}
	} else {
		if _, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
			GuildID:   req.GuildID,
			ChannelID: req.ChannelID,
			UserID:    req.Sender,
			Timestamp: time.Now(),
			Action:    req.TransferType,
		}); err != nil {
			err = fmt.Errorf("error create activity log: %v", err)
			return nil, err
		}
	}
	withdrawalAmount := util.WeiToEther(signedTx.Value())
	transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()

	res := &response.InDiscordWalletWithdrawResponse{
		FromDiscordID:    req.Sender,
		ToAddress:        req.Recipients[0],
		Amount:           transferredAmount,
		Cryptocurrency:   req.Cryptocurrency,
		TxHash:           signedTx.Hash().Hex(),
		TxURL:            fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
		WithdrawalAmount: withdrawalAmount,
		TransactionFee:   transactionFee,
	}

	err = e.sendTransferLogs(req, []response.InDiscordWalletTransferResponse{
		{
			FromDiscordID:  req.Sender,
			Amount:         transferredAmount,
			Cryptocurrency: token.Symbol,
			TxHash:         signedTx.Hash().Hex(),
			TxUrl:          fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
			TransactionFee: transactionFee,
		},
	})
	if err != nil {
		e.log.Errorf(err, "[entity.InDiscordWalletWithdraw] sendTransferLogs failed")
	}
	return res, nil
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
	user, err := e.GetOneOrUpsertUser(discordID)
	if err != nil {
		e.log.Fields(logger.Fields{"discord_id": discordID}).Error(err, "[entity.InDiscordWalletBalances] GetOneOrUpsertUser() failed")
		return nil, err
	}

	tokens, err := e.GetGuildTokens(guildID)
	if err != nil {
		err = fmt.Errorf("failed to get global default tokens - err: %v", err)
		return nil, err
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

	coinIDs := make([]string, len(tokens))
	for i, token := range tokens {
		coinIDs[i] = token.CoinGeckoID
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

func (e *Entity) SearchCoins(query string) ([]model.CoingeckoSupportedTokens, error) {
	token, err := e.repo.CoingeckoSupportedTokens.GetOne(query)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"query": query}).Error(err, "[entity.SearchCoins] repo.CoingeckoSupportedTokens.GetOne() failed")
		return nil, err
	}
	if err == nil {
		return []model.CoingeckoSupportedTokens{*token}, nil
	}

	searchQ := coingeckosupportedtokens.ListQuery{Symbol: query}
	tokens, err := e.repo.CoingeckoSupportedTokens.List(searchQ)
	if err != nil {
		e.log.Fields(logger.Fields{"searchQ": searchQ}).Error(err, "[entity.SearchCoins] repo.CoingeckoSupportedTokens.List() failed")
		return nil, err
	}

	return tokens, nil
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

func (e *Entity) queryCoins(guildID, query string) ([]model.CoingeckoSupportedTokens, *response.GetCoinResponse, error) {
	config, err := e.repo.GuildConfigDefaultTicker.GetOneByGuildIDAndQuery(guildID, query)
	// if default ticker was set then return ...
	if err == nil {
		coin, err, code := e.svc.CoinGecko.GetCoin(config.DefaultTicker)
		if err != nil {
			e.log.Fields(logger.Fields{"default_ticker": config.DefaultTicker, "code": code}).Error(err, "[entity.queryCoins] svc.CoinGecko.GetCoin failed")
			return nil, nil, err
		}
		return []model.CoingeckoSupportedTokens{{ID: coin.ID, Name: coin.Name, Symbol: coin.Symbol}}, coin, nil
	}

	// ... else SearchCoins()
	searchResult, err := e.SearchCoins(query)
	// searchResult, err, code := e.svc.CoinGecko.SearchCoins(query)
	if err != nil {
		e.log.Fields(logger.Fields{"query": query}).Error(err, "[entity.queryCoins] svc.CoinGecko.SearchCoins failed")
		return nil, nil, err
	}
	switch len(searchResult) {
	case 0:
		e.log.Fields(logger.Fields{"query": query}).Error(err, "[entity.queryCoins] svc.CoinGecko.SearchCoins - no data found")
		return nil, nil, fmt.Errorf("coin %s not found", query)
	case 1:
		coin, err, code := e.svc.CoinGecko.GetCoin(searchResult[0].ID)
		if err != nil {
			e.log.Fields(logger.Fields{"coind_id": searchResult[0].ID, "code": code}).Error(err, "[entity.queryCoins] svc.CoinGecko.GetCoin failed")
			return nil, nil, err
		}
		return searchResult, coin, nil
	default:
		// if multiple search results then respond as suggestions
		return searchResult, nil, nil
	}
}

func (e *Entity) CompareToken(base, target, interval, guildID string) (*response.CompareTokenReponseData, error) {
	baseSearch, baseCoin, err := e.queryCoins(guildID, base)
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID, "base": base}).Error(err, "[entity.CompareToken] queryCoins failed")
		return nil, err
	}
	targetSearch, targetCoin, err := e.queryCoins(guildID, target)
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID, "target": target}).Error(err, "[entity.CompareToken] queryCoins failed")
		return nil, err
	}

	// if multiple coins (either base or target) found then return suggestions
	if len(baseSearch) > 1 || len(targetSearch) > 1 {
		return &response.CompareTokenReponseData{BaseCoinSuggestions: baseSearch, TargetCoinSuggestions: targetSearch}, nil
	}
	baseID := baseSearch[0].ID
	targetID := targetSearch[0].ID

	// get coins ohlc
	baseOhlc, err, code := e.svc.CoinGecko.GetHistoryCoinInfo(baseID, interval)
	if err != nil {
		e.log.Fields(logger.Fields{"base_id": baseID, "code": code}).Error(err, "[entity.CompareToken] svc.CoinGecko.GetHistoryCoinInfo failed")
		return nil, err
	}
	targetOhlc, err, code := e.svc.CoinGecko.GetHistoryCoinInfo(targetID, interval)
	if err != nil {
		e.log.Fields(logger.Fields{"target_id": targetID, "code": code}).Error(err, "[entity.CompareToken] svc.CoinGecko.GetHistoryCoinInfo failed")
		return nil, err
	}

	res := &response.CompareTokenReponseData{BaseCoin: baseCoin, TargetCoin: targetCoin}
	size := len(baseOhlc)
	if size > len(targetOhlc) {
		size = len(targetOhlc)
	}
	for i := 0; i < size; i++ {
		// ohlc format: [(time), (open), (high), (low), (close)]
		targetPrice := targetOhlc[i][1]
		if targetPrice == 0 {
			continue
		}
		ratio := baseOhlc[i][1] / targetPrice
		timeStr := time.UnixMilli(int64(baseOhlc[i][0])).Format("01-02")
		res.Ratios = append(res.Ratios, ratio)
		res.Times = append(res.Times, timeStr)
	}
	if size > 0 {
		res.From = time.UnixMilli(int64(baseOhlc[0][0])).Format("January 02, 2006")
		res.To = time.UnixMilli(int64(baseOhlc[len(baseOhlc)-1][0])).Format("January 02, 2006")
	}
	return res, nil
}

func (e *Entity) SetGuildDefaultTicker(req request.GuildConfigDefaultTickerRequest) error {
	return e.repo.GuildConfigDefaultTicker.UpsertOne(&model.GuildConfigDefaultTicker{
		Query:         req.Query,
		GuildID:       req.GuildID,
		DefaultTicker: req.DefaultTicker,
	})
}

func (e *Entity) GetGuildDefaultTicker(req request.GetGuildDefaultTickerRequest) (*response.GetGuildDefaultTickerResponse, error) {
	defaultTicker, err := e.repo.GuildConfigDefaultTicker.GetOneByGuildIDAndQuery(req.GuildID, req.Query)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID, "query": req.Query}).Error(err, "[entity.GetGuildDefaultTicker] repo.GuildConfigDefaultTicker.GetOneByGuildIDAndQuery() failed")
		return nil, err
	}
	return &response.GetGuildDefaultTickerResponse{
		Data: defaultTicker,
	}, nil
}

func (e *Entity) GetUserWatchlist(req request.GetUserWatchlistRequest) (*response.GetWatchlistResponse, error) {
	q := userwatchlistitem.UserWatchlistQuery{
		UserID: req.UserID,
		Offset: req.Page * req.Size,
		Limit:  req.Size,
	}
	list, _, err := e.repo.UserWatchlistItem.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetUserWatchlist] repo.UserWatchlistItem.List() failed")
		return nil, err
	}

	coinGeckoIds := make([]string, 0)
	pairs := make([]string, 0)
	for _, item := range list {
		if strings.Contains(item.Symbol, "/") {
			pairs = append(pairs, item.Symbol)
		}
		if item.CoinGeckoID == "" {
			continue
		}
		coinGeckoIds = append(coinGeckoIds, item.CoinGeckoID)
	}
	if len(coinGeckoIds) == 0 && len(pairs) == 0 {
		coinGeckoIds = e.getDefaultWatchlistIDs()
	}
	if len(coinGeckoIds) == 0 && len(pairs) == 0 {
		return &response.GetWatchlistResponse{Data: nil}, nil
	}

	// CoinGeckoAPI | get ticker market data
	data := make([]response.CoinMarketItemData, 0)
	if len(coinGeckoIds) > 0 {
		cgData, err, code := e.svc.CoinGecko.GetCoinsMarketData(coinGeckoIds)
		if err != nil {
			e.log.Fields(logger.Fields{"ids": coinGeckoIds, "code": code}).Error(err, "[entity.GetUserWatchlist] svc.CoinGecko.GetCoinsMarketData() failed")
			return nil, err
		}
		data = append(data, cgData...)
	}

	// BinanceAPI | get pair market data
	for _, pair := range pairs {
		klineData, err, code := e.svc.Binance.GetKlinesBySymbol(pair)
		if err != nil {
			e.log.Fields(logger.Fields{"pair": pair, "code": code}).Error(err, "[entity.GetUserWatchlist] svc.Binance.GetKlinesBySymbol() failed")
			return nil, err
		}
		symbol := strings.Replace(pair, "/", "", 1)
		item := response.CoinMarketItemData{
			Name:   pair,
			Symbol: symbol,
			IsPair: true,
		}
		for _, klineItem := range klineData {
			price, _ := strconv.ParseFloat(klineItem.OPrice, 64)
			item.SparkLineIn7d.Price = append(item.SparkLineIn7d.Price, price)
		}
		latestPrice := item.SparkLineIn7d.Price[len(item.SparkLineIn7d.Price)-1]
		oldPrice := item.SparkLineIn7d.Price[0]
		item.CurrentPrice = latestPrice
		item.PriceChangePercentage7dInCurrency = (latestPrice - oldPrice) / oldPrice * 100
		data = append(data, item)
	}
	return &response.GetWatchlistResponse{Data: data}, nil
}

func (e *Entity) getDefaultWatchlistIDs() []string {
	return []string{"bitcoin", "ethereum", "binancecoin", "fantom", "internet-computer", "solana", "avalanche-2", "matic-network"}
}

func (e *Entity) AddToWatchlist(req request.AddToWatchlistRequest) (*response.AddToWatchlistResponse, error) {
	isPair := false
	// e.g. btc/usdt
	if strings.Contains(req.Symbol, "/") {
		isPair = true
	}
	switch {
	case isPair:
		// binance klines API requires symbol to be uppercase
		req.Symbol = strings.ToUpper(req.Symbol)
		_, err, code := e.svc.Binance.GetExchangeInfo(req.Symbol)
		if err != nil {
			e.log.Fields(logger.Fields{"symbol": req.Symbol, "code": code}).Error(err, "[entity.AddToWatchlist] svc.Binance.GetExchangeInfo() failed")
			return nil, baseerrs.ErrRecordNotFound
		}
	case req.CoinGeckoID == "" && !isPair:
		coins, err := e.SearchCoins(req.Symbol)
		// coins, err, code := e.svc.CoinGecko.SearchCoins(req.Symbol)
		if err != nil {
			e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[entity.AddToWatchlist] svc.CoinGecko.SearchCoins() failed")
			return nil, err
		}
		if len(coins) > 1 {
			return &response.AddToWatchlistResponse{
				Data: &response.AddToWatchlistResponseData{Suggestions: coins},
			}, nil
		}
		if len(coins) == 0 {
			e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[entity.AddToWatchlist] svc.CoinGecko.SearchCoins() - no data found")
			return nil, baseerrs.ErrRecordNotFound
		}
		req.CoinGeckoID = coins[0].ID
	}

	listQ := userwatchlistitem.UserWatchlistQuery{CoinGeckoID: req.CoinGeckoID, UserID: req.UserID}
	if isPair {
		listQ.Symbol = req.Symbol
	}
	_, total, err := e.repo.UserWatchlistItem.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.AddToWatchlist] repo.UserWatchlistItem.List() failed")
		return nil, err
	}
	if total == 1 {
		return nil, baseerrs.ErrConflict
	}

	err = e.repo.UserWatchlistItem.Create(&model.UserWatchlistItem{
		UserID:      req.UserID,
		Symbol:      req.Symbol,
		CoinGeckoID: req.CoinGeckoID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.AddToWatchlist] repo.UserWatchlistItem.Create() failed")
		return nil, err
	}
	return &response.AddToWatchlistResponse{Data: nil}, nil
}

func (e *Entity) RemoveFromWatchlist(req request.RemoveFromWatchlistRequest) error {
	rows, err := e.repo.UserWatchlistItem.Delete(req.UserID, req.Symbol)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.RemoveFromWatchlist] repo.UserWatchlistItem.Delete() failed")
	}
	if rows == 0 {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.RemoveFromWatchlist] item not found")
		return baseerrs.ErrRecordNotFound
	}
	return err
}

func (e *Entity) RefreshCoingeckoSupportedTokensList() (int64, error) {
	tokens, err, code := e.svc.CoinGecko.GetSupportedCoins()
	if err != nil {
		e.log.Fields(logger.Fields{"code": code}).Error(err, "[entity.RefreshCoingeckoSupportedTokensList] svc.CoinGecko.GetSupportedCoins() failed")
		return 0, err
	}
	e.log.Infof("[entity.RefreshCoingeckoSupportedTokensList] svc.CoinGecko.GetSupportedCoins() - found %d items", len(tokens))

	updatedRows := int64(0)
	for _, token := range tokens {
		model := model.CoingeckoSupportedTokens{
			ID:     token.ID,
			Name:   token.Name,
			Symbol: token.Symbol,
		}
		rowsAffected, err := e.repo.CoingeckoSupportedTokens.Upsert(&model)
		if err != nil {
			e.log.Fields(logger.Fields{"token": token}).Error(err, "[entity.RefreshCoingeckoSupportedTokensList] repo.CoingeckoSupportedTokens.Upsert() failed")
			continue
		}
		updatedRows += rowsAffected
	}
	return updatedRows, nil
}
