package entities

import (
	"errors"
	"fmt"
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

func (e *Entity) SubmitOnchainTransfer(req request.SubmitOnchainTransferRequest) ([]response.SubmitOnchainTransfer, error) {
	// validate transfer
	message := fmt.Sprintf("[Transfer onchain to %s] %s", req.Recipients, req.Message)
	amountEach := req.Amount / float64(len(req.Recipients))
	supportedToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(req.Token))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
				UserID:          &req.Sender,
				GuildID:         &req.GuildID,
				ChannelID:       &req.ChannelID,
				Action:          &req.TransferType,
				Receiver:        req.Recipients,
				NumberReceivers: 1,
				Duration:        &req.Duration,
				Amount:          amountEach,
				Status:          consts.OffchainTipBotTrasferStatusFail,
				FullCommand:     &req.FullCommand,
				FailReason:      consts.OffchainTipBotFailReasonTokenNotSupported,
				Image:           req.Image,
				Message:         message,
			})
			return nil, errors.New(consts.OffchainTipBotFailReasonTokenNotSupported)
		}
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entity.SubmitOnchainTransfer] repo.OffchainTipBotTokens.GetBySymbol() failed")
		return nil, err
	}

	insufficientBal := &model.OffchainTipBotActivityLog{
		UserID:          &req.Sender,
		GuildID:         &req.GuildID,
		ChannelID:       &req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        req.Recipients,
		NumberReceivers: len(req.Recipients),
		TokenID:         supportedToken.ID.String(),
		Duration:        &req.Duration,
		Amount:          amountEach,
		Status:          consts.OffchainTipBotTrasferStatusFail,
		FullCommand:     &req.FullCommand,
		FailReason:      consts.OffchainTipBotFailReasonNotEnoughBalance,
		Image:           req.Image,
		Message:         message,
	}
	userBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(req.Sender, supportedToken.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(insufficientBal)
			return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
		}
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[entity.SubmitOnchainTransfer] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return nil, err
	}
	if req.All {
		req.Amount = userBal.Amount
		amountEach = req.Amount / float64(len(req.Recipients))
	}
	if userBal.Amount == 0 || float64(userBal.Amount) < req.Amount {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(insufficientBal)
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// update sender balanace
	err = e.repo.OffchainTipBotUserBalances.UpdateUserBalance(&model.OffchainTipBotUserBalance{UserID: req.Sender, TokenID: supportedToken.ID, Amount: userBal.Amount - req.Amount})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.SubmitOnchainTransfer] repo.OffchainTipBotUserBalances.UpdateUserBalance() failed")
		return nil, err
	}

	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{supportedToken.CoinGeckoID}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": supportedToken.CoinGeckoID}).Error(err, "[entity.SubmitOnchainTransfer] svc.CoinGecko.GetCoinPrice() failed")
		return nil, err
	}

	// create pending transactions
	list := make([]*model.OnchainTipBotTransaction, len(req.Recipients))
	res := make([]response.SubmitOnchainTransfer, len(req.Recipients))
	for i, r := range req.Recipients {
		// store new onchain tx
		list[i] = &model.OnchainTipBotTransaction{
			SenderDiscordID:    req.Sender,
			RecipientDiscordID: r,
			GuildID:            req.GuildID,
			ChannelID:          req.ChannelID,
			Amount:             amountEach,
			TokenSymbol:        req.Token,
			Each:               req.Each,
			All:                req.All,
			TransferType:       req.TransferType,
			FullCommand:        req.FullCommand,
			Message:            req.Message,
			Image:              req.Image,
			Status:             "pending",
		}
		// response data
		res[i] = response.SubmitOnchainTransfer{
			SenderID:    req.Sender,
			RecipientID: r,
			Amount:      amountEach,
			Symbol:      supportedToken.TokenSymbol,
			AmountInUSD: amountEach * tokenPrice[supportedToken.CoinGeckoID],
		}
	}
	if err := e.repo.OnchainTipBotTransaction.UpsertMany(list); err != nil {
		e.log.Fields(logger.Fields{"list": list}).Error(err, "[entity.SubmitOnchainTransfer] repo.OnchainTipBotTransaction.UpsertMany() failed")
		return nil, err
	}

	// dm recipient
	for _, r := range req.Recipients {
		e.notifyPendingTransfer(r)
	}

	// notify tip to channel
	notifyReq := request.OffchainTransferRequest{
		Sender:       req.Sender,
		Recipients:   req.Recipients,
		Platform:     req.Platform,
		GuildID:      req.GuildID,
		ChannelID:    req.ChannelID,
		Amount:       req.Amount,
		Token:        req.Token,
		Each:         req.Each,
		All:          req.All,
		TransferType: req.TransferType,
		FullCommand:  req.FullCommand,
		Duration:     req.Duration,
		Image:        req.Image,
		Message:      req.Message,
	}
	e.sendLogNotify(notifyReq, amountEach)

	// notify tip to other platform: twitter, telegram, ...
	e.NotifyTipFromPlatforms(notifyReq, amountEach, tokenPrice[supportedToken.CoinGeckoID])

	return res, nil
}

func (e *Entity) ClaimOnchainTransfer(req request.ClaimOnchainTransferRequest) (*response.ClaimOnchainTransfer, error) {
	tx, err := e.repo.OnchainTipBotTransaction.GetOne(req.ClaimID)
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
	onchainTx, _, err := e.transferOnchain(tx.Amount,
		accounts.Account{Address: common.HexToAddress(req.Address)},
		tx.Amount,
		token, -1, false)
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
	al, err := e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
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
	})
	if err != nil {
		e.log.Fields(logger.Fields{"tx": tx}).Error(err, "[entity.ClaimOnchainTransfer] repo.OffchainTipBotActivityLogs.CreateActivityLog() failed")
		return nil, err
	}

	_, err = e.repo.OffchainTipBotTransferHistories.CreateTransferHistories([]model.OffchainTipBotTransferHistory{
		{
			SenderID:   &tx.SenderDiscordID,
			ReceiverID: tx.RecipientDiscordID,
			GuildID:    tx.GuildID,
			LogID:      al.ID.String(),
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
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/933342303546929203.png?size=240&quality=lossless",
			Name:    "On-chain tippings!",
		},
		Description: fmt.Sprintf("<@%s>, you received %d tipping(s) from others.\nYou can claim each transfer by using\n`$claim <Claim ID> <your recipient address>`", userID, len(txs)),
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
				Name:   "Author",
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
