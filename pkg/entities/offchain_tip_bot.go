package entities

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo/offchain_tip_bot_contract"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) TransferToken(req request.OffchainTransferRequest) ([]response.OffchainTipBotTransferToken, error) {
	// check supported tokens
	log := model.OffchainTipBotActivityLog{
		UserID:          &req.Sender,
		GuildID:         &req.GuildID,
		ChannelID:       &req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        req.Recipients,
		NumberReceivers: len(req.Recipients),
		Duration:        &req.Duration,
		Amount:          req.Amount,
		FullCommand:     &req.FullCommand,
		Image:           req.Image,
		Message:         req.Message,
	}
	log.Status = consts.OffchainTipBotTrasferStatusFail
	log.FailReason = consts.OffchainTipBotFailReasonGetProfileFailed

	// get senderProfile, recipientProfiles by discordID
	mpReqs := request.MochiPayTransferRequest{}
	senderProfile, err := e.svc.MochiProfile.GetByDiscordID(req.Sender)
	if err != nil {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
		return []response.OffchainTipBotTransferToken{}, errors.New(log.FailReason)
	}

	mpReqs.From = &request.Wallet{
		ProfileGlobalId: senderProfile.ID,
	}

	for _, v := range req.Recipients {
		profile, err := e.svc.MochiProfile.GetByDiscordID(v)
		if err != nil {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
			return []response.OffchainTipBotTransferToken{}, errors.New(log.FailReason)
		}
		mpReqs.Tos = append(mpReqs.Tos, &request.Wallet{
			ProfileGlobalId: profile.ID,
		})
	}

	// validate amount
	if !req.All && req.Amount == 0 {
		log.Status = consts.OffchainTipBotTrasferStatusFail
		log.FailReason = consts.OffchainTipBotFailReasonInvalidAmount
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
		return []response.OffchainTipBotTransferToken{}, errors.New(log.FailReason)
	}

	// validate token
	supportedToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(req.Token))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Status = consts.OffchainTipBotTrasferStatusFail
			log.FailReason = consts.OffchainTipBotFailReasonTokenNotSupported
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
			return []response.OffchainTipBotTransferToken{}, errors.New(log.FailReason)
		}
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entity.TransferToken] repo.OffchainTipBotTokens.GetBySymbol() failed")
		return nil, err
	}

	mpReqs.TokenId = supportedToken.ID.String()
	mpReqs.Note = req.Message

	// validate balance
	log.TokenID = supportedToken.ID.String()
	insufficientBalanceLog := log
	insufficientBalanceLog.Status = consts.OffchainTipBotTrasferStatusFail
	insufficientBalanceLog.FailReason = consts.OffchainTipBotFailReasonNotEnoughBalance
	senderBalance, err := e.svc.MochiPay.GetBalance(senderProfile.ID, req.Token, "0")
	if err != nil {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&insufficientBalanceLog)
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[entity.TransferToken] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return nil, err
	}

	amount, err := util.StringToBigInt(senderBalance.Data[0].Amount)
	if err != nil {
		log.FailReason = consts.OffchainTipBotFailReasonInvalidAmount
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
		return []response.OffchainTipBotTransferToken{}, errors.New(log.FailReason)
	}

	if amount.Cmp(big.NewInt(0)) != 1 {
		log.FailReason = consts.OffchainTipBotFailReasonInvalidAmount
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
		return []response.OffchainTipBotTransferToken{}, errors.New(log.FailReason)
	}

	// case transfer all
	intVal, _ := util.BigIntToFloat(amount, float64(senderBalance.Data[0].Token.Decimal))
	if req.All {
		req.Amount = intVal
	}

	amountEachRecipient := req.Amount / float64(len(req.Recipients))
	insufficientBalanceLog.Amount = amountEachRecipient
	mpReqs.Amount = make([]string, len(req.Recipients))
	for i := range mpReqs.Amount {
		mpReqs.Amount[i] = strconv.FormatFloat(amountEachRecipient, 'f', 4, 64)
	}

	err = e.svc.MochiPay.Transfer(mpReqs)
	if err != nil {
		log.Status = consts.OffchainTipBotTrasferStatusFail
		log.FailReason = consts.OffchainTipBotFailReasonMochiPayTransferFailed
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
		return []response.OffchainTipBotTransferToken{}, errors.New(insufficientBalanceLog.FailReason)
	}

	// logs + histories
	log.Amount = amountEachRecipient
	log.Status = consts.OffchainTipBotTrasferStatusSuccess
	log.FailReason = ""
	err = e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TransferToken] repo.OffchainTipBotActivityLogs.CreateActivityLog() failed")
	}

	listTransferHistories := make([]model.OffchainTipBotTransferHistory, 0)
	for _, recipient := range req.Recipients {
		listTransferHistories = append(listTransferHistories, model.OffchainTipBotTransferHistory{
			SenderID:   &req.Sender,
			ReceiverID: recipient,
			GuildID:    req.GuildID,
			LogID:      log.ID.String(),
			Status:     consts.OffchainTipBotTrasferStatusSuccess,
			Amount:     amountEachRecipient,
			Token:      supportedToken.TokenSymbol,
			Action:     req.TransferType,
		})
	}

	transferHistories, err := e.repo.OffchainTipBotTransferHistories.CreateTransferHistories(listTransferHistories)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TransferToken] repo.OffchainTipBotTransferHistories.CreateTransferHistories() failed")
	}

	// notify tip to channel
	e.sendLogNotify(req, amountEachRecipient)

	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{supportedToken.CoinGeckoID}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": supportedToken.CoinGeckoID}).Error(err, "[entity.TransferToken] svc.CoinGecko.GetCoinPrice() failed")
	}

	// notify tip to other platform: twitter, telegram, ...
	e.NotifyTipFromPlatforms(req, amountEachRecipient, tokenPrice[supportedToken.CoinGeckoID])

	e.sendTipBotLogs(log, supportedToken.TokenSymbol, "")

	return e.MappingTransferTokenResponse(req.Token, amountEachRecipient, tokenPrice[supportedToken.CoinGeckoID], transferHistories), nil
}

