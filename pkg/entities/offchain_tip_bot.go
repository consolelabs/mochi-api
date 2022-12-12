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
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) TransferToken(req request.OffchainTransferRequest) ([]response.OffchainTipBotTransferToken, error) {
	// check supported tokens
	amountEachRecipient := req.Amount / float64(len(req.Recipients))
	supportedToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(req.Token))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
				UserID:          req.Sender,
				GuildID:         req.GuildID,
				ChannelID:       req.ChannelID,
				Action:          &req.TransferType,
				Receiver:        req.Recipients,
				NumberReceivers: len(req.Recipients),
				Duration:        &req.Duration,
				Amount:          amountEachRecipient,
				Status:          consts.OffchainTipBotTrasferStatusFail,
				FullCommand:     &req.FullCommand,
				FailReason:      consts.OffchainTipBotFailReasonTokenNotSupported,
				Image:           req.Image,
				Message:         req.Message,
			})
			return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonTokenNotSupported)
		}
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[repo.OffchainTipBotTokens.GetBySymbol] - failed to get check supported token")
		return nil, err
	}

	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{supportedToken.CoinGeckoID}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": supportedToken.CoinGeckoID}).Error(err, "[svc.CoinGecko.GetCoinPrice] - failed to get coin price from Coingecko")
		return nil, err
	}

	// check user bals, both not have record user_bals + amount in record user_bals = 0 -> return not enough bals
	modelNotEnoughBalance := &model.OffchainTipBotActivityLog{
		UserID:          req.Sender,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        req.Recipients,
		NumberReceivers: len(req.Recipients),
		TokenID:         supportedToken.ID.String(),
		Duration:        &req.Duration,
		Amount:          amountEachRecipient,
		Status:          consts.OffchainTipBotTrasferStatusFail,
		FullCommand:     &req.FullCommand,
		FailReason:      consts.OffchainTipBotFailReasonNotEnoughBalance,
		Image:           req.Image,
		Message:         req.Message,
	}
	userBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(req.Sender, supportedToken.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
			return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
		}
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID] - failed to get user balance")
		return nil, err
	}

	if req.All {
		if userBal.Amount == 0 {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
			return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
		}
		req.Amount = userBal.Amount
		amountEachRecipient = req.Amount / float64(len(req.Recipients))
	}

	if float64(userBal.Amount) < req.Amount {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
		return []response.OffchainTipBotTransferToken{}, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// case record offchain_tip_bot_user_balances if not exist yet, CreateIfNotExists here to ensure that
	// TODO(trkhoi): gorm not support upsert batch model, temp do like this will find another way later
	for _, recipient := range req.Recipients {
		e.repo.OffchainTipBotUserBalances.CreateIfNotExists(&model.OffchainTipBotUserBalance{
			UserID:  recipient,
			TokenID: supportedToken.ID,
		})
	}

	// create activity log
	al, err := e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
		UserID:          req.Sender,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        req.Recipients,
		TokenID:         supportedToken.ID.String(),
		NumberReceivers: len(req.Recipients),
		Duration:        &req.Duration,
		Amount:          amountEachRecipient,
		Status:          consts.OffchainTipBotTrasferStatusSuccess,
		FullCommand:     &req.FullCommand,
		FailReason:      "",
		Image:           req.Image,
		Message:         req.Message,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotActivityLogs.CreateActivityLog] - failed to create activity log")
		return nil, err
	}

	// create transfer histories for each transfer
	listTransferHistories := make([]model.OffchainTipBotTransferHistory, 0)
	for _, recipient := range req.Recipients {
		listTransferHistories = append(listTransferHistories, model.OffchainTipBotTransferHistory{
			SenderID:   req.Sender,
			ReceiverID: recipient,
			GuildID:    req.GuildID,
			LogID:      al.ID.String(),
			Status:     consts.OffchainTipBotTrasferStatusSuccess,
			Amount:     amountEachRecipient,
			Token:      supportedToken.TokenSymbol,
			Action:     req.TransferType,
		})
	}
	transferHistories, err := e.repo.OffchainTipBotTransferHistories.CreateTransferHistories(listTransferHistories)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotTransferHistories.CreateTransferHistories] - failed to create transfer histories")
		return nil, err
	}

	// update recipients balances
	err = e.repo.OffchainTipBotUserBalances.UpdateListUserBalances(req.Recipients, supportedToken.ID, amountEachRecipient)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotUserBalances.UpdateListUserBalances] - failed to update recipients balances")
		return nil, err
	}

	// update sender balanace
	err = e.repo.OffchainTipBotUserBalances.UpdateUserBalance(&model.OffchainTipBotUserBalance{UserID: req.Sender, TokenID: supportedToken.ID, Amount: userBal.Amount - req.Amount})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotUserBalances.UpdateUserBalance] - failed to update sender balance")
		return nil, err
	}

	// notify tip to channel
	e.sendLogNotify(req, amountEachRecipient)

	// notify tip to other platform: twitter, telegram, ...
	e.NotifyTipFromPlatforms(req, amountEachRecipient, tokenPrice[supportedToken.CoinGeckoID])

	return e.MappingTransferTokenResponse(req.Token, amountEachRecipient, tokenPrice[supportedToken.CoinGeckoID], transferHistories), nil
}

