package entities

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	onchaintipbottransaction "github.com/defipod/mochi/pkg/repo/onchain_tip_bot_transaction"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

// func (e *Entity) SubmitOnchainTransfer(req request.SubmitOnchainTransferRequest) ([]response.SubmitOnchainTransfer, error) {
// 	// validate transfer
// 	message := fmt.Sprintf("[Transfer onchain to %s] %s", req.Recipients, req.Message)
// 	amountEach := req.Amount / float64(len(req.Recipients))
// 	supportedToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(req.Token))
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
// 				UserID:          &req.Sender,
// 				GuildID:         &req.GuildID,
// 				ChannelID:       &req.ChannelID,
// 				Action:          &req.TransferType,
// 				Receiver:        req.Recipients,
// 				NumberReceivers: 1,
// 				Duration:        &req.Duration,
// 				Amount:          amountEach,
// 				Status:          consts.OffchainTipBotTrasferStatusFail,
// 				FullCommand:     &req.FullCommand,
// 				FailReason:      consts.OffchainTipBotFailReasonTokenNotSupported,
// 				Image:           req.Image,
// 				Message:         message,
// 			})
// 			return nil, errors.New(consts.OffchainTipBotFailReasonTokenNotSupported)
// 		}
// 		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entity.SubmitOnchainTransfer] repo.OffchainTipBotTokens.GetBySymbol() failed")
// 		return nil, err
// 	}

// 	insufficientBal := &model.OffchainTipBotActivityLog{
// 		UserID:          &req.Sender,
// 		GuildID:         &req.GuildID,
// 		ChannelID:       &req.ChannelID,
// 		Action:          &req.TransferType,
// 		Receiver:        req.Recipients,
// 		NumberReceivers: len(req.Recipients),
// 		TokenID:         supportedToken.ID.String(),
// 		Duration:        &req.Duration,
// 		Amount:          amountEach,
// 		Status:          consts.OffchainTipBotTrasferStatusFail,
// 		FullCommand:     &req.FullCommand,
// 		FailReason:      consts.OffchainTipBotFailReasonNotEnoughBalance,
// 		Image:           req.Image,
// 		Message:         message,
// 	}
// 	userBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(req.Sender, supportedToken.ID)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(insufficientBal)
// 			return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
// 		}
// 		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[entity.SubmitOnchainTransfer] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
// 		return nil, err
// 	}
// 	if req.All {
// 		req.Amount = userBal.Amount
// 		amountEach = req.Amount / float64(len(req.Recipients))
// 	}
// 	if userBal.Amount == 0 || float64(userBal.Amount) < req.Amount {
// 		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(insufficientBal)
// 		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
// 	}

// 	// update sender balanace
// 	batch := []model.OffchainTipBotUserBalance{{UserID: req.Sender, TokenID: supportedToken.ID, ChangedAmount: -req.Amount}}
// 	err = e.UpsertBatchOfUserBalances("tip-onchain", supportedToken.TokenSymbol, batch)
// 	if err != nil {
// 		e.log.Fields(logger.Fields{
// 			"totalAmount": req.Amount, "amountEach": amountEach, "sender": req.Sender,
// 			"recipients": len(req.Recipients), "token": supportedToken.TokenSymbol,
// 		}).Error(err, "[entity.SubmitOnchainTransfer] entity.UpsertBatchOfUserBalances() failed")
// 		return nil, err
// 	}

// 	// create pending transactions
// 	list := make([]*model.OnchainTipBotTransaction, len(req.Recipients))
// 	for i, r := range req.Recipients {
// 		list[i] = &model.OnchainTipBotTransaction{
// 			SenderDiscordID:    req.Sender,
// 			RecipientDiscordID: r,
// 			GuildID:            req.GuildID,
// 			ChannelID:          req.ChannelID,
// 			Amount:             amountEach,
// 			TokenSymbol:        req.Token,
// 			Each:               req.Each,
// 			All:                req.All,
// 			TransferType:       req.TransferType,
// 			FullCommand:        req.FullCommand,
// 			Message:            req.Message,
// 			Image:              req.Image,
// 			Status:             "pending",
// 		}
// 	}
// 	if err := e.repo.OnchainTipBotTransaction.UpsertMany(list); err != nil {
// 		e.log.Fields(logger.Fields{"list": list}).Error(err, "[entity.SubmitOnchainTransfer] repo.OnchainTipBotTransaction.UpsertMany() failed")
// 		return nil, err
// 	}