func (e *Entity) sendLogNotify(req request.OffchainTransferRequest, amountEachRecipient float64) {
	if req.TransferType != consts.OffchainTipBotTransferTypeTip {
		return
	}
	// Do not return error here, just log it
	configNotifyChannels, err := e.repo.OffchainTipBotConfigNotify.GetByGuildID(req.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID}).Error(err, "[entity.sendLogNotify] repo.OffchainTipBotConfigNotify.GetByGuildID() failed")
		return
	}

	for _, configNotifyChannel := range configNotifyChannels {
		if strings.EqualFold(req.Token, configNotifyChannel.Token) || configNotifyChannel.Token == "*" {
			var recipients []string
			for _, recipient := range req.Recipients {
				recipients = append(recipients, fmt.Sprintf("<@%s>", recipient))
			}
			recipientsStr := strings.Join(recipients, ", ")
			descriptionFormat := "<@%s> has sent %s **%g %s** at <#%s>"
			if len(req.Recipients) > 1 {
				descriptionFormat = "<@%s> has sent %s **%g %s** each at <#%s>"
			}
			description := fmt.Sprintf(descriptionFormat, req.Sender, recipientsStr, amountEachRecipient, strings.ToUpper(req.Token), req.ChannelID)
			if req.Message != "" {
				description += fmt.Sprintf(" with messge\n\n  <:conversation:1032608818930139249> **%s**", req.Message)
			}
			title := fmt.Sprintf("<:tip:933384794627248128> %s <:tip:933384794627248128>", strings.ToUpper(req.TransferType))

			err := e.svc.Discord.SendTipActivityLogs(configNotifyChannel.ChannelID, req.Sender, title, description, req.Image)
			if err != nil {
				e.log.Fields(logger.Fields{"channel_id": configNotifyChannel.ChannelID}).Error(err, "[entity.sendLogNotify] discord.ChannelMessageSendEmbed() failed")
			}
		}
	}
}