func (e *Entity) sendLogNotify(req request.OffchainTransferRequest, amountEachRecipient float64) {
	if req.TransferType == consts.OffchainTipBotTransferTypeTip {
		// Do not return error here, just log it
		configNotifyChannels, err := e.repo.OffchainTipBotConfigNotify.GetByGuildID(req.GuildID)
		if err != nil {
			e.log.Fields(logger.Fields{"guild_id": req.GuildID}).Error(err, "[repo.OffchainTipBotConfigNotify.GetByGuildID] - failed to get config notify channels")
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
					e.log.Fields(logger.Fields{"channel_id": configNotifyChannel.ChannelID}).Error(err, "[discord.ChannelMessageSendEmbed] - failed to send message to channel")
				}
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
				e.log.Fields(logger.Fields{"req": req, "recipient": recipient}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalances] - failed to get user balances")
				return
			}

			dmChannel, err := e.discord.UserChannelCreate(recipient)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req, "recipient": recipient}).Error(err, "[discord.UserChannelCreate] - failed to create DM channel")
				return
			}

			balsEmbed := make([]*discordgo.MessageEmbedField, 0)

			for _, bal := range recipientBals {
				balsEmbed = append(balsEmbed, &discordgo.MessageEmbedField{
					Name:   bal.Name,
					Inline: true,
					Value:  fmt.Sprintf("%s %.2f %s `$%.2f`", util.GetEmojiToken(bal.Symbol), bal.Balances, bal.Symbol, bal.BalancesInUSD),
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
				e.log.Fields(logger.Fields{"req": req, "recipient": recipient}).Error(err, "[discord.ChannelMessageSendEmbeds] - failed to send DM to recipient")
				return
			}
		}
	}
}

func (e *Entity) MappingTransferTokenResponse(tokenSymbol string, amount float64, price float64, transferHistories []model.OffchainTipBotTransferHistory) (res []response.OffchainTipBotTransferToken) {
	for _, transferHistory := range transferHistories {
		res = append(res, response.OffchainTipBotTransferToken{
			SenderID:    transferHistory.SenderID,
			RecipientID: transferHistory.ReceiverID,
			Amount:      amount,
			Symbol:      tokenSymbol,
			AmountInUSD: amount * price,
		})
	}
	return res
}

