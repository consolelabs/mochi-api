package entities

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	usertokenpricealert "github.com/defipod/mochi/pkg/repo/user_token_price_alert"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetHistoricalMarketChart(req *request.GetMarketChartRequest) (*response.CoinPriceHistoryResponse, error, int) {
	data, err, statusCode := e.svc.CoinGecko.GetHistoricalMarketData(req)
	if err != nil {
		return nil, err, statusCode
	}
	// handle quest logs
	log := &model.QuestUserLog{
		UserID: req.DiscordID,
		Action: model.QuestAction(model.TICKER),
	}
	if err := e.UpdateUserQuestProgress(log); err != nil {
		e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.GetHistoricalMarketChart] entity.UpdateUserQuestProgress() failed")
	}
	return data, nil, http.StatusOK
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

// TODO: refactor
func (e *Entity) transferOnchain(toAccount accounts.Account, amount float64, token model.Token, nonce int, all bool) (*types.Transaction, float64, error) {
	chain := e.dcwallet.Chain(token.ChainID)
	if chain == nil {
		return nil, 0, errors.New("cryptocurrency not supported")
	}
	signedTx, amount, err := chain.TransferOnchain(
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

func (e *Entity) GetSupportedTokens() (tokens []model.Token, err error) {
	tokens, err = e.repo.Token.GetAllSupported()
	if err != nil {
		err = fmt.Errorf("failed to get supported tokens - err: %v", err)
		return
	}
	return
}

func (e *Entity) GetCoinData(coinID string) (*response.GetCoinResponse, error, int) {
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

func (e *Entity) GetGuildDefaultTicker(req request.GetGuildDefaultTickerRequest) (*model.GuildConfigDefaultTicker, error) {
	defaultTicker, err := e.repo.GuildConfigDefaultTicker.GetOneByGuildIDAndQuery(req.GuildID, req.Query)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID, "query": req.Query}).Error(err, "[entity.GetGuildDefaultTicker] repo.GuildConfigDefaultTicker.GetOneByGuildIDAndQuery() failed")
		return nil, err
	}
	return defaultTicker, nil
}

