package entities

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	tokenSuportReq "github.com/defipod/mochi/pkg/repo/user_token_support_request"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
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

func (e *Entity) notifyDiscordTokenApproved(req model.UserTokenSupportRequest) error {
	description := fmt.Sprintf("Your token request for %s has been approved! Now you can make %s transaction with $tip and $airdrop! <:pumpeet:930840081554624632>", req.TokenName, req.TokenName)
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
	if err := e.svc.Discord.SendMessage(req.ChannelID, msgSend); err != nil {
		e.log.Fields(logger.Fields{
			"guidelineChannelID": e.cfg.MochiTokenRequestChannelID,
			"msg":                msgSend,
		}).Error(err, "[entity.CreateProposalChannelConfig] e.svc.Discord.SendMessage failed")
		return err
	}
	return nil
}

func (e *Entity) notifyDiscordTokenRejected(req model.UserTokenSupportRequest) error {
	description := fmt.Sprintf("Because of some technical barrier, we regret to inform you that your token %s can’t be supported!\n", req.TokenName) +
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
		}).Error(err, "[entity.CreateProposalChannelConfig] e.svc.Discord.SendMessage failed")
		return err
	}
	return nil
}
