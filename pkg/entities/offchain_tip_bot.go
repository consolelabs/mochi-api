package entities

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	query "github.com/defipod/mochi/pkg/repo/guild_config_log_channel"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochipay"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) TransferToken(req request.OffchainTransferRequest) (*response.OffchainTipBotTransferToken, error) {
	e.log.Fields(logger.Fields{"req": req}).Info("receive new transfer request")
	// get senderProfile, recipientProfiles by discordID
	transferReq := request.MochiPayTransferRequest{}
	transferReq.From = &request.Wallet{
		ProfileGlobalId: req.Sender,
	}

	for _, r := range req.Recipients {
		transferReq.Tos = append(transferReq.Tos, &request.Wallet{
			ProfileGlobalId: r,
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
	senderBalance, err := e.svc.MochiPay.GetBalance(req.Sender, req.Token, req.ChainID)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[entity.TransferToken] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return nil, err
	}

	if len(senderBalance.Data) == 0 {
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	bal, err := util.StringToBigInt(senderBalance.Data[0].Amount)
	if err != nil {
		return nil, errors.New(consts.OffchainTipBotFailReasonInvalidAmount)
	}

	if bal.Cmp(big.NewInt(0)) != 1 {
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	// calculate transferred amount for each recipient
	var amountEach float64
	if req.Each && !req.All {
		amountEach = req.Amount
	} else {
		amountEach = req.Amount / float64(len(req.Recipients))
	}
	amountEachStr := strconv.FormatFloat(amountEach, 'f', -1, 64)

	transferReq.Amount = make([]string, len(req.Recipients))
	for i := range transferReq.Amount {
		transferReq.Amount[i] = amountEachStr
	}

	//validate tip range
	tipRangeConfig, err := e.repo.GuildConfigTipRange.GetByGuildID(req.GuildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID}).Error(err, "[entity.TransferToken] repo.GuildConfigTipRange.GetByGuildID() failed")
		return nil, errors.New("get price coingecko failed")
	}

	if tipRangeConfig != nil {
		//TODO: move this get tokenPrice block code out if using for others validate
		tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{token.CoinGeckoId}, "usd")
		if err != nil {
			e.log.Fields(logger.Fields{"token": token.CoinGeckoId}).Error(err, "[entity.TransferToken] svc.CoinGecko.GetCoinPrice() failed")
		}

		//only validate if have tokenPrice && tipRangeConfig
		if tipRangeConfig.Min != nil && tokenPrice[token.CoinGeckoId] > 0 {
			if tipRangeConfig.Min != nil && tokenPrice[token.CoinGeckoId]*amountEach < *tipRangeConfig.Min {
				return nil, errors.New("tip amount < min tip range")
			}
			if tipRangeConfig.Max != nil && tokenPrice[token.CoinGeckoId]*amountEach > *tipRangeConfig.Max {
				return nil, errors.New("tip amount > max tip range")
			}
		}
	}

	tx, err := e.svc.MochiPay.Transfer(transferReq)
	if err != nil {
		return nil, errors.New(consts.OffchainTipBotFailReasonMochiPayTransferFailed)
	}

	// notify tip to channel
	e.sendLogNotify(req, int(token.Decimal), amountEachStr)

	return &response.OffchainTipBotTransferToken{
		Id:          tx.Data.Id,
		AmountEach:  amountEach,
		TotalAmount: req.Amount,
		TxId:        tx.Data.TxId,
	}, nil
}

func (e *Entity) sendLogNotify(req request.OffchainTransferRequest, decimal int, amountEachStr string) {
	if req.TransferType != consts.OffchainTipBotTransferTypeTip && req.TransferType != consts.OffchainTipBotTransferTypeAirdrop {
		return
	}
	// Do not return error here, just log it
	configNotifyChannels, err := e.repo.GuildConfigLogChannel.Get(query.Query{LogType: "tip", GuildId: req.GuildID})
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID}).Error(err, "[entity.sendLogNotify] repo.OffchainTipBotConfigNotify.GetByGuildID() failed")
		return
	}

	for _, configNotifyChannel := range configNotifyChannels {
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
		description := fmt.Sprintf(descriptionFormat, req.Sender, recipientsStr, amountEachStr, strings.ToUpper(req.Token), req.ChannelID)
		if req.Message != "" {
			description += fmt.Sprintf("\n<a:_:1095990167350816869> **%s**", req.Message)
		}
		author := &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: "https://cdn.discordapp.com/emojis/1093923019988148354.gif?size=240&quality=lossless",
		}

		err := e.svc.Discord.SendTipActivityLogs(configNotifyChannel.ChannelId, req.Sender, author, description, req.Image)
		if err != nil {
			e.log.Fields(logger.Fields{"channel_id": configNotifyChannel.ChannelId}).Error(err, "[entity.sendLogNotify] discord.ChannelMessageSendEmbed() failed")
		}
	}
}