func (e *Entity) GetUserWatchlist(req request.GetUserWatchlistRequest) (*response.GetWatchlistResponse, error) {
	q := userwatchlistitem.UserWatchlistQuery{
		UserID: req.UserID,
		Offset: req.Page * req.Size,
		Limit:  req.Size,
	}
	list, total, err := e.repo.UserWatchlistItem.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetUserWatchlist] repo.UserWatchlistItem.List() failed")
		return nil, err
	}

	tickers := make([]string, 0)
	pairs := make([]model.UserWatchlistItem, 0)
	isDefault := false
	for _, item := range list {
		if strings.Contains(item.Symbol, "/") {
			pairs = append(pairs, item)
		}
		tickers = append(tickers, item.CoinGeckoID)
	}
	if len(tickers) == 0 && len(pairs) == 0 {
		tickers = e.getDefaultWatchlistIDs()
		isDefault = true
	}
	if len(tickers) == 0 && len(pairs) == 0 {
		return nil, nil
	}

	// CoinGeckoAPI | get ticker market data
	data := make([]response.CoinMarketItemData, 0)
	if len(tickers) > 0 {
		cgData, err, code := e.svc.CoinGecko.GetCoinsMarketData(tickers)
		if err != nil {
			e.log.Fields(logger.Fields{"ids": tickers, "code": code}).Error(err, "[entity.GetUserWatchlist] svc.CoinGecko.GetCoinsMarketData() failed")
			return nil, err
		}
		data = append(data, cgData...)
	}

	for _, pair := range pairs {
		queries := strings.Split(pair.CoinGeckoID, "/")
		item := response.CoinMarketItemData{}
		// process fiat data
		if pair.IsFiat {
			fiatData, err := e.svc.Nghenhan.GetFiatHistoricalChart(queries[0], queries[1], "h", 168)
			if err != nil {
				e.log.Fields(logger.Fields{"pair": pair}).Error(err, "[entity.GetUserWatchlist] Nghenhan.GetFiatHistoricalChart failed")
				return nil, err
			}
			fiatRate := []float64{}
			for _, v := range fiatData.Data {
				cPrice, _ := strconv.ParseFloat(v.ClosePrice, 64)
				fiatRate = append(fiatRate, cPrice)
			}
			if len(fiatRate) == 0 {
				e.log.Fields(logger.Fields{"pair": pair}).Error(err, "[entity.GetUserWatchlist] Nghenhan.GetFiatHistoricalChart returned no data")
				return nil, err
			}

			lastestPrice := fiatRate[len(fiatRate)-1]
			oldPrice := fiatRate[0]
			if oldPrice == 0 {
				item.PriceChangePercentage7dInCurrency = 100
			} else {
				item.PriceChangePercentage7dInCurrency = (lastestPrice - oldPrice) / oldPrice * 100
			}
			item.Symbol = pair.Symbol
			item.IsPair = true
			item.SparkLineIn7d.Price = fiatRate
			item.CurrentPrice = lastestPrice
		} else {
			comparisonData, err := e.CompareToken(queries[0], queries[1], "7", "")
			if err != nil {
				e.log.Fields(logger.Fields{"pair": pair}).Error(err, "[entity.GetUserWatchlist] entity.CompareToken() failed")
				return nil, err
			}

			item.Symbol = pair.Symbol
			item.IsPair = true
			item.Image = fmt.Sprintf("%s||%s", comparisonData.BaseCoin.Image.Small, comparisonData.TargetCoin.Image.Small)
			item.SparkLineIn7d.Price = comparisonData.Ratios
			if len(comparisonData.Ratios) > 0 {
				latestPrice := item.SparkLineIn7d.Price[len(item.SparkLineIn7d.Price)-1]
				oldPrice := item.SparkLineIn7d.Price[0]
				item.CurrentPrice = latestPrice
				item.PriceChangePercentage7dInCurrency = (latestPrice - oldPrice) / oldPrice * 100
			}
		}
		data = append(data, item)
	}
	for i := range data {
		data[i].IsDefault = isDefault
	}
	// handle quest logs
	log := &model.QuestUserLog{
		UserID: req.UserID,
		Action: model.QuestAction(model.WATCHLIST),
	}
	if err := e.UpdateUserQuestProgress(log); err != nil {
		e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.GetUserWatchlist] entity.UpdateUserQuestProgress() failed")
	}
	return &response.GetWatchlistResponse{
		Pagination: &response.PaginationResponse{
			Total: total,
			Pagination: model.Pagination{
				Page: int64(req.Page),
				Size: int64(req.Size),
			},
		},
		Data: data,
	}, nil
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
	case req.IsFiat:
		req.CoinGeckoID = req.Symbol

	case isPair && req.CoinGeckoID == "":
		queries := strings.Split(req.Symbol, "/")
		data, err := e.CompareToken(queries[0], queries[1], "7", "")
		if err != nil {
			e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[entity.AddToWatchlist] e.CompareToken() failed")
			return nil, baseerrs.ErrRecordNotFound
		}
		hasSuggestions := data.BaseCoinSuggestions != nil && len(data.BaseCoinSuggestions) > 0
		if hasSuggestions {
			return &response.AddToWatchlistResponse{
				Data: &response.AddToWatchlistResponseData{
					BaseSuggestions:   data.BaseCoinSuggestions,
					TargetSuggestions: data.TargetCoinSuggestions,
				},
			}, nil
		}
		req.CoinGeckoID = fmt.Sprintf("%s/%s", data.BaseCoin.ID, data.TargetCoin.ID)

	case !isPair && req.CoinGeckoID == "":
		tokens, err := e.SearchCoins(req.Symbol)
		// coins, err, code := e.svc.CoinGecko.SearchCoins(req.Symbol)
		if err != nil {
			e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[entity.AddToWatchlist] svc.CoinGecko.SearchCoins() failed")
			return nil, err
		}
		if len(tokens) > 1 {
			return &response.AddToWatchlistResponse{
				Data: &response.AddToWatchlistResponseData{BaseSuggestions: tokens},
			}, nil
		}
		if len(tokens) == 0 {
			e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[entity.AddToWatchlist] svc.CoinGecko.SearchCoins() - no data found")
			return nil, baseerrs.ErrRecordNotFound
		}
		req.CoinGeckoID = tokens[0].ID

	}

	listQ := userwatchlistitem.UserWatchlistQuery{CoinGeckoID: req.CoinGeckoID, UserID: req.UserID}
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
		IsFiat:      req.IsFiat,
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