func (e *Entity) NotifyTipFromPlatforms(req request.OffchainTransferRequest, amountEachRecipient float64, price float64) {
	// TODO(trkhoi): handle for all platforms
	if req.Platform == "telegram" {
		// send DM to all recipients
		for _, recipient := range req.Recipients {
			recipientBals, err := e.GetUserBalances(recipient)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req, "recipient": recipient}).Error(err, "[entity.NotifyTipFromPlatforms] repo.OffchainTipBotUserBalances.GetUserBalances() failed")
				return
			}

			dmChannel, err := e.discord.UserChannelCreate(recipient)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req, "recipient": recipient}).Error(err, "[entity.NotifyTipFromPlatforms] discord.UserChannelCreate() failed")
				return
			}

			balsEmbed := make([]*discordgo.MessageEmbedField, 0)

			for _, bal := range recipientBals {
				balsEmbed = append(balsEmbed, &discordgo.MessageEmbedField{
					Name:   bal.Name,
					Inline: true,
					Value:  fmt.Sprintf("%s %.2f %s `$%.2f`", util.GetEmoji(bal.Symbol), bal.Balances, bal.Symbol, bal.BalancesInUSD),
				})
			}

			msgTip := []*discordgo.MessageEmbed{
				{
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: "https://i.imgur.com/qj7iPqz.png",
					},
					Author: &discordgo.MessageEmbedAuthor{
						IconURL: "https://cdn.discordapp.com/emojis/942088817391849543.png?size=240&quality=lossless",
						Name:    fmt.Sprintf("Tips from %s", req.Platform),
					},
					Description: fmt.Sprintf("<@!%s> has sent you **%.2f %s** (\u2248 $%.2f)", req.Sender, amountEachRecipient, strings.ToUpper(req.Token), amountEachRecipient*price),
					Footer: &discordgo.MessageEmbedFooter{
						Text: fmt.Sprintf("Type /feedback to report•Mochi Bot • %s", time.Now().Format("2006-01-02 15:04:05")),
					},
					Color: 0x9fffe4,
				},
				{
					Author: &discordgo.MessageEmbedAuthor{
						IconURL: "https://cdn.discordapp.com/emojis/933342303546929203.png?size=240&quality=lossless",
						Name:    "Your balances",
					},
					Footer: &discordgo.MessageEmbedFooter{
						Text: fmt.Sprintf("Type /feedback to report•Mochi Bot • %s", time.Now().Format("2006-01-02 15:04:05")),
					},
					Color:  0x9fffe4,
					Fields: balsEmbed,
				},
				{
					Description: "You can now link telegram with your discord account. Type `$telegram config <telegram_username>` to link your account",
					Footer: &discordgo.MessageEmbedFooter{
						Text: fmt.Sprintf("Type /feedback to report•Mochi Bot • %s", time.Now().Format("2006-01-02 15:04:05")),
					},
					Color: 0x9fffe4,
				},
			}
			_, err = e.discord.ChannelMessageSendEmbeds(dmChannel.ID, msgTip)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req, "recipient": recipient}).Error(err, "[entity.NotifyTipFromPlatforms] discord.ChannelMessageSendEmbeds() failed")
				return
			}
		}
	}
}

func (e *Entity) MappingTransferTokenResponse(tokenSymbol string, amount float64, price float64, transferHistories []model.OffchainTipBotTransferHistory) (res []response.OffchainTipBotTransferToken) {
	for _, transferHistory := range transferHistories {
		res = append(res, response.OffchainTipBotTransferToken{
			SenderID:    *transferHistory.SenderID,
			RecipientID: transferHistory.ReceiverID,
			Amount:      amount,
			Symbol:      tokenSymbol,
			AmountInUSD: amount * price,
		})
	}
	return res
}

