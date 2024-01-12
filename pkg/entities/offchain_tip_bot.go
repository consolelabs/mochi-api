package entities

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
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

	tx, transferErr := e.svc.MochiPay.Transfer(transferReq)
	if transferErr != nil {
		e.log.Fields(logger.Fields{"token": token.CoinGeckoId}).Error(transferErr, "[entity.TransferToken] svc.MochiPay.Transfer() failed")
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
	logger := e.log.Fields(logger.Fields{"component": "entity.TransferV2", "req": req})
	logger.Info("receive new transfer request")
	template := parseTemplate(req)

	// validate token
	token, err := e.svc.MochiPay.GetToken(req.Token, req.ChainID)
	if err != nil {
		logger.Error(err, "[entity.TransferTokenV2] svc.MochiPay.GetToken() failed")
		return nil, err
	}

	// convert total transfer amount
	totalAmount := util.FloatToBigInt(req.Amount, token.Decimal)

	// validate balance
	if err := e.validateTransferBalance(totalAmount, req); err != nil {
		logger.Error(err, "[entity.TransferTokenV2] svc.MochiPay.GetToken() failed")
		return nil, err
	}

	// calculate respective amount for each recipient
	amountPerRecipient := new(big.Int)
	if req.Each && !req.All {
		amountPerRecipient = totalAmount
	} else {
		amountPerRecipient.Quo(totalAmount, big.NewInt(int64(len(req.Recipients))))
	}
	fAmountPerTx := util.BigIntToFloat(amountPerRecipient, int(token.Decimal))

	//validate tip range
	err = e.validateTipRange(req.GuildID, fAmountPerTx, token)
	if err != nil {
		logger.Error(err, "validateTipRange() failed")
		return nil, err
	}

	// compose transfer request
	transferReq := e.composeTransferRequest(req, token, amountPerRecipient, template)
	// send request to Mochi Pay
	res, err := e.svc.MochiPay.TransferV2(transferReq)
	if err != nil {
		logger.Error(err, "svc.MochiPay.TransferV2() failed")
		return nil, errors.New(consts.OffchainTipBotFailReasonMochiPayTransferFailed)
	}

	if len(res.Data) == 0 {
		e.log.Error(err, "no transfer response")
		return nil, errors.New(consts.OffchainTipBotFailReasonMochiPayTransferFailed)
	}

	internalId := res.Data[0].InternalId
	externalId := res.Data[0].ExternalId
	id := res.Data[0].ID

	return &response.TransferTokenV2Data{
		Id:          id,
		AmountEach:  fAmountPerTx,
		TotalAmount: req.Amount,
		TxId:        internalId,
		ExternalId:  externalId,
	}, nil
}

func (e *Entity) composeTransferRequest(req request.TransferV2Request, token *mochipay.Token, amount *big.Int, template interface{}) mochipay.TransferV2Request {
	recipients := make([]*mochipay.Wallet, 0)
	for _, r := range req.Recipients {
		recipients = append(recipients, &mochipay.Wallet{
			ProfileGlobalId: r,
		})
	}

	amounts := make([]string, 0)
	for range req.Recipients {
		amt, _ := new(big.Float).SetInt(amount).Float64()
		amounts = append(amounts, fmt.Sprintf("%v", strconv.FormatFloat(amt, 'f', -1, 64)))
	}

	// compose metadata
	metadata := map[string]interface{}{
		"message":         req.Message,
		"moniker":         req.Moniker,
		"original_tx_id":  req.OriginalTxId,
		"original_amount": req.OriginalAmount,
		"channel_id":      req.ChannelId,
		"channel_name":    req.ChannelName,
		"channel_url":     req.ChannelUrl,
		"channel_avatar":  req.ChannelAvatar,
		"template":        template,
	}
	if req.Metadata != nil {
		for k, v := range req.Metadata {
			metadata[k] = v
		}
	}

	return mochipay.TransferV2Request{
		From: &mochipay.Wallet{
			ProfileGlobalId: req.Sender,
		},
		Tos:      recipients,
		Platform: req.Platform,
		Metadata: metadata,
		Action:   req.TransferType,
		TokenId:  token.Id,
		Amount:   amounts,
	}
}

func (e *Entity) validateTipRange(guildID string, amount float64, token *mochipay.Token) error {
	if guildID == "" {
		return nil
	}

	config, err := e.repo.GuildConfigTipRange.GetByGuildID(guildID)
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[entity.validateTipRange] repo.GuildConfigTipRange.GetByGuildID() failed")
		return errors.New("failed to validate tip range")
	}

	prices, err := e.svc.CoinGecko.GetCoinPrice([]string{token.CoinGeckoId}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": token.CoinGeckoId}).Error(err, "[entity.validateTipRange] svc.CoinGecko.GetCoinPrice() failed")
		return nil
	}

	// no token tokenPrice data -> ignore
	tokenPrice, ok := prices[token.CoinGeckoId]
	if !ok {
		return nil
	}

	// validate tip amount
	if config.Min != nil && tokenPrice*amount < *config.Min {
		return fmt.Errorf("transfer amount must worth more than $%v", *config.Min)
	}
	if config.Max != nil && tokenPrice*amount > *config.Max {
		return fmt.Errorf("transfer amount must worth less than $%v", *config.Max)
	}

	return nil
}

func (e *Entity) validateTransferBalance(total *big.Int, req request.TransferV2Request) error {
	if !req.All && total.Cmp(big.NewInt(0)) != 1 {
		return errors.New(consts.OffchainTipBotFailReasonInvalidAmount)
	}

	// validate balance
	senderBalance, err := e.svc.MochiPay.GetBalance(req.Sender, req.Token, req.ChainID)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token, "user": req.Sender}).Error(err, "[entity.TransferTokenV2] repo.OffchainTipBotUserBalances.GetUserBalanceByTokenID() failed")
		return err
	}

	if len(senderBalance.Data) == 0 {
		return errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	bal, err := util.StringToBigInt(senderBalance.Data[0].Amount)
	if err != nil {
		return errors.New(consts.OffchainTipBotFailReasonInvalidAmount)
	}

	if bal.Cmp(total) < 0 {
		return errors.New(consts.OffchainTipBotFailReasonNotEnoughBalance)
	}

	return nil
}

func parseTemplate(req request.TransferV2Request) (template interface{}) {
	hashtags := regexp.MustCompile(`#[a-zA-Z0-9_]+`).FindAll([]byte(req.Message), -1)
	for _, hashtag := range hashtags {
		productHashtag, err := e.GetProductHashtag(request.GetProductHashtagRequest{Alias: string(hashtag)[1:]})
		if err != nil || productHashtag == nil {
			e.log.Fields(logger.Fields{"hashtag": string(hashtag)}).Error(err, "[entity.TransferTokenV2] GetProductHashtag() failed")
			continue
		}

		template = productHashtag.ProductHashtag
		break
	}

	if template != nil {
		return
	}

	if req.ThemeId == 0 {
		return
	}

	theme, err := e.repo.ProductTheme.GetByID(req.ThemeId)
	if err != nil {
		e.log.Fields(logger.Fields{"theme_id": req.ThemeId}).Error(err, "[entity.TransferTokenV2] repo.ProductTheme.GetByID() failed")
		return
	}

	hashtag, err := e.repo.ProductHashtag.GetBySlug(theme.Slug)
	if err != nil {
		e.log.Fields(logger.Fields{"slug": theme.Slug}).Error(err, "[entity.TransferTokenV2] repo.ProductHashtag.GetBySlug() failed")
		return
	}

	template = hashtag
	return
}