func (e *Entity) GetFiatHistoricalExchangeRates(req request.GetFiatHistoricalExchangeRatesRequest) (*response.GetFiatHistoricalExchangeRatesResponse, error) {
	since := time.Now().AddDate(0, 0, -req.Days).UnixMilli()
	interval := "h"
	limit := req.Days * 24
	if req.Days > 90 {
		interval = "d"
		limit = req.Days
	}
	fiatData, err := e.svc.Nghenhan.GetFiatHistoricalChart(req.Base, req.Target, interval, limit)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetFiatHistoricalExchangeRates] Nghenhan.GetFiatHistoricalChart failed")
		return nil, err
	}

	rates := []float64{}
	times := []time.Time{}
	for _, v := range fiatData.Data {
		if v.OpenTime < int(since) {
			continue
		}
		// get price array
		cPrice, _ := strconv.ParseFloat(v.ClosePrice, 64)
		rates = append(rates, cPrice)
		// get time array
		t := time.Unix(0, int64(v.OpenTime)*int64(time.Millisecond))
		times = append(times, t)
	}
	if len(rates) == 0 || len(times) == 0 {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetFiatHistoricalExchangeRates] Nghenhan.GetFiatHistoricalChart returned no data")
		return nil, err
	}
	latest := rates[len(rates)-1]

	labels := []string{}
	for _, t := range times {
		date := t.Format("01-02")
		labels = append(labels, date)
	}

	return &response.GetFiatHistoricalExchangeRatesResponse{
		LatestRate: latest,
		Times:      labels,
		Rates:      rates,
		From:       times[0].Format("January 02, 2006"),
		To:         times[len(times)-1].Format("January 02, 2006"),
	}, nil
}