func (e *Entity) OffchainTipBotWithdraw(req request.OffchainWithdrawRequest) (*response.OffchainTipBotWithdraw, error) {
	al := model.OffchainTipBotActivityLog{
		UserID:          &req.Recipient,
		GuildID:         &req.GuildID,
		ChannelID:       &req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        []string{req.RecipientAddress},
		NumberReceivers: 1,
		Duration:        &req.Duration,
		Amount:          req.Amount,
		FullCommand:     &req.FullCommand,
	}
	// check supported tokens
	offchainToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(req.Token)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entity.OffchainTipBotWithdraw] repo.OffchainTipBotTokens.GetBySymbol() failed")
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		al.Status = consts.OffchainTipBotTrasferStatusFail
		al.FailReason = consts.OffchainTipBotFailReasonTokenNotSupported
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&al)
		return nil, errors.New(consts.OffchainTipBotFailReasonTokenNotSupported)
	}
	al.ServiceFee = offchainToken.ServiceFee
	al.FeeAmount = offchainToken.ServiceFee * req.Amount
	al.TokenID = offchainToken.ID.String()

	// check recipient balance
	insufficientBalanceLog := al
	insufficientBalanceLog.Status = consts.OffchainTipBotTrasferStatusFail
	insufficientBalanceLog.FailReason = consts.OffchainTipBotFailReasonNotEnoughBalance
	recipientBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(req.Recipient, offchainToken.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Recipient}).Error(err, "[entity.OffchainTipBotWithdraw] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&insufficientBalanceLog)
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// TODO: need another approach
	if req.All {
		req.Amount = util.RoundFloat(recipientBal.Amount/(1+offchainToken.ServiceFee), 5)
		// avoid round up lead to insufficent bal
		if float64(recipientBal.Amount) < req.Amount+offchainToken.ServiceFee*req.Amount {
			req.Amount = req.Amount - 0.00001
		}
		al.FeeAmount = offchainToken.ServiceFee * req.Amount
	}

	if float64(recipientBal.Amount) < req.Amount+offchainToken.ServiceFee*req.Amount {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&insufficientBalanceLog)
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// temp get from old tokens because not have flow from coingecko to offchain_tip_bot_tokens yet
	token, err := e.repo.Token.GetBySymbol(req.Token, true)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entity.OffchainTipBotWithdraw] repo.Token.GetBySymbol() failed")
		return nil, err
	}

	// execute tx
	var (
		withdrawalAmount  *big.Float
		transactionFee    float64
		transferredAmount float64
		txHash            string
		txErr             error
	)
	if strings.ToLower(req.Token) == "sol" {
		txHash, transferredAmount, txErr = e.solana.Transfer(e.cfg.SolanaCentralizedWalletPrivateKey, req.RecipientAddress, req.Amount, false)
	} else {
		var signedTx *types.Transaction
		signedTx, transferredAmount, txErr = e.transferOnchain(accounts.Account{Address: common.HexToAddress(req.RecipientAddress)}, req.Amount, token, -1, false)
		if signedTx != nil {
			txHash = signedTx.Hash().Hex()
			withdrawalAmount = util.WeiToEther(signedTx.Value())
			transactionFee, _ = util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()
		}
	}
	if txErr != nil {
		failedTxLog := al
		failedTxLog.Status = consts.OffchainTipBotTrasferStatusFail
		failedTxLog.FailReason = fmt.Sprintf("%s [%s]", txErr.Error(), txHash)
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&failedTxLog)
		e.log.Fields(logger.Fields{"txHash": txHash, "amount": transferredAmount, "token": req.Token}).Error(txErr, "[entity.OffchainTipBotWithdraw] transfer tx failed")
		return nil, txErr
	}

	// update author balance
	batch := []model.OffchainTipBotUserBalance{{UserID: req.Recipient, TokenID: offchainToken.ID, ChangedAmount: -transferredAmount - offchainToken.ServiceFee*req.Amount}}
	err = e.UpsertBatchOfUserBalances(req.TransferType, offchainToken.TokenSymbol, batch)
	if err != nil {
		e.log.Fields(logger.Fields{"batch": batch}).Error(err, "[entity.OffchainTipBotWithdraw] entity.UpsertBatchOfUserBalances() failed")
		return nil, err
	}

	// execute tx success -> create offchain_tip_bot_activity_logs + offchain_tip_bot_transfer_histories and update offchain_tip_bot_user_balances
	successLog := al
	successLog.Status = consts.OffchainTipBotTrasferStatusSuccess
	successLog.Amount = transferredAmount
	err = e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&successLog)
	if err == nil {
		e.repo.OffchainTipBotTransferHistories.CreateTransferHistories([]model.OffchainTipBotTransferHistory{{
			SenderID:   &req.Recipient,
			ReceiverID: req.Recipient,
			GuildID:    req.GuildID,
			LogID:      successLog.ID.String(),
			Status:     consts.OffchainTipBotTrasferStatusSuccess,
			Amount:     transferredAmount,
			Token:      offchainToken.TokenSymbol,
			Action:     req.TransferType,
			ServiceFee: offchainToken.ServiceFee,
			FeeAmount:  offchainToken.ServiceFee * req.Amount,
			TxHash:     txHash,
		}})
	}

	e.sendTipBotLogs(successLog, offchainToken.TokenSymbol, "")

	return &response.OffchainTipBotWithdraw{
		UserDiscordID:  req.Recipient,
		ToAddress:      req.RecipientAddress,
		Amount:         transferredAmount,
		Symbol:         req.Token,
		TxHash:         txHash,
		TxUrl:          fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, txHash),
		WithdrawAmount: withdrawalAmount,
		TransactionFee: transactionFee,
	}, nil
}