func (e *Entity) OffchainTipBotWithdraw(req request.OffchainWithdrawRequest) (*response.OffchainTipBotWithdraw, error) {
	// check supported tokens
	offchainToken, err := e.repo.OffchainTipBotTokens.GetBySymbol(req.Token)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
				UserID:          req.Recipient,
				GuildID:         req.GuildID,
				ChannelID:       req.ChannelID,
				Action:          &req.TransferType,
				Receiver:        []string{req.Recipient},
				NumberReceivers: 1,
				Duration:        &req.Duration,
				Amount:          req.Amount,
				Status:          consts.OffchainTipBotTrasferStatusFail,
				FullCommand:     &req.FullCommand,
				FailReason:      consts.OffchainTipBotFailReasonTokenNotSupported,
			})
			return nil, errors.New(consts.OffchainTipBotFailReasonTokenNotSupported)
		}
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[repo.OffchainTipBotTokens.GetBySymbol] - failed to get check supported token")
		return nil, err
	}

	// check recipient balance
	modelNotEnoughBalance := &model.OffchainTipBotActivityLog{
		UserID:          req.Recipient,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        []string{req.Recipient},
		NumberReceivers: 1,
		TokenID:         offchainToken.ID.String(),
		Duration:        &req.Duration,
		Amount:          req.Amount,
		Status:          consts.OffchainTipBotTrasferStatusFail,
		FullCommand:     &req.FullCommand,
		FailReason:      consts.OffchainTipBotFailReasonNotEnoughBalance,
		ServiceFee:      offchainToken.ServiceFee,
		FeeAmount:       offchainToken.ServiceFee * req.Amount,
	}
	recipientBal, err := e.repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID(req.Recipient, offchainToken.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
			return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
		}
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Recipient}).Error(err, "[repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID] - failed to get user balance")
		return nil, err
	}

	if float64(recipientBal.Amount) < req.Amount+offchainToken.ServiceFee*req.Amount {
		e.repo.OffchainTipBotActivityLogs.CreateActivityLog(modelNotEnoughBalance)
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// temp get from old tokens because not have flow from coingecko to offchain_tip_bot_tokens yet
	token, err := e.repo.Token.GetBySymbol(strings.ToLower(req.Token), true)
	if err != nil {
		return nil, err
	}

	// execute tx
	signedTx, transferredAmount, err := e.transferOffchain(req.Amount,
		accounts.Account{Address: common.HexToAddress(req.RecipientAddress)},
		req.Amount,
		token, -1, req.All)
	if err != nil {
		err = fmt.Errorf("error transfer: %v", err)
		return nil, err
	}

	// execute tx success -> create offchain_tip_bot_activity_logs + offchain_tip_bot_transfer_histories and update offchain_tip_bot_user_balances
	al, err := e.repo.OffchainTipBotActivityLogs.CreateActivityLog(&model.OffchainTipBotActivityLog{
		UserID:          req.Recipient,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Action:          &req.TransferType,
		Receiver:        []string{req.RecipientAddress},
		TokenID:         offchainToken.ID.String(),
		NumberReceivers: 1,
		Duration:        &req.Duration,
		Amount:          transferredAmount,
		Status:          consts.OffchainTipBotTrasferStatusSuccess,
		FullCommand:     &req.FullCommand,
		FailReason:      "",
		ServiceFee:      offchainToken.ServiceFee,
		FeeAmount:       offchainToken.ServiceFee * req.Amount,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotActivityLogs.CreateActivityLog] - failed to create activity log")
		return nil, err
	}

	_, err = e.repo.OffchainTipBotTransferHistories.CreateTransferHistories([]model.OffchainTipBotTransferHistory{{
		SenderID:   req.Recipient,
		ReceiverID: req.Recipient,
		GuildID:    req.GuildID,
		LogID:      al.ID.String(),
		Status:     consts.OffchainTipBotTrasferStatusSuccess,
		Amount:     transferredAmount,
		Token:      offchainToken.TokenSymbol,
		Action:     req.TransferType,
		ServiceFee: offchainToken.ServiceFee,
		FeeAmount:  offchainToken.ServiceFee * req.Amount,
	}})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotTransferHistories.CreateTransferHistories] - failed to create transfer histories")
		return nil, err
	}

	err = e.repo.OffchainTipBotUserBalances.UpdateUserBalance(&model.OffchainTipBotUserBalance{UserID: req.Recipient, TokenID: offchainToken.ID, Amount: recipientBal.Amount - transferredAmount - offchainToken.ServiceFee*req.Amount})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.OffchainTipBotUserBalances.UpdateUserBalance] - failed to update sender balance")
		return nil, err
	}

	withdrawalAmount := util.WeiToEther(signedTx.Value())
	transactionFee, _ := util.WeiToEther(new(big.Int).Sub(signedTx.Cost(), signedTx.Value())).Float64()

	return &response.OffchainTipBotWithdraw{
		UserDiscordID:  req.Recipient,
		ToAddress:      req.RecipientAddress,
		Amount:         transferredAmount,
		Symbol:         req.Token,
		TxHash:         signedTx.Hash().Hex(),
		TxUrl:          fmt.Sprintf("%s/%s", token.Chain.TxBaseURL, signedTx.Hash().Hex()),
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