func (e *Entity) AddTokenPriceAlert(req request.AddTokenPriceAlertRequest) (*response.AddTokenPriceAlertResponse, error) {
	if req.Value <= 0 || req.Symbol == "" {
		e.log.Fields(logger.Fields{
			"price":  req.Value,
			"symbol": req.Symbol,
		}).Error(nil, "[Entity][AddTokenPriceAlert] invalid alert value or token symbol")
		return nil, baseerrs.ErrBadRequest
	}

	if err := req.AlertType.IsValidAlertType(); err != nil {
		e.log.Fields(logger.Fields{
			"alert_type": req.AlertType,
		}).Error(err, "[Entity][AddTokenPriceAlert] invalid alert type")
		return nil, err
	}

	if err := req.Frequency.IsValidAlertFrequency(); err != nil {
		e.log.Fields(logger.Fields{
			"frequency": req.Frequency,
		}).Error(err, "[Entity][AddTokenPriceAlert] invalid alert frequency")
		return nil, err
	}

	// check if req.Price is percentage
	isPercent := false
	req.Symbol = strings.ToUpper(req.Symbol)
	if req.AlertType == model.ChangeIsOver || req.AlertType == model.ChangeIsUnder {
		isPercent = true
		req.Value = util.RoundFloat(req.Value, 2)
	} else {
		req.Value = util.RoundFloat(req.Value, 8)
	}
	if req.Value == 0 {
		e.log.Fields(logger.Fields{"req.Value": req.Value}).Error(nil, "[entity.AddTokenPriceAlert] parse percentage value failed")
		return nil, baseerrs.ErrInvalidAlertValue
	}

	// fetch req.Symbol's current price
	var alertPair = req.Symbol + "USDT"
	var alertPrice = req.Value
	pairInfo, err, _ := e.svc.Binance.GetTickerPrice(alertPair)
	if err != nil {
		e.log.Fields(logger.Fields{"req.symbol": req.Symbol}).Error(err, "[entity.AddTokenPriceAlert] e.svc.Binance.GetTickerPrice() failed")
		return nil, baseerrs.ErrTokenNotFound
	}
	currentPrice, err := strconv.ParseFloat(pairInfo.Price, 64)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.AddTokenPriceAlert] strconv.ParseFloat() failed")
		return nil, err
	}

	// calculate trigger value if input value is percentage
	if isPercent {
		req.PriceByPercent = util.RoundFloat(currentPrice*req.Value/100+currentPrice, 8) // calculates target price to trigger alert
		alertPrice = req.PriceByPercent
	}

	// Generate Redis alert key based on Alert Type
	alertKey := req.AlertType.GetRedisKeyPrefix()
	if alertKey == "" {
		e.log.Fields(logger.Fields{"req.AlertType": req.AlertType}).Error(err, "[entity.AddTokenPriceAlert] req.AlertType.GetRedisKeyPrefix() failed")
		return nil, baseerrs.ErrBadRequest
	}
	alertKey = alertKey + ":" + strings.ToLower(alertPair)

	// check a price was already configured ?
	listQ := usertokenpricealert.UserTokenPriceAlertQuery{Symbol: req.Symbol, UserDiscordID: req.UserDiscordID, Value: req.Value}
	items, total, err := e.repo.UserTokenPriceAlert.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.AddTokenPriceAlert] repo.UserTokenPriceAlert.List() failed")
		return nil, err
	}

	// update other property if request price was already configured
	var alertID int
	if total >= 1 {
		var fetchedAlert = items[0]
		alertID = fetchedAlert.ID
		fetchedAlert.AlertType = req.AlertType
		fetchedAlert.Frequency = req.Frequency
		fetchedAlert.Value = req.Value
		fetchedAlert.PriceByPercent = req.PriceByPercent
		fetchedAlert.UpdatedAt = time.Now().UTC()
		err = e.repo.UserTokenPriceAlert.Update(&fetchedAlert)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.AddTokenPriceAlert] repo.UserTokenPriceAlert.Update() failed")
			return nil, err
		}
	} else {
		alertID, err = e.repo.UserTokenPriceAlert.Create(&model.UserTokenPriceAlert{
			UserDiscordID:  req.UserDiscordID,
			Symbol:         req.Symbol,
			AlertType:      req.AlertType,
			Frequency:      req.Frequency,
			Value:          req.Value,
			PriceByPercent: req.PriceByPercent,
			SnoozedTo:      time.Now().UTC(),
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      time.Now().UTC(),
			Currency:       "USDT",
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.AddTokenPriceAlert] repo.UserTokenPriceAlert.Create() failed")
			return nil, err
		}
	}

	err = e.cache.ZSet(alertKey, alertID, alertPrice)
	if err != nil {
		e.log.Fields(logger.Fields{"alertKey": alertKey, "alertID": alertID, "alertPrice": alertPrice}).Error(err, "[entity.AddTokenPriceAlert] e.cache.ZSet() failed")
		return nil, err
	}

	return &response.AddTokenPriceAlertResponse{Data: nil}, nil
}

func (e *Entity) GetUserListPriceAlert(req request.GetUserListPriceAlertRequest) (*[]model.UserTokenPriceAlert, error) {
	q := usertokenpricealert.UserTokenPriceAlertQuery{
		UserDiscordID: req.UserDiscordID,
		Offset:        req.Page * req.Size,
		Limit:         req.Size,
	}
	list, _, err := e.repo.UserTokenPriceAlert.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetUserListPriceAlert] repo.UserTokenPriceAlert.List() failed")
		return nil, err
	}
	return &list, nil
}