// 	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{supportedToken.CoinGeckoID}, "usd")
// 	if err != nil {
// 		e.log.Fields(logger.Fields{"token": supportedToken.CoinGeckoID}).Error(err, "[entity.SubmitOnchainTransfer] svc.CoinGecko.GetCoinPrice() failed")
// 		return nil, err
// 	}

// 	// dm recipient
// 	res := make([]response.SubmitOnchainTransfer, len(req.Recipients))
// 	for i, r := range req.Recipients {
// 		e.notifyPendingTransfer(r)
// 		res[i] = response.SubmitOnchainTransfer{
// 			SenderID:    req.Sender,
// 			RecipientID: r,
// 			Amount:      amountEach,
// 			Symbol:      supportedToken.TokenSymbol,
// 			AmountInUSD: amountEach * tokenPrice[supportedToken.CoinGeckoID],
// 		}
// 	}

// 	// notify tip to channel
// 	notifyReq := request.OffchainTransferRequest{
// 		Sender:       req.Sender,
// 		Recipients:   req.Recipients,
// 		Platform:     req.Platform,
// 		GuildID:      req.GuildID,
// 		ChannelID:    req.ChannelID,
// 		Amount:       req.Amount,
// 		Token:        req.Token,
// 		Each:         req.Each,
// 		All:          req.All,
// 		TransferType: req.TransferType,
// 		Image:        req.Image,
// 		Message:      req.Message,
// 	}
// 	e.sendLogNotify(notifyReq, 0)

// 	// notify tip to other platform: twitter, telegram, ...
// 	e.NotifyTipFromPlatforms(notifyReq, amountEach, tokenPrice[supportedToken.CoinGeckoID])

// 	return res, nil
// }

