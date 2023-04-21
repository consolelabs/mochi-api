package entities

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

// TODO@anhnh: define error codes
func (e *Entity) TransferToken(req request.OffchainTransferRequest) (*response.OffchainTipBotTransferToken, error) {
	e.log.Fields(logger.Fields{"req": req}).Info("receive new transfer request")
	// get senderProfile, recipientProfiles by discordID
	transferReq := request.MochiPayTransferRequest{}
	senderProfile, err := e.svc.MochiProfile.GetByDiscordID(req.Sender)
	if err != nil {
		return nil, errors.New(consts.OffchainTipBotFailReasonGetProfileFailed)
	}

	transferReq.From = &request.Wallet{
		ProfileGlobalId: senderProfile.ID,
	}

	for _, v := range req.Recipients {
		profile, err := e.svc.MochiProfile.GetByDiscordID(v)
		if err != nil {
			return nil, errors.New(consts.OffchainTipBotFailReasonGetProfileFailed)
		}
		transferReq.Tos = append(transferReq.Tos, &request.Wallet{
			ProfileGlobalId: profile.ID,
		})
	}

	// validate amount
	if !req.All && req.Amount == 0 {
		return nil, errors.New(consts.OffchainTipBotFailReasonInvalidAmount)
	}

	// validate token
	token, err := e.svc.MochiPay.GetToken(req.Token, req.ChainID)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entity.TransferToken] svc.MochiPay.GetToken() failed")
		return nil, err
	}
	transferReq.TokenId = token.Id
	transferReq.Note = req.Message

	// validate balance
	senderBalance, err := e.svc.MochiPay.GetBalance(senderProfile.ID, req.Token, req.ChainID)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[entity.TransferToken] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return nil, err
	}

	bal, err := util.StringToBigInt(senderBalance.Data[0].Amount)
	if err != nil {
		return nil, errors.New(consts.OffchainTipBotFailReasonInvalidAmount)
	}

	if bal.Cmp(big.NewInt(0)) != 1 {
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// if transfer all -> amount = sender balance
	fBal := util.BigIntToFloat(bal, int(senderBalance.Data[0].Token.Decimal))
	if req.All {
		req.Amount = fBal
	}

	// calculate transferred amount for each recipient
	var amountEach float64
	if req.Each && !req.All {
		amountEach = req.Amount
	} else {
		amountEach = req.Amount / float64(len(req.Recipients))
	}
	if fBal < amountEach {
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	transferReq.Amount = make([]string, len(req.Recipients))
	for i := range transferReq.Amount {
		transferReq.Amount[i] = strconv.FormatFloat(amountEach, 'f', int(token.Decimal), 64)
	}

	err = e.svc.MochiPay.Transfer(transferReq)
	if err != nil {
		return nil, errors.New(consts.OffchainTipBotFailReasonMochiPayTransferFailed)
	}

	// notify tip to channel
	e.sendLogNotify(req, int(token.Decimal))

	// tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{req.Token}, "usd")
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"token": token.CoinGeckoId}).Error(err, "[entity.TransferToken] svc.CoinGecko.GetCoinPrice() failed")
	// }

	// notify tip to other platform: twitter, telegram, ...
	// e.NotifyTipFromPlatforms(req, amountEach, tokenPrice[token.CoinGeckoId])

	// e.sendTipBotLogs(req, token.Symbol, "")

	return &response.OffchainTipBotTransferToken{
		AmountEach:  amountEach,
		TotalAmount: req.Amount,
		// Token:       token,
	}, nil
}

func (e *Entity) sendLogNotify(req request.OffchainTransferRequest, decimal int) {
	if req.TransferType != consts.OffchainTipBotTransferTypeTip && req.TransferType != consts.OffchainTipBotTransferTypeAirdrop {
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
			descriptionFormat := ""
			name := ""
			switch req.TransferType {
			case "tip":
				name = "Someone sent out money"
				descriptionFormat = "<@%s> has just sent %s **%s %s** at <#%s>"
				if req.Each {
					descriptionFormat = "<@%s> has just sent %s **%s %s** each at <#%s>"
				}
			case "airdrop":
				name = "Someone dropped money"
				descriptionFormat = "<@%s> has just airdropped %s **%s %s** at <#%s>"
			}
			description := fmt.Sprintf(descriptionFormat, req.Sender, recipientsStr, req.AmountString, strings.ToUpper(req.Token), req.ChannelID)
			if req.Message != "" {
				description += fmt.Sprintf("\n<a:_:1095990167350816869> **%s**", req.Message)
			}
			author := &discordgo.MessageEmbedAuthor{
				Name:    name,
				IconURL: "https://cdn.discordapp.com/emojis/1093923019988148354.gif?size=240&quality=lossless",
			}

			err := e.svc.Discord.SendTipActivityLogs(configNotifyChannel.ChannelID, req.Sender, author, description, req.Image)
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

// TODO: remove once mochi-pay migration is done
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

func (e *Entity) notifyDepositTx(bal float64, userID, explorerUrl, signature string, token model.OffchainTipBotToken, priceInUSD float64) error {
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

	amountInUSD := bal * priceInUSD
	balanceInUSD := userBal.Amount * priceInUSD
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:approve:1013775501757780098> Confirm Transaction"),
		Description: fmt.Sprintf("Your **%s** (%s) deposit has been confirmed.", token.TokenName, token.TokenSymbol),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Deposited",
				Inline: true,
				Value:  fmt.Sprintf("%.6f %s\n`$%.2f`", bal, token.TokenSymbol, amountInUSD),
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

func (e *Entity) sendTipBotLogs(req request.OffchainTransferRequest, token, depositSource string) error {
	if req.GuildID == "" || req.GuildID == "DM" {
		return nil
	}
	guild, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[entity.sendTipBotLogs] repo.DiscordGuilds.GetByID() failed")
		return err
	}
	if guild.LogChannel == "" {
		return nil
	}

	token = strings.ToUpper(token)
	senderID := req.Sender
	amount := req.Amount

	var description string
	switch req.TransferType {
	case "tip":
		var recipients []string
		for _, r := range req.Recipients {
			recipients = append(recipients, fmt.Sprintf("<@%s>", r))
		}
		recipientsStr := strings.Join(recipients, ",")
		description = fmt.Sprintf("<@%s> has sent %s **%g %s** each at <#%s>", senderID, recipientsStr, amount, token, req.ChannelID)
		if len(description) > 2000 {
			description = fmt.Sprintf("<@%s> has sent %d people **%g %s** each at <#%s>", senderID, len(recipients), amount, token, req.ChannelID)
		}
	// case "withdraw":
	// 	description = fmt.Sprintf("<@%s> has made a withdrawal of **%g %s** to address `%s` at <#%s>", senderID, log.Amount, token, log.Receiver[0], *log.ChannelID)
	// case "deposit":
	// 	description = fmt.Sprintf("<@%s> has just deposited **%g %s** from address `%s`", senderID, log.Amount, token, depositSource)
	case "airdrop":
		description = fmt.Sprintf("<@%s> has just airdropped **%g %s** at <#%s>", senderID, req.Amount, token, req.ChannelID)
	}
	err = e.svc.Discord.SendGuildActivityLogs(guild.LogChannel, senderID, strings.ToUpper(req.TransferType), description)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.sendTipBotLogs] svc.Discord.SendGuildActivityLogs() failed")
		return err
	}
	return nil
}
