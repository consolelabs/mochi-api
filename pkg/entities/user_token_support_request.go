package entities

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	tokenSupportReq "github.com/defipod/mochi/pkg/repo/user_token_support_request"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochipay"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetUserRequestTokens(req request.GetUserSupportTokenRequest) (tokens []model.UserTokenSupportRequest, pagination *response.PaginationResponse, err error) {
	page, _ := strconv.Atoi(req.Page)
	size, _ := strconv.Atoi(req.Size)
	tokens, total, err := e.repo.UserTokenSupportRequest.List(tokenSupportReq.ListQuery{Offset: page * size, Limit: size, Status: req.Status})
	if err != nil {
		err = fmt.Errorf("failed to get user requested tokens - err: %v", err)
		return nil, nil, err
	}
	return tokens, &response.PaginationResponse{
		Pagination: model.Pagination{
			Page: int64(page),
			Size: int64(size),
		},
		Total: total,
	}, nil
}

func (e *Entity) CreateUserTokenSupportRequest(req request.CreateUserTokenSupportRequest) (*model.UserTokenSupportRequest, error) {
	l := e.log.Fields(logger.Fields{"req": req})
	chainIdStr := util.ConvertInputToChainId(req.TokenChain)
	chainID, err := strconv.Atoi(chainIdStr)
	if err != nil {
		l.Error(errors.ErrInvalidChain, "invalid chain")
		return nil, errors.ErrInvalidChain
	}
	// check duplicated request
	reqs, _, err := e.repo.UserTokenSupportRequest.List(tokenSupportReq.ListQuery{TokenChainID: &chainID, TokenAddress: req.TokenAddress})
	if err != nil {
		l.Error(err, "[entity.CreateUserTokenSupportRequest] repo.UserTokenSupportRequest.List() failed")
		return nil, err
	}
	if len(reqs) > 0 {
		l.Error(errors.ErrTokenRequestExisted, "[entity.CreateUserTokenSupportRequest] token request already existed")
		return nil, errors.ErrTokenRequestExisted
	}

	// get token info
	token, err := e.findTokenByContractAddress(chainID, req.TokenAddress)
	if err != nil {
		l.Error(err, "[entity.CreateUserTokenSupportRequest] findTokenByContractAddress() failed")
		return nil, errors.ErrTokenNotFound
	}

	// check duplicated from mochi-pay
	found, err := e.svc.MochiPay.GetToken(token.Symbol, fmt.Sprint(chainID))
	if err == nil && found != nil {
		if strings.EqualFold(found.Address, req.TokenAddress) {
			l.Error(err, "[entity.CreateUserTokenSupportRequest] repo.UserTokenSupportRequest.List() failed")
			return nil, errors.ErrTokenRequestExisted
		}
	}

	tokenReq := &model.UserTokenSupportRequest{
		UserDiscordID: req.UserDiscordID,
		GuildID:       req.GuildID,
		ChannelID:     req.ChannelID,
		MessageID:     req.MessageID,
		TokenAddress:  req.TokenAddress,
		TokenChainID:  chainID,
		Status:        model.TokenSupportPending,
		CoinGeckoID:   token.ID,
		TokenName:     token.Name,
		Symbol:        strings.ToUpper(token.Symbol),
		Decimal:       token.Decimal,
		Icon:          token.Image.Large,
	}
	err = e.repo.UserTokenSupportRequest.CreateWithHook(tokenReq, func(id int) error {
		return e.notifyDiscordTokenRequest(id, req)
	})
	if err != nil {
		e.log.Fields(logger.Fields{"requestID": tokenReq.ID}).Error(err, "[entity.CreateUserTokenSupportRequest] notifyDiscordTokenRequest() failed")
		return nil, err
	}
	return tokenReq, nil
}