func (e *Entity) ClaimOnchainTransfer(req request.ClaimOnchainTransferRequest) (*response.ClaimOnchainTransfer, error) {
	tx, err := e.repo.OnchainTipBotTransaction.GetOnePending(req.ClaimID)
	if err != nil {
		e.log.Fields(logger.Fields{"claim_id": req.ClaimID}).Error(err, "[entity.ClaimOnchainTransfer] repo.OnchainTipBotTransaction.GetOne() failed")
		return nil, err
	}
	if req.UserID != tx.RecipientDiscordID {
		e.log.Fields(logger.Fields{"claim_id": req.ClaimID}).Error(err, "[entity.ClaimOnchainTransfer] requesting user is not authorized to claim")
		return nil, gorm.ErrRecordNotFound
	}
	message := fmt.Sprintf("[Transfer onchain to %s] %s", req.Address, tx.Message)
	token, err := e.repo.Token.GetBySymbol(tx.TokenSymbol, true)
	if err != nil {
		e.log.Fields(logger.Fields{"token": tx.TokenSymbol}).Error(err, "[entity.ClaimOnchainTransfer] repo.Token.GetBySymbol() failed")
		return nil, err
	}

	// send onchain tx
	onchainTx, _, err := e.transferOnchain(accounts.Account{Address: common.HexToAddress(req.Address)}, tx.Amount, token, -1, false)
	if err != nil {
		e.log.Fields(logger.Fields{"addr": req.Address, "amount": tx.Amount, "token": token.Symbol}).Error(err, "[entity.ClaimOnchainTransfer] e.transferOnchain() failed")
		return nil, err
	}

	// update onchain tx
	tx.ClaimedAt = time.Now().UTC()
	tx.Status = "claimed"
	tx.TxHash = onchainTx.Hash().Hex()
	tx.RecipientAddress = &req.Address
	err = e.repo.OnchainTipBotTransaction.UpsertMany([]*model.OnchainTipBotTransaction{tx})
	if err != nil {
		e.log.Fields(logger.Fields{"tx": tx}).Error(err, "[entity.SubmitOnchainTransfer] repo.OnchainTipBotTransaction.UpsertMany() failed")
		return nil, err
	}

	// create activity log
	supportedToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(tx.TokenSymbol))
	if err != nil {
		e.log.Fields(logger.Fields{"token": tx.TokenSymbol}).Error(err, "[entity.ClaimOnchainTransfer] repo.OffchainTipBotTokens.GetBySymbol() failed")
		return nil, err
	}
	log := &model.OffchainTipBotActivityLog{
		UserID:          &tx.SenderDiscordID,
		GuildID:         &tx.GuildID,
		ChannelID:       &tx.ChannelID,
		Action:          &tx.TransferType,
		Receiver:        []string{tx.RecipientDiscordID},
		TokenID:         supportedToken.ID.String(),
		NumberReceivers: 1,
		Duration:        &tx.Duration,
		Amount:          tx.Amount,
		Status:          consts.OffchainTipBotTrasferStatusSuccess,
		FullCommand:     &tx.FullCommand,
		FailReason:      "",
		Image:           tx.Image,
		Message:         message,
	}
	err = e.repo.OffchainTipBotActivityLogs.CreateActivityLog(log)
	if err != nil {
		e.log.Fields(logger.Fields{"tx": tx}).Error(err, "[entity.ClaimOnchainTransfer] repo.OffchainTipBotActivityLogs.CreateActivityLog() failed")
		return nil, err
	}

	_, err = e.repo.OffchainTipBotTransferHistories.CreateTransferHistories([]model.OffchainTipBotTransferHistory{
		{
			SenderID:   &tx.SenderDiscordID,
			ReceiverID: tx.RecipientDiscordID,
			GuildID:    tx.GuildID,
			LogID:      log.ID.String(),
			Status:     consts.OffchainTipBotTrasferStatusSuccess,
			Amount:     tx.Amount,
			Token:      supportedToken.TokenSymbol,
			Action:     tx.TransferType,
		},
	})
	if err != nil {
		e.log.Fields(logger.Fields{"tx": tx}).Error(err, "[entity.ClaimOnchainTransfer] repo.OffchainTipBotTransferHistories.CreateTransferHistories() failed")
		return nil, err
	}

	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{supportedToken.CoinGeckoID}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": supportedToken.CoinGeckoID}).Error(err, "[entity.ClaimOnchainTransfer] svc.CoinGecko.GetCoinPrice() failed")
		return nil, err
	}
	return &response.ClaimOnchainTransfer{
		SubmitOnchainTransfer: response.SubmitOnchainTransfer{
			SenderID:    tx.SenderDiscordID,
			RecipientID: tx.RecipientDiscordID,
			Amount:      tx.Amount,
			Symbol:      tx.TokenSymbol,
			AmountInUSD: tx.Amount * tokenPrice[supportedToken.CoinGeckoID],
		},
		RecipientAddress: req.Address,
		TxHash:           tx.TxHash,
		TxUrl:            fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, tx.TxHash),
	}, nil
}

func (e *Entity) notifyPendingTransfer(userID string) error {
	q := onchaintipbottransaction.ListQuery{RecipientDiscordID: userID, Status: "pending"}
	txs, err := e.repo.OnchainTipBotTransaction.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": q}).Error(err, "[entity.notifyPendingTransfer] repo.OnchainTipBotTransaction.List() failed")
		return err
	}
	dmChannel, err := e.discord.UserChannelCreate(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[entity.notifyPendingTransfer] discord.UserChannelCreate() failed")
		return err
	}
	if len(txs) == 0 {
		e.log.Fields(logger.Fields{"userID": userID}).Info("[entity.notifyPendingTransfer] no pending tx")
		return nil
	}

	claimIDs := make([]string, len(txs))
	amounts := make([]string, len(txs))
	senders := make([]string, len(txs))
	for i, tx := range txs {
		claimIDs[i] = fmt.Sprint(tx.ID)
		amounts[i] = fmt.Sprintf("%.2f %s", tx.Amount, tx.TokenSymbol)
		senders[i] = fmt.Sprintf("<@%s>", tx.SenderDiscordID)
	}
	description := fmt.Sprintf("<:pointingright:1058304352944656384> You have received %d tips. You can transfer each to your crypto wallet by `$claim <Claim ID> <your recipient address>`.", len(txs))
	description += "\n<:pointingright:1058304352944656384> You can find the recipient address in your crypto wallet (Eg: Metamask, Phantom, ...)."
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/933342303546929203.png?size=240&quality=lossless",
			Name:    "Claim your tip!",
		},
		Description: description,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Claim ID",
				Value:  strings.Join(claimIDs, "\n"),
				Inline: true,
			},
			{
				Name:   "Amount",
				Value:  strings.Join(amounts, "\n"),
				Inline: true,
			},
			{
				Name:   "Sender",
				Value:  strings.Join(senders, "\n"),
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Type /feedback to report•Mochi Bot • %s", time.Now().Format("2006-01-02 15:04:05")),
		},
		Color: 0x9FFFE4,
	}
	_, err = e.discord.ChannelMessageSendEmbeds(dmChannel.ID, []*discordgo.MessageEmbed{embed})
	if err != nil {
		e.log.Fields(logger.Fields{"dm": dmChannel.ID, "userID": userID}).Error(err, "[entity.notifyPendingTransfer] discord.ChannelMessageSendEmbeds() failed")
	}
	return err
}

