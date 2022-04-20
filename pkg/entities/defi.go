package entities

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

const (
	getBeefyLPPriceURL    = "https://api.beefy.finance/lps"
	getBeefyTokenPriceURL = "https://api.beefy.finance/prices"
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
	fmt.Println("req:", req)
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
			FromDiscordID: req.Sender,
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
			FromDiscordID:  req.Sender,
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

	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Cryptocurrency))
	if err != nil {
		err = fmt.Errorf("error getting token info: %v", err)
		return
	}

	signedTx, txBaseURL, err := e.transfer(fromAccount,
		accounts.Account{Address: common.HexToAddress(req.Recipients[0])},
		req.Amount,
		token, -1, req.All)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return
	}

	_, err = e.repo.DiscordBotTransaction.Create(model.DiscordBotTransaction{
		TxHash:        signedTx.Hash().Hex(),
		FromDiscordID: req.Sender,
		ToAddress:     req.Recipients[0],
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
		FromDiscordId:    req.Sender,
		ToAddress:        req.Recipients[0],
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
	if _, err := util.FetchData(getBeefyTokenPriceURL, &tokenPrices); err != nil {
		return nil, err
	}

	// get lp prices
	if _, err := util.FetchData(getBeefyLPPriceURL, &tokenPrices); err != nil {
		return nil, err
	}

	res := make(map[string]float64)

	for k, v := range tokenPrices {
		res[k] = float64(v)
	}

	return res, nil
}

func (e *Entity) InDiscordWalletBalances(discordID string) (*response.UserBalancesResponse, error) {
	response := &response.UserBalancesResponse{}

	user, err := e.repo.Users.GetOne(discordID)
	if err != nil {
		err = fmt.Errorf("failed to get user %s: %v", discordID, err)
		return nil, err
	}

	tokens, err := e.repo.Token.GetAllSupported()
	if err != nil {
		err = fmt.Errorf("failed to get supported tokens - err: %v", err)
	}

	if user.InDiscordWalletAddress.String == "" {
		if err = e.generateInDiscordWallet(user); err != nil {
			err = fmt.Errorf("cannot generate in-discord wallet: %v", err)
			return nil, err
		}
	}

	balances, err := e.dcwallet.FTM().Balances(user.InDiscordWalletAddress.String, tokens)
	if err != nil {
		err = fmt.Errorf("cannot get user balances: %v", err)
		return nil, err
	}
	response.Balances = balances

	tokenPrices, err := e.GetBeefyTokenPrices()
	if err != nil {
		err = fmt.Errorf("cannot get user balances: %v", err)
		return nil, err
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