func (e *Entity) findTokenByContractAddress(chainID int, address string) (*response.GetCoinByContractResponseData, error) {
	var platformID string
	// special handling for solana
	if chainID == 999 {
		platformID = "solana"
	} else {
		platforms, err := e.svc.CoinGecko.GetAssetPlatforms()
		if err != nil {
			e.log.Error(err, "[entity.findTokenByContractAddress] svc.CoinGecko.GetAssetPlatforms() failed")
			return nil, err
		}

		for _, p := range platforms {
			if p.ChainIdentifier != nil && *p.ChainIdentifier == int64(chainID) {
				platformID = p.ID
			}
		}
	}

	// find with coingecko first
	coin, err := e.svc.CoinGecko.GetCoinByContract(platformID, address)
	// if found, return
	if err == nil {
		platformDetail, ok := coin.DetailPlatforms[platformID]
		if ok {
			coin.Decimal = platformDetail.DecimalPlace
		}
		return coin, nil
	}
	// if fail to fetch data -> continue looking up
	if err != nil {
		e.log.Fields(logger.Fields{"ChainID": chainID}).Error(err, "[entity.findTokenByContractAddress] svc.CoinGecko.GetCoinByContract() failed")
		// return nil, err
	}

	// now try to find with covalent
	res, err, status := e.svc.Covalent.GetHistoricalTokenPrices(chainID, "usd", address)
	if err != nil {
		e.log.Fields(logger.Fields{"ChainID": chainID, "status": status}).Error(err, "[entity.findTokenByContractAddress] svc.CoinGecko.GetCoinByContract() failed")
		return nil, err
	}
	if len(res.Data) == 0 {
		e.log.Fields(logger.Fields{"ChainID": chainID, "contract": address}).Error(err, "[entity.findTokenByContractAddress] Token contract not found")
		err = fmt.Errorf("Token contract %s not found", address)
		return nil, err
	}
	return &response.GetCoinByContractResponseData{
		ID:      "",
		Name:    res.Data[0].Name,
		Symbol:  res.Data[0].Symbol,
		Decimal: res.Data[0].Decimals,
	}, nil
}

func (e *Entity) ApproveTokenSupportRequest(id int) (*model.UserTokenSupportRequest, error) {
	req, err := e.repo.UserTokenSupportRequest.Get(id)
	if err != nil {
		e.log.Fields(logger.Fields{
			"id": id,
		}).Error(err, "[entity.ApproveTokenSupportRequest] repo.UserTokenSupportRequest.Get() failed")
		return nil, err
	}

	// TODO: remove after migrating token to mochi pay
	offchainToken := &model.OffchainTipBotToken{
		TokenName:   req.TokenName,
		TokenSymbol: req.Symbol,
		CoinGeckoID: req.CoinGeckoID,
		Icon:        &req.Icon,
		Status:      1,
	}
	err = e.repo.OffchainTipBotTokens.Create(offchainToken)
	if err != nil {
		e.log.Fields(logger.Fields{"token": offchainToken}).Error(err, "[entity.ApproveTokenSupportRequest] repo.OffchainTipBotTokens.Create() failed")
		return nil, err
	}

	// create token in mochi-pay
	err = e.svc.MochiPay.CreateToken(mochipay.CreateTokenRequest{
		Id:          offchainToken.ID.String(),
		Name:        offchainToken.TokenName,
		Symbol:      offchainToken.TokenSymbol,
		Decimal:     int64(req.Decimal),
		ChainId:     fmt.Sprint(req.TokenChainID),
		Address:     req.TokenAddress,
		Icon:        req.Icon,
		CoinGeckoId: req.CoinGeckoID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"token": offchainToken}).Error(err, "[entity.ApproveTokenSupportRequest] svc.MochiPay.CreateToken() failed")
		return nil, err
	}

	return e.updateStatusTokenRequest(id, model.TokenSupportApproved)
}

func (e *Entity) RejectTokenSupportRequest(id int) (*model.UserTokenSupportRequest, error) {
	return e.updateStatusTokenRequest(id, model.TokenSupportRejected)
}