func (e *Entity) TotalBalances() ([]response.TotalBalances, error) {
	tokens, err := e.repo.Token.GetAll()
	if err != nil {
		e.log.Error(err, "[entities.migrateBalance] - failed to get supported tokens")
		return nil, err
	}
	balances, err := e.balances("0x4ec16127E879464bEF6ab310084FAcEC1E4Fe465", tokens)
	if err != nil {
		e.log.Error(err, "[entities.migrateBalance] - failed to get balance")
		return nil, err
	}

	totalBalances := make([]response.TotalBalances, 0)

	for key, value := range balances {
		coingeckoID := ""
		// select coingecko id
		for _, t := range tokens {
			if t.Symbol == key {
				coingeckoID = t.CoinGeckoID
			}
		}
		// get coingecko price
		var tokenPrice float64

		tokenPrices, err := e.svc.CoinGecko.GetCoinPrice([]string{coingeckoID}, "usd")
		if err != nil {
			e.log.Fields(logger.Fields{"token": key}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
			return nil, err
		}
		if len(tokenPrices) > 0 {
			tokenPrice = tokenPrices[coingeckoID]
		} else {
			tokenPrice = 1.0
		}

		if value > 0 {
			totalBalances = append(totalBalances, response.TotalBalances{
				Symbol:      key,
				Amount:      value,
				AmountInUsd: value * tokenPrice,
			})
		}

	}

	return totalBalances, nil
}

func (e *Entity) TotalOffchainBalances() ([]response.TotalOffchainBalances, error) {
	// cal exist offchain balance
	totalOffchainBalances, err := e.repo.OffchainTipBotUserBalances.SumAmountByTokenId()
	if err != nil {
		e.log.Error(err, "[e.repo.TotalOffchainBalances] - failed to get total offchain balance")
		return nil, err
	}

	mappingTotalOffchainBalances := make([]response.TotalOffchainBalances, 0)

	for _, totalOffchainBalance := range totalOffchainBalances {
		// get coingecko price
		var tokenPrice float64
		tokenPrices, err := e.svc.CoinGecko.GetCoinPrice([]string{totalOffchainBalance.CoinGeckoId}, "usd")
		if err != nil {
			e.log.Fields(logger.Fields{"token": totalOffchainBalance.CoinGeckoId}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
			return nil, err
		}
		if len(tokenPrices) > 0 {
			tokenPrice = tokenPrices[totalOffchainBalance.CoinGeckoId]
		} else {
			tokenPrice = 1.0
		}

		mappingTotalOffchainBalances = append(mappingTotalOffchainBalances, response.TotalOffchainBalances{
			Symbol:      totalOffchainBalance.TokenSymbol,
			Amount:      totalOffchainBalance.Total,
			AmountInUsd: totalOffchainBalance.Total * tokenPrice,
		})
	}
	return mappingTotalOffchainBalances, nil
}

func (e *Entity) TotalFee() ([]response.TotalFeeWithdraw, error) {
	return e.repo.OffchainTipBotTransferHistories.TotalFeeFromWithdraw()
}

func (e *Entity) UpdateTokenFee(req request.OffchainUpdateTokenFee) error {
	return e.repo.OffchainTipBotTokens.UpdateTokenFee(req.Symbol, req.ServiceFee)
}

func (e *Entity) GetAllTipBotTokens() ([]model.OffchainTipBotToken, error) {
	return e.repo.OffchainTipBotTokens.GetAll()
}

func (e *Entity) GetContracts(req request.TipBotGetContractsRequest) ([]model.OffchainTipBotContract, error) {
	contracts, err := e.repo.OffchainTipBotContract.List(offchain_tip_bot_contract.ListQuery{ChainID: req.ChainID, IsEVM: req.IsEVM, SupportDeposit: req.SupportDeposit})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetContracts] repo.OffchainTipBotContract.List() failed")
		return nil, err
	}
	for i, c := range contracts {
		chain, err := e.repo.OffchainTipBotChain.GetByID(c.ChainID)
		if err != nil {
			e.log.Fields(logger.Fields{"chainID": c.ChainID}).Error(err, "[entity.GetContracts] repo.OffchainTipBotChain.GetByID() failed")
			return nil, err
		}
		contracts[i].Chain = &chain
	}
	return contracts, nil
}