func (e *Entity) GetSpecificAlert(alertIDStr string) (*model.UserTokenPriceAlert, error) {
	alertID, err := strconv.Atoi(alertIDStr)
	if err != nil {
		e.log.Fields(logger.Fields{"alertID": alertIDStr}).Error(err, "[entidy.GetSpecificAlert] strconv.Atoi() failed")
		return nil, err
	}
	item, err := e.repo.UserTokenPriceAlert.GetById(alertID)
	if err == gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"alertID": alertID}).Error(err, "[entity.GetSpecificAlert] repo.UserTokenPriceAlert.GetOne() record not found")
		return nil, err
	}
	if err != nil {
		e.log.Fields(logger.Fields{"alertID": alertID}).Error(err, "[entity.GetSpecificAlert] repo.UserTokenPriceAlert.GetOne() failed")
		return nil, err
	}
	return &item, nil
}

func (e *Entity) UpdateSpecificPriceAlert(item model.UserTokenPriceAlert) error {
	err := e.repo.UserTokenPriceAlert.Update(&item)
	if err != nil {
		e.log.Fields(logger.Fields{}).Error(err, "[entity.UpdateSpecificPriceAlert] repo.UserTokenPriceAlert.Update() failed")
		return err
	}
	return nil
}

func (e *Entity) RemoveTokenPriceAlert(alertIDStr string) error {
	alertID, err := strconv.Atoi(alertIDStr)
	if err != nil {
		e.log.Fields(logger.Fields{"alertID": alertIDStr}).Error(err, "[entidy.RemoveTokenPriceAlert] strconv.Atoi() failed")
		return err
	}

	alert, err := e.repo.UserTokenPriceAlert.GetById(alertID)
	if err != nil {
		e.log.Fields(logger.Fields{"id": alertID}).Error(err, "[entity.RemoveTokenPriceAlert] repo.UserTokenPriceAlert.GetById() failed")
		return baseerrs.ErrRecordNotFound
	}

	err = e.repo.UserTokenPriceAlert.DeleteByID(alertID)
	if err != nil {
		e.log.Fields(logger.Fields{"id": alertID}).Error(err, "[entity.RemoveTokenPriceAlert] repo.UserTokenPriceAlert.Delete() failed")
	}

	var direction string
	if strings.Contains(alert.AlertType.GetRedisKeyPrefix(), "up") {
		direction = "up"
	} else {
		direction = "down"
	}

	if direction != "" && alert.PriceByPercent != 0 {
		err = e.RemovePriceAlertZCache(strings.ToLower(alert.Symbol+"USDT"), direction, fmt.Sprintf("%v", alert.PriceByPercent))
	} else {
		err = e.RemovePriceAlertZCache(strings.ToLower(alert.Symbol+"USDT"), direction, fmt.Sprintf("%v", alert.Value))
	}

	if err != nil {
		e.log.Fields(logger.Fields{"alert.Symbol": alert.Symbol, "alert.Value": alert.Value}).Error(err, "[entity.RemoveTokenPriceAlert] e.RemovePriceAlertZCache() failed")
	}

	return nil
}

func (e *Entity) GetBinanceCoinPrice(symbol string) (*response.GetTickerPriceResponse, error, int) {
	searchPair := strings.ToUpper(symbol + "usdt")
	data, err, statusCode := e.svc.Binance.GetTickerPrice(searchPair)
	if err != nil {
		e.log.Fields(logger.Fields{"req.Symbol": searchPair}).Error(err, "[entity.GetBinanceCoinData] e.svc.Binance.GetAvgPriceBySymbol() failed")
		return nil, baseerrs.ErrTokenNotFound, statusCode
	}

	return data, nil, http.StatusOK
}