func (e *Entity) updateStatusTokenRequest(id int, status model.TokenSupportRequestStatus) (*model.UserTokenSupportRequest, error) {
	req, err := e.repo.UserTokenSupportRequest.Get(id)
	if err != nil {
		e.log.Fields(logger.Fields{
			"id":     id,
			"status": status,
		}).Error(err, "[entity.updateStatusTokenRequest] repo.UserTokenSupportRequest.Get() failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}
	req.Status = status
	err = e.repo.UserTokenSupportRequest.UpdateWithHook(req, func(id int) error {
		switch status {
		case model.TokenSupportApproved:
			return e.notifyDiscordTokenApproved(*req)
		case model.TokenSupportRejected:
			return e.notifyDiscordTokenRejected(*req)
		default:
			return fmt.Errorf("invalid token support status")
		}
	})
	if err != nil {
		e.log.Fields(logger.Fields{
			"id":     id,
			"status": status,
		}).Error(err, "[entity.updateStatusTokenRequest] repo.UserTokenSupportRequest.Update() failed")
		return nil, err
	}
	return req, nil
}

func (e *Entity) notifyDiscordTokenRequest(requestID int, req request.CreateUserTokenSupportRequest) error {
	guild, _ := e.svc.Discord.GetGuild(req.GuildID)
	user, _ := e.svc.Discord.GetUser(req.UserDiscordID)

	guildName := req.GuildID
	userName := req.UserDiscordID
	if guild != nil {
		guildName = guild.Name
	}
	if user != nil {
		userName = fmt.Sprintf("%s#%s", user.Username, user.Discriminator)
	}

	description := fmt.Sprintf("<@%s> wants to add the following token into his/her server.\n\n", req.UserDiscordID)
	fields := []*discordgo.MessageEmbedField{
		{Name: "Guild", Value: fmt.Sprintf("`%s`", guildName)},
		{Name: "Submitter", Value: fmt.Sprintf("`%s`", userName)},
		{Name: "Token", Value: fmt.Sprintf("`%s`", req.TokenAddress)},
		{Name: "Chain", Value: fmt.Sprintf("`%s`", strings.ToUpper(req.TokenChain))},
	}
	msgSend := discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Token-add submission received",
				Description: description,
				Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				Type:        discordgo.EmbedTypeArticle,
				Fields:      fields,
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Approve",
						Style:    discordgo.SuccessButton,
						Disabled: false,
						CustomID: fmt.Sprintf("token-request-approve-%d", requestID),
					},
					discordgo.Button{
						Label:    "Reject",
						Style:    discordgo.DangerButton,
						Disabled: false,
						CustomID: fmt.Sprintf("token-request-reject-%d", requestID),
					},
				},
			},
		},
	}
	if err := e.svc.Discord.SendMessage(e.cfg.MochiTokenRequestChannelID, msgSend); err != nil {
		e.log.Fields(logger.Fields{
			"guidelineChannelID": e.cfg.MochiTokenRequestChannelID,
			"msg":                msgSend,
		}).Error(err, "[entity.notifyDiscordTokenRequest] e.svc.Discord.SendMessage failed")
		return err
	}
	return nil
}

func (e *Entity) notifyDiscordTokenApproved(req model.UserTokenSupportRequest) error {
	description := fmt.Sprintf("Your token request for %s has been approved! Now you can make %s transaction with $tip and $airdrop! <:pumpeet:930840081554624632>", req.Symbol, req.Symbol)
	msgSend := discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "<:approve:1013775501757780098> Your token has been approved",
				Description: description,
				Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				Type:        discordgo.EmbedTypeArticle,
			},
		},
	}
	if err := e.svc.Discord.SendDM(req.UserDiscordID, msgSend); err != nil {
		e.log.Fields(logger.Fields{
			"guidelineChannelID": e.cfg.MochiTokenRequestChannelID,
			"msg":                msgSend,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.svc.Discord.SendDM failed")
		return err
	}
	return nil
}

func (e *Entity) notifyDiscordTokenRejected(req model.UserTokenSupportRequest) error {
	description := fmt.Sprintf("Because of some technical barrier, we regret to inform you that your token %s can’t be supported!\n", req.TokenAddress) +
		"Please check out and try some supported token by $token list. <:nekolove:>"
	msgSend := discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "<:revoke:967285238055174195> Your token has been rejected",
				Description: description,
				Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				Type:        discordgo.EmbedTypeArticle,
			},
		},
	}
	if err := e.svc.Discord.SendDM(req.UserDiscordID, msgSend); err != nil {
		e.log.Fields(logger.Fields{
			"guidelineChannelID": e.cfg.MochiTokenRequestChannelID,
			"msg":                msgSend,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.svc.Discord.SendDM failed")
		return err
	}
	return nil
}