func (e *Entity) HandleIncomingDeposit(req request.TipBotDepositRequest) error {
	e.log.Fields(logger.Fields{"txHash": req.TxHash, "chainID": req.ChainID}).Info("receiving new deposit ...")
	chain, err := e.repo.OffchainTipBotChain.GetByChainID(req.ChainID)
	if err != nil {
		e.log.Fields(logger.Fields{"chainID": req.ChainID}).Error(err, "[entity.HandleIncomingDeposit] repo.OffchainTipBotChain.GetByChainID() failed")
		return err
	}
	if req.TokenSymbol == "" {
		req.TokenSymbol = chain.Currency
	}
	signedAtUnixMs := req.SignedAt.UnixMilli()
	getContractAssignmentQ := offchain_tip_bot_contract.GetAssignContractQuery{Address: req.ToAddress, TokenSymbol: req.TokenSymbol, SignedAt: &signedAtUnixMs}
	assignedContract, err := e.repo.OffchainTipBotContract.GetAssignContract(getContractAssignmentQ)
	if err != nil || assignedContract.UserID == "" {
		e.log.Fields(logger.Fields{"getContractAssignmentQ": getContractAssignmentQ}).Error(err, "[entity.HandleIncomingDeposit] repo.OffchainTipBotContract.GetAssignContract() failed")
		return nil
	}

	token, err := e.repo.Token.GetBySymbol(req.TokenSymbol, true)
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": req.TokenSymbol}).Error(err, "[entity.HandleIncomingDeposit] repo.Token.GetBySymbol() failed")
		return err
	}
	offchainToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(req.TokenSymbol)
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": req.TokenSymbol}).Error(err, "[entity.HandleIncomingDeposit] repo.OffchainTipBotTokens.GetBySymbol() failed")
		return err
	}
	tokenID := offchainToken.ID
	transferType := "deposit"
	userID := assignedContract.UserID
	// if erc-20 token OR solana => no calculation
	amount := req.Amount
	// else native token -> calculate
	if req.TokenContract == "" && !strings.EqualFold(req.TokenSymbol, "SOL") {
		amount /= math.Pow10(token.Decimals)
	}

	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{token.CoinGeckoID}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token.CoinGeckoID": token.CoinGeckoID}).Error(err, "[entity.HandleIncomingDeposit] svc.CoinGecko.GetCoinPrice() failed")
		return err
	}
	priceInUSD := tokenPrice[offchainToken.CoinGeckoID]
	// create deposit log
	depositLog := model.OffchainTipBotDepositLog{
		ChainID:     assignedContract.ChainID,
		TxHash:      req.TxHash,
		TokenID:     offchainToken.ID,
		FromAddress: req.FromAddress,
		ToAddress:   req.ToAddress,
		Amount:      amount,
		AmountInUSD: priceInUSD * amount,
		UserID:      userID,
		BlockNumber: req.BlockNumber,
		SignedAt:    req.SignedAt,
	}
	err = e.repo.OffchainTipBotDepositLog.CreateMany([]model.OffchainTipBotDepositLog{depositLog})
	if err != nil {
		e.log.Fields(logger.Fields{"depositLog": depositLog}).Error(err, "[entity.HandleIncomingDeposit] repo.OffchainTipBotDepositLog.CreateMany() failed")
		return err
	}

	// create activity log
	log := model.OffchainTipBotActivityLog{
		Action:          &transferType,
		Receiver:        []string{userID},
		TokenID:         tokenID.String(),
		NumberReceivers: 1,
		Amount:          amount,
		Status:          consts.OffchainTipBotTrasferStatusSuccess,
		FailReason:      "",
		Message:         req.TxHash,
		UserID:          &userID,
	}

	// update author balance
	batch := []model.OffchainTipBotUserBalance{{UserID: userID, TokenID: tokenID, ChangedAmount: amount, Amount: amount}}
	err = e.UpsertBatchOfUserBalances(transferType, offchainToken.TokenSymbol, batch)
	if err != nil {
		e.log.Fields(logger.Fields{"batch": batch}).Error(err, "[entity.HandleIncomingDeposit] entity.UpsertBatchOfUserBalances() failed")
		return err
	}

	// store logs / histories
	err = e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&log)
	if err != nil {
		e.log.Fields(logger.Fields{"model": log}).Error(err, "[entity.HandleIncomingDeposit] repo.OffchainTipBotActivityLogs.CreateActivityLog() failed")
	}
	histories := []model.OffchainTipBotTransferHistory{
		{
			ReceiverID: userID,
			LogID:      log.ID.String(),
			Status:     consts.OffchainTipBotTrasferStatusSuccess,
			Amount:     amount,
			Token:      offchainToken.TokenSymbol,
			Action:     transferType,
			TxHash:     req.TxHash,
		},
	}
	_, err = e.repo.OffchainTipBotTransferHistories.CreateTransferHistories(histories)
	if err != nil {
		e.log.Fields(logger.Fields{"histories": histories}).Error(err, "[entity.HandleIncomingDeposit] repo.OffchainTipBotTransferHistories.CreateTransferHistories() failed")
	}

	// discord notification
	err = e.notifyDepositTx(amount, userID, chain.ExplorerURL, req.TxHash, *offchainToken, priceInUSD)
	if err != nil {
		e.log.Fields(logger.Fields{"amount": amount, "userID": userID, "token": *offchainToken}).Error(err, "[entity.HandleIncomingDeposit] e.notifyDepositTx() failed")
	}

	// sweep tokens
	err = e.sweepTokens(req, assignedContract.ContractID.String(), token, chain.IsEVM)
	if err != nil {
		e.log.Fields(logger.Fields{"contract": req.ToAddress, "userID": userID, "token": req.TokenSymbol}).Error(err, "[entity.HandleIncomingDeposit] e.sweepTokens() failed")
	}
	e.sendTipBotLogs(log, offchainToken.TokenSymbol, req.FromAddress)
	return nil
}