func (e *Entity) TransferTokenV2(req request.TransferV2Request) (*response.TransferTokenV2Data, error) {
	e.log.Fields(logger.Fields{"component": "TransferV2", "req": req}).Info("receive new transfer request")
	// get senderProfile, recipientProfiles by discordID
	transferReq := mochipay.TransferV2Request{
		From: &mochipay.Wallet{
			ProfileGlobalId: req.Sender,
		},
		Platform: req.Platform,
		Metadata: req.Metadata,
		Action:   req.TransferType,
	}

	for _, r := range req.Recipients {
		transferReq.Tos = append(transferReq.Tos, &mochipay.Wallet{
			ProfileGlobalId: r,
		})
	}

	// validate token
	token, err := e.svc.MochiPay.GetToken(req.Token, req.ChainID)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entity.TransferTokenV2] svc.MochiPay.GetToken() failed")
		return nil, err
	}
	transferReq.TokenId = token.Id

	// validate amount
	amount := util.FloatToBigInt(req.Amount, token.Decimal)
	if !req.All && amount.Cmp(big.NewInt(0)) != 1 {
		return nil, errors.New(consts.OffchainTipBotFailReasonInvalidAmount)
	}

	// validate balance
	senderBalance, err := e.svc.MochiPay.GetBalance(req.Sender, req.Token, req.ChainID)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[entity.TransferTokenV2] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return nil, err
	}

	if len(senderBalance.Data) == 0 {
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	bal, err := util.StringToBigInt(senderBalance.Data[0].Amount)
	if err != nil {
		return nil, errors.New(consts.OffchainTipBotFailReasonInvalidAmount)
	}

	if bal.Cmp(amount) != 1 {
		return nil, errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	amountEach := new(big.Int)
	if req.Each && !req.All {
		amountEach = amount
	} else {
		amountEach.Quo(amount, big.NewInt(int64(len(req.Recipients))))
	}
	amountEachF := util.BigIntToFloat(amountEach, int(token.Decimal))

	transferReq.Amount = make([]string, len(req.Recipients))
	for i := range transferReq.Amount {
		amt, _ := new(big.Float).SetInt(amountEach).Float64()
		transferReq.Amount[i] = fmt.Sprintf("%v", strconv.FormatFloat(amt, 'f', -1, 64))
	}

	//validate tip range
	err = e.validateTipRange(req.GuildID, token.CoinGeckoId, amountEachF)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TransferTokenV2] validateTipRange() failed")
		return nil, err
	}

	res, err := e.svc.MochiPay.TransferV2(transferReq)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TransferTokenV2] svc.MochiPay.TransferV2() failed")
		return nil, errors.New(consts.OffchainTipBotFailReasonMochiPayTransferFailed)
	}

	// notify tip to channel
	// e.sendLogNotify(req, int(token.Decimal), amountEachStr)

	if len(res.Data) == 0 {
		e.log.Error(err, "[entity.TransferTokenV2] no transfer response")
		return nil, errors.New(consts.OffchainTipBotFailReasonMochiPayTransferFailed)
	}

	internalId := res.Data[0].InternalId
	externalId := res.Data[0].ExternalId
	id := res.Data[0].ID

	return &response.TransferTokenV2Data{
		Id:          id,
		AmountEach:  amountEachF,
		TotalAmount: req.Amount,
		TxId:        internalId,
		ExternalId:  externalId,
	}, nil
}

func (e *Entity) validateTipRange(guildID, coinGeckoID string, amount float64) error {
	//validate tip range
	tipRangeConfig, err := e.repo.GuildConfigTipRange.GetByGuildID(guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[entity.TransferToken] repo.GuildConfigTipRange.GetByGuildID() failed")
		return errors.New("failed to get tip range config")
	}

	// no tip range restriction -> ignore
	if tipRangeConfig == nil {
		return nil
	}

	tokenPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{coinGeckoID}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": coinGeckoID}).Error(err, "[entity.TransferToken] svc.CoinGecko.GetCoinPrice() failed")
		return err
	}

	// if no token price data -> ignore
	if tokenPrice[coinGeckoID] == 0 {
		return nil
	}

	usdAmount := tokenPrice[coinGeckoID] * amount
	if tipRangeConfig.Min != nil && usdAmount < *tipRangeConfig.Min {
		return errors.New("tip amount < min tip range")
	}
	if tipRangeConfig.Max != nil && usdAmount > *tipRangeConfig.Max {
		return errors.New("tip amount > max tip range")
	}

	return nil
}