func (e *Entity) GetUserOnchainTransfers(userID, status string) ([]model.OnchainTipBotTransaction, error) {
	q := onchaintipbottransaction.ListQuery{RecipientDiscordID: userID, Status: status}
	txs, err := e.repo.OnchainTipBotTransaction.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": q}).Error(err, "[entity.GetUserOnchainTransfers] repo.OnchainTipBotTransaction.List() failed")
		return nil, err
	}
	return txs, nil
}

func (e *Entity) OnchainBalances(address string, tokens []model.Token) (map[string]*big.Int, error) {
	balances := make(map[string]*big.Int, 0)
	for _, token := range tokens {
		chain := e.dcwallet.Chain(token.ChainID)
		if chain == nil {
			err := errors.New("chain not supported")
			e.log.Fields(logger.Fields{"chainID": token.ChainID}).Errorf(err, "[entity.OnchainBalances] %s", err.Error())
			continue
		}
		key := strings.ToUpper(token.Symbol)
		switch token.IsNative {
		case true:
			nativeBal, err := chain.RawNativeBalance(address, token)
			if err != nil {
				e.log.Fields(logger.Fields{"token": key}).Error(err, "[entity.OnchainBalances] chain.RawNativeBalance() failed")
				continue
			}
			balances[key] = nativeBal
		default:
			tokenBalance, err := chain.RawErc20TokenBalance(address, token)
			if err != nil {
				e.log.Fields(logger.Fields{"token": key}).Errorf(err, "[entity.OnchainBalances] chain.RawErc20TokenBalance() failed")
				continue
			}
			balances[key] = tokenBalance
		}
	}
	return balances, nil
}

func (e *Entity) GetPendingOnchainBalances(userID string) ([]response.GetUserBalances, error) {
	q := onchaintipbottransaction.ListQuery{RecipientDiscordID: userID, Status: "pending"}
	txs, err := e.repo.OnchainTipBotTransaction.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": q}).Error(err, "[entity.GetPendingOnchainBalances] repo.OnchainTipBotTransaction.List() failed")
		return nil, err
	}
	balanceMap := make(map[string]response.GetUserBalances)
	for _, tx := range txs {
		curBalance, ok := balanceMap[tx.TokenSymbol]
		if !ok {
			balanceMap[tx.TokenSymbol] = response.GetUserBalances{
				Symbol:   tx.TokenSymbol,
				Balances: tx.Amount,
			}
			continue
		}
		curBalance.Balances += tx.Amount
	}
	data := make([]response.GetUserBalances, 0, len(balanceMap))
	for symbol, uBal := range balanceMap {
		token, err := e.repo.Token.GetBySymbol(symbol, true)
		if err != nil {
			e.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[entity.GetPendingOnchainBalances] repo.Token.GetBySymbol() failed")
			return nil, err
		}
		if strings.EqualFold(token.Symbol, "icy") {
			token.CoinGeckoID = "icy"
		}
		uBal.Name = token.Name
		uBal.ID = token.CoinGeckoID
		tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{token.CoinGeckoID}, "usd")
		if err != nil {
			e.log.Fields(logger.Fields{"id": token.CoinGeckoID}).Error(err, "[entity.GetPendingOnchainBalances] repo.Token.GetBySymbol() failed")
			return nil, err
		}
		uBal.BalancesInUSD = uBal.Balances * tokenPrice[token.CoinGeckoID]
		data = append(data, uBal)
	}
	return data, nil
}