func (e *Entity) sweepTokens(req request.TipBotDepositRequest, contractID string, token model.Token, isEVM bool) error {
	// EVM
	if isEVM {
		tx, err := e.abi.SweepTokens(req.ToAddress, int64(req.ChainID), token)
		if err != nil {
			e.log.Fields(logger.Fields{"address": req.ToAddress, "chainID": req.ChainID, "token": token}).Error(err, "[entity.sweepTokens] e.abi.SweepTokens() failed")
			return err
		}
		e.repo.OffchainTipBotContract.UpdateSweepTime(contractID, time.Now())
		e.log.Infof("[entity.sweepTokens] sucessfully sweep EVM tokens: %s", tx.Hash().Hex())
		return nil
	}

	// 2. Solana
	contract, err := e.repo.OffchainTipBotContract.GetByID(contractID)
	if err != nil {
		e.log.Fields(logger.Fields{"contractID": contractID}).Error(err, "[entity.sweepTokens] repo.OffchainTipBotContract.GetByID() failed")
		return err
	}
	if contract.PrivateKey == "" {
		e.log.Fields(logger.Fields{"contract": req.ToAddress}).Infof("[entity.sweepTokens] Solana contract has no PK")
		return nil
	}
	pk, err := util.DecodeCFB(e.cfg.SolanaPKSecretKey, contract.PrivateKey)
	if err != nil {
		e.log.Fields(logger.Fields{"contract": req.ToAddress}).Infof("[entity.sweepTokens] util.DecodeCFB() failed")
		return err
	}
	txHash, _, err := e.solana.Transfer(pk, e.solana.GetCentralizedWalletAddress(), 0, true)
	if err != nil {
		e.log.Fields(logger.Fields{"address": req.ToAddress, "chainID": req.ChainID}).Error(err, "[entity.sweepTokens] e.solana.Transfer() failed")
		return err
	}
	e.repo.OffchainTipBotContract.UpdateSweepTime(contractID, time.Now())
	e.log.Infof("[entity.sweepTokens] sucessfully sweep Solana: %s", txHash)
	return nil
}

func (e *Entity) notifyDepositTx(amount float64, userID, explorerUrl, signature string, token model.OffchainTipBotToken, priceInUSD float64) error {
	dmChannel, err := e.discord.UserChannelCreate(userID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": userID}).Error(err, "[entity.notifyDepositTx] discord.UserChannelCreate() failed")
		return err
	}
	userBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(userID, token.ID)
	if err != nil {
		e.log.Fields(logger.Fields{"tokenID": token.ID, "userID": userID}).Error(err, "[entity.notifyDepositTx] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return err
	}

	amountInUSD := amount * priceInUSD
	balanceInUSD := userBal.Amount * priceInUSD
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:approve:1013775501757780098> Confirm Transaction"),
		Description: fmt.Sprintf("Your **%s** (%s) deposit has been confirmed.", token.TokenName, token.TokenSymbol),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Deposited",
				Inline: true,
				Value:  fmt.Sprintf("%.6f %s\n`$%.2f`", amount, token.TokenSymbol, amountInUSD),
			},
			{
				Name:   "Balance",
				Inline: true,
				Value:  fmt.Sprintf("%.6f %s\n`$%.2f`", userBal.Amount, token.TokenSymbol, balanceInUSD),
			},
			{
				Name:  "Transaction ID",
				Value: fmt.Sprintf("[`%s`](%s/tx/%s)", signature, explorerUrl, signature),
			},
		},
	}
	_, err = e.discord.ChannelMessageSendEmbed(dmChannel.ID, embed)
	if err != nil {
		e.log.Fields(logger.Fields{"dmChannel.ID": dmChannel.ID}).Error(err, "[entity.notifyDepositTx] discord.ChannelMessageSendEmbed() failed")
	}
	return err
}

