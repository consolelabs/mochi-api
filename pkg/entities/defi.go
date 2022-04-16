package entities

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
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

const (
	getBeefyLPPriceURL    = `https://api.beefy.finance/lps`
	getBeefyTokenPriceURL = `https://api.beefy.finance/prices`
)

type HistoricalMarketChartResponse struct {
	Prices [][]float64 `json:"prices"`
}

type SearchCoinsResponse struct {
	Coins []CoinDataResponse `json:"coins"`
}

type CoinDataResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	MarketCapRank int    `json:"market_cap_rank"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}

type CoinPriceHistoryResponse struct {
	Timestamps []string  `json:"timestamps"`
	Prices     []float64 `json:"prices"`
	From       string    `json:"from"`
	To         string    `json:"to"`
}

func fetchHistoricalMarketData(req *request.GetMarketChartRequest, result interface{}) (int, error) {
	client := &http.Client{Timeout: time.Second * 60}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d", req.CoinID, req.Currency, req.Days)
	resp, err := client.Get(url)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, json.Unmarshal(bytes, result)
}

func searchForCoins(query, result interface{}) (int, error) {
	client := &http.Client{Timeout: time.Second * 60}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/search?query=%s", query)
	resp, err := client.Get(url)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, json.Unmarshal(bytes, result)
}

func (e *Entity) GetHistoricalMarketChart(c *gin.Context) (*CoinPriceHistoryResponse, error, int) {
	req, err := request.ValidateRequest(c)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	var resp *HistoricalMarketChartResponse
	statusCode, err := fetchHistoricalMarketData(req, &resp)
	if err != nil || statusCode != 200 {
		if statusCode != 404 {
			return nil, fmt.Errorf("failed to fetch historical market data - coin %s: %v", req.CoinID, err), statusCode
		}

		var searchResp *SearchCoinsResponse
		_, err := searchForCoins(req.CoinID, &searchResp)
		if err != nil || searchResp == nil || len(searchResp.Coins) == 0 {
			return nil, fmt.Errorf("failed to search for coins by query %s: %v", req.CoinID, err), http.StatusNotFound
		}
		req.CoinID = searchResp.Coins[0].ID
		statusCode, err := fetchHistoricalMarketData(req, &resp)
		if err != nil || statusCode != 200 {
			return nil, fmt.Errorf("failed to fetch historical market data 2 - coin %s: %v", req.CoinID, err), statusCode
		}
	}

	data := CoinPriceHistoryResponse{}
	for _, p := range resp.Prices {
		timestamp := time.UnixMilli(int64(p[0])).Format("01-02")
		data.Timestamps = append(data.Timestamps, timestamp)
		data.Prices = append(data.Prices, p[1])
	}

	from := time.UnixMilli(int64(resp.Prices[0][0])).Format("January 02, 2006")
	data.From = from
	to := time.UnixMilli(int64(resp.Prices[len(resp.Prices)-1][0])).Format("January 02, 2006")
	data.To = to

	return &data, nil, http.StatusOK
}

func (e *Entity) generateInDiscordWallet(user *model.User) error {
	if !user.InDiscordWalletAddress.Valid {
		inDiscordWalletNumber := e.repo.Users.GetLatestWalletNumber() + 1
		inDiscordAddress, err := e.dcwallet.GetAccountByWalletNumber(inDiscordWalletNumber)
		if err != nil {
			err = fmt.Errorf("error getting wallet address: %v", err)
			return err
		}

		user.InDiscordWalletNumber = model.JSONNullInt64{NullInt64: sql.NullInt64{Int64: int64(inDiscordWalletNumber), Valid: true}}
		user.InDiscordWalletAddress = model.JSONNullString{NullString: sql.NullString{String: inDiscordAddress.Address.Hex(), Valid: true}}

		if err = e.repo.Users.UpsertOne(*user); err != nil {
			err = fmt.Errorf("error upsert user: %v", err)
			return err
		}
	}

	return nil
}

func (e *Entity) InDiscordWalletTransfer(req request.TransferRequest) ([]response.InDiscordWalletTransferResponse, []string) {
	res := []response.InDiscordWalletTransferResponse{}
	errs := []string{}

	fromUser, err := e.repo.Users.GetOne(req.FromDiscordID)
	if err != nil {
		errs = append(errs, fmt.Sprintf("user not found: %v", err))
		return nil, errs
	}
	if err = e.generateInDiscordWallet(fromUser); err != nil {
		errs = append(errs, fmt.Sprintf("cannot generate in-discord wallet: %v", err))
		return nil, errs
	}

	toUsers, err := e.repo.Users.GetByDiscordIDs(req.ToDiscordIDs)
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

	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency))
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

		signedTx, txBaseURL, err := e.transfer(fromAcc, toAcc, amountEach, token, nonce, req.All)
		if err != nil {
			errs = append(errs, fmt.Sprintf("error transfer: %v", err))
			continue
		}
		nonce = int(signedTx.Nonce()) + 1
		transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()
		transferredAmount := float64(signedTx.Value().Int64()) / float64(math.Pow10(18))

		_, err = e.repo.DiscordBotTransaction.Create(model.DiscordBotTransaction{
			TxHash:        signedTx.Hash().Hex(),
			FromDiscordID: req.FromDiscordID,
			ToDiscordID:   toUser.ID,
			GuildID:       req.GuildID,
			ChannelID:     req.ChannelID,
			Amount:        transferredAmount,
			TokenID:       token.ID,
			Type:          "TRANSFER",
		})
		if err != nil {
			errs = append(errs, fmt.Sprintf("error create tx: %v", err))
			continue
		}

		res = append(res, response.InDiscordWalletTransferResponse{
			FromDiscordID:  req.FromDiscordID,
			ToDiscordID:    toUser.ID,
			Amount:         transferredAmount,
			Cryptocurrency: req.Cryptocurrency,
			TxHash:         signedTx.Hash().Hex(),
			TxUrl:          fmt.Sprintf("%s/%s", txBaseURL, signedTx.Hash().Hex()),
			TransactionFee: transactionFee,
		})
	}
	if len(errs) == 0 {
		errs = nil
	}

	return res, errs
}

func (e *Entity) InDiscordWalletWithdraw(req request.WithdrawRequest) (res response.InDiscordWalletWithdrawResponse, err error) {
	fromUser, err := e.repo.Users.GetOne(req.FromDiscordID)
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

	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency))
	if err != nil {
		err = fmt.Errorf("error getting token info: %v", err)
		return
	}

	signedTx, txBaseURL, err := e.transfer(fromAccount,
		accounts.Account{Address: common.HexToAddress(req.ToAddress)},
		req.Amount,
		token, -1, req.All)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return
	}

	_, err = e.repo.DiscordBotTransaction.Create(model.DiscordBotTransaction{
		TxHash:        signedTx.Hash().Hex(),
		FromDiscordID: req.FromDiscordID,
		ToAddress:     req.ToAddress,
		GuildID:       req.GuildID,
		ChannelID:     req.ChannelID,
		Amount:        req.Amount,
		TokenID:       token.ID,
		Type:          "WITHDRAW",
	})
	if err != nil {
		err = fmt.Errorf("error create tx: %v", err)
		return
	}

	withdrawalAmount := util.WeiToEther(signedTx.Value())
	transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()

	res = response.InDiscordWalletWithdrawResponse{
		FromDiscordId:    req.FromDiscordID,
		ToAddress:        req.ToAddress,
		Amount:           req.Amount,
		Cryptocurrency:   req.Cryptocurrency,
		TxHash:           signedTx.Hash().Hex(),
		TxURL:            fmt.Sprintf("%s/%s", txBaseURL, signedTx.Hash().Hex()),
		WithdrawalAmount: withdrawalAmount,
		TransactionFee:   transactionFee,
	}
	return
}

func (e *Entity) transfer(fromAccount accounts.Account, toAccount accounts.Account, amount float64, token model.Token, nonce int, all bool) (*types.Transaction, string, error) {
	var (
		signedTx  *types.Transaction
		txBaseURL string
		err       error
	)

	switch token.ChainID {
	case 250: // ftm
		signedTx, err = e.dcwallet.FTM().Transfer(
			fromAccount,
			toAccount,
			amount,
			token,
			nonce,
			all,
		)
		if err != nil {
			err = fmt.Errorf("error transfer: %v", err)
			return nil, "", err
		}
		txBaseURL = "https://ftmscan.com/tx"
	default:
		return nil, "", errors.New("cryptocurrency not supported")
	}

	return signedTx, txBaseURL, nil
}

func (e *Entity) GetBeefyTokenPrices() (map[string]float64, error) {

	tokenPrices := make(map[string]float64)

	// get token prices
	if err := util.FetchData(getBeefyTokenPriceURL, &tokenPrices); err != nil {
		return nil, err
	}

	// get lp prices
	if err := util.FetchData(getBeefyLPPriceURL, &tokenPrices); err != nil {
		return nil, err
	}

	res := make(map[string]float64)

	for k, v := range tokenPrices {
		res[k] = float64(v)
	}

	return res, nil
}

func (e *Entity) InDiscordWalletBalances(discordID string) (response response.UserBalancesResponse, err error) {
	var ua *model.User
	switch {
	case discordID != "":
		ua, err = e.repo.Users.GetOne(discordID)
	default:
		err = fmt.Errorf("discord_id is required")
		return
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("unverified user")
			return
		}
		err = fmt.Errorf("failed to get user address, err: %v", err)
		return
	}

	tokens, err := e.repo.Token.GetAllSupported()
	if err != nil {
		err = fmt.Errorf("failed to get supported tokens - err: %v", err)
	}

	if ua.InDiscordWalletAddress.String == "" {
		if err = e.generateInDiscordWallet(ua); err != nil {
			err = fmt.Errorf("cannot generate in-discord wallet: %v", err)
			return
		}
		return e.InDiscordWalletBalances(discordID)
	}

	balances, err := e.dcwallet.FTM().Balances(ua.InDiscordWalletAddress.String, tokens)
	if err != nil {
		err = fmt.Errorf("cannot get user balances: %v", err)
		return
	}
	response.Balances = balances

	tokenPrices, err := e.GetBeefyTokenPrices()
	if err != nil {
		err = fmt.Errorf("cannot get user balances: %v", err)
		return
	}

	response.BalancesInUSD = make(map[string]float64)
	for tokenSymbol, balance := range balances {
		tokenPriceInUSD, ok := tokenPrices[strings.ToUpper(tokenSymbol)]
		if !ok {
			// get wrap version
			tokenPriceInUSD, ok = tokenPrices[strings.ToUpper("w"+tokenSymbol)]
		}
		if !ok {
			err = fmt.Errorf("failed to get token prices: %v", err)
			continue
		}
		response.BalancesInUSD[tokenSymbol] = balance * tokenPriceInUSD
	}

	return
}

func (e *Entity) GetSupportedTokens() (tokens []model.Token, err error) {
	tokens, err = e.repo.Token.GetAllSupported()
	if err != nil {
		err = fmt.Errorf("failed to get supported tokens - err: %v", err)
		return
	}
	return
}
