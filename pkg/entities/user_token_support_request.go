package entities

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	tokenSuportReq "github.com/defipod/mochi/pkg/repo/user_token_support_request"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
	"gorm.io/gorm"
)

func (e *Entity) CreateUserTokenSupportRequest(req request.CreateUserTokenSupportRequest) (*model.UserTokenSupportRequest, error) {
	chainIdStr := util.ConvertInputToChainId(req.TokenChain)
	chainId, err := strconv.Atoi(chainIdStr)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(errors.ErrInvalidChain, "invalid chain")
		return nil, errors.ErrInvalidChain
	}
	reqs, err := e.repo.UserTokenSupportRequest.List(tokenSuportReq.ListQuery{TokenChainID: &chainId, TokenAddress: req.TokenAddress})
	if err != nil {
		e.log.Fields(logger.Fields{
			"ChainID":      chainId,
			"TokenAddress": req.TokenAddress,
		}).Error(err, "[entity.CreateUserTokenSupportRequest] repo.UserTokenSupportRequest.List() failed")
		return nil, err
	}
	if len(reqs) > 0 {
		e.log.Fields(logger.Fields{
			"ChainID":      chainId,
			"TokenAddress": req.TokenAddress,
		}).Error(errors.ErrTokenRequestExisted, "[entity.CreateUserTokenSupportRequest] token request already existed")
		return nil, errors.ErrTokenRequestExisted
	}
	tokenReq := &model.UserTokenSupportRequest{
		UserDiscordID: req.UserDiscordID,
		ChannelID:     req.ChannelID,
		MessageID:     req.MessageID,
		TokenName:     req.TokenName,
		TokenAddress:  req.TokenAddress,
		TokenChainID:  chainId,
		Status:        model.TokenSupportPending,
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

func (e *Entity) ApproveTokenSupportRequest(id int) (*model.UserTokenSupportRequest, error) {
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
	if err := e.repo.UserTokenSupportRequest.Update(req); err != nil {
		e.log.Fields(logger.Fields{
			"id":     id,
			"status": status,
		}).Error(err, "[entity.updateStatusTokenRequest] repo.UserTokenSupportRequest.Update() failed")
		return nil, err
	}
	return req, nil
}

func (e *Entity) notifyDiscordTokenRequest(requestID int, req request.CreateUserTokenSupportRequest) error {
	description := fmt.Sprintf("<@%s> wants to add the following token into his/her server.\n\n", req.UserDiscordID) +
		"Token name\n" +
		fmt.Sprintf("```%s```", req.TokenName) +
		"Token address\n" +
		fmt.Sprintf("```%s```", req.TokenAddress) +
		"Chain name\n" +
		fmt.Sprintf("```%s```", req.TokenChain)
	msgSend := discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Token-add submission received",
				Description: description,
				Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				Type:        discordgo.EmbedTypeArticle,
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
		}).Error(err, "[entity.CreateProposalChannelConfig] e.svc.Discord.SendMessage failed")
		return err
	}
	return nil
}