func (e *Entity) GetOneDepositTx(chainID string, txHash string) (*model.OffchainTipBotDepositLog, error) {
	return e.repo.OffchainTipBotDepositLog.GetOne(chainID, txHash)
}

func (e *Entity) UpsertBatchOfUserBalances(action, tokenSymbol string, list []model.OffchainTipBotUserBalance) error {
	err := e.repo.OffchainTipBotUserBalances.UpsertBatch(list)
	if err != nil {
		e.log.Fields(logger.Fields{"batch": len(list)}).Error(err, "[entity.UpsertBatchOfUserBalances] repo.OffchainTipBotUserBalances.UpsertBatch() failed")
		return err
	}
	snapshots := make([]model.OffchainTipBotUserBalanceSnapshot, 0, len(list))
	for _, ub := range list {
		snapshots = append(snapshots, model.OffchainTipBotUserBalanceSnapshot{
			UserID:        ub.UserID,
			TokenID:       ub.TokenID,
			Action:        action,
			TokenSymbol:   tokenSymbol,
			ChangedAmount: ub.ChangedAmount,
			Amount:        ub.Amount,
		})
	}
	err = e.repo.OffchainTipBotUserBalanceSnapshot.CreateBatch(snapshots)
	if err != nil {
		e.log.Fields(logger.Fields{"batch": len(list)}).Error(err, "[entity.UpsertBatchOfUserBalances] repo.OffchainTipBotUserBalanceSnapshot.CreateBatch() failed")
	}
	return nil
}

func (e *Entity) sendTipBotLogs(log model.OffchainTipBotActivityLog, token, depositSource string) error {
	if log.GuildID == nil || *log.GuildID == "" || *log.GuildID == "DM" {
		return nil
	}
	guild, err := e.repo.DiscordGuilds.GetByID(*log.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": *log.GuildID}).Error(err, "[entity.sendTipBotLogs] repo.DiscordGuilds.GetByID() failed")
		return err
	}
	if guild.LogChannel == "" {
		return nil
	}

	token = strings.ToUpper(token)
	senderID := *log.UserID
	amount := log.Amount

	var description string
	switch *log.Action {
	case "tip":
		var recipients []string
		for _, r := range log.Receiver {
			recipients = append(recipients, fmt.Sprintf("<@%s>", r))
		}
		recipientsStr := strings.Join(recipients, ",")
		description = fmt.Sprintf("<@%s> has sent %s **%g %s** each at <#%s>", senderID, recipientsStr, amount, token, *log.ChannelID)
		if len(description) > 2000 {
			description = fmt.Sprintf("<@%s> has sent %d people **%g %s** each at <#%s>", senderID, len(recipients), amount, token, *log.ChannelID)
		}
	case "withdraw":
		description = fmt.Sprintf("<@%s> has made a withdrawal of **%g %s** to address `%s` at <#%s>", senderID, log.Amount, token, log.Receiver[0], *log.ChannelID)
	case "deposit":
		description = fmt.Sprintf("<@%s> has just deposited **%g %s** from address `%s`", senderID, log.Amount, token, depositSource)
	case "airdrop":
		description = fmt.Sprintf("<@%s> has just airdropped **%g %s** at <#%s>", senderID, log.Amount, token, *log.ChannelID)
	}
	err = e.svc.Discord.SendGuildActivityLogs(guild.LogChannel, senderID, strings.ToUpper(*log.Action), description)
	if err != nil {
		e.log.Fields(logger.Fields{"activity": log}).Error(err, "[entity.sendTipBotLogs] svc.Discord.SendGuildActivityLogs() failed")
		return err
	}
	return nil
}
