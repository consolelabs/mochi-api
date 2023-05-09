package entities

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) CreateVault(req *request.CreateVaultRequest) (*model.Vault, error) {
	// auto generate vault address when desig mode = false
	walletAddress := ""
	walletNumber := -1
	if !req.DesigMode {
		latestWalletNumber, err := e.repo.Vault.GetLatestWalletNumber()
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - e.repo.Vault.GetLatestWalletNumber failed")
			return nil, err
		}

		account, err := e.vaultwallet.GetAccountByWalletNumber(int(latestWalletNumber.Int64))
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - e.vaultwallet.GetAccountByWalletNumber failed")
			return nil, err
		}

		walletAddress = account.Address.Hex()
		walletNumber = int(latestWalletNumber.Int64) + 1
	}

	vault, err := e.repo.Vault.Create(&model.Vault{
		GuildId:       req.GuildId,
		Name:          req.Name,
		Threshold:     req.Threshold,
		WalletAddress: walletAddress,
		WalletNumber:  int64(walletNumber),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - e.repo.Vault.Create failed")
		return nil, err
	}

	// default for vault creator will be added as treasurer
	_, err = e.repo.Treasurer.Create(&model.Treasurer{
		VaultId:       vault.Id,
		GuildId:       req.GuildId,
		UserDiscordId: req.VaultCreator,
		Role:          consts.VaultCreatorRole,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - add treasurer failed")
		return nil, err
	}

	return vault, nil
}

func (e *Entity) GetVault(guildId string) ([]model.Vault, error) {
	return e.repo.Vault.GetByGuildId(guildId)
}

func (e *Entity) GetVaultInfo() (*model.VaultInfo, error) {
	return e.repo.VaultInfo.Get()
}

func (e *Entity) GetVaultConfigChannel(guildId string) (*model.VaultConfig, error) {
	vaultConfig, err := e.repo.VaultConfig.GetByGuildId(guildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return vaultConfig, nil
}

func (e *Entity) CreateVaultConfigChannel(req *request.CreateVaultConfigChannelRequest) error {
	return e.repo.VaultConfig.Create(&model.VaultConfig{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	})
}

func (e *Entity) CreateConfigThreshold(req *request.CreateConfigThresholdRequest) (*model.Vault, error) {
	vault, err := e.repo.Vault.GetByNameAndGuildId(req.Name, req.GuildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vault not found")
		}
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateConfigThreshold] - e.repo.Vault.GetByNameAndGuildId failed")
		return nil, err
	}

	_, err = e.repo.Vault.UpdateThreshold(&model.Vault{
		GuildId:   req.GuildId,
		Name:      req.Name,
		Threshold: req.Threshold,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateConfigThreshold] - e.repo.Vault.UpdateThreshold failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId:   req.GuildId,
		VaultId:   vault.Id,
		Action:    consts.TreasurerConfigThresholdType,
		Threshold: req.Threshold,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTransaction.Create failed")
		return nil, err
	}
	return vault, nil
}

func (e *Entity) AddTreasurerToVault(req *request.AddTreasurerToVaultRequest) (*model.Treasurer, error) {
	treasurer, err := e.repo.Treasurer.Create(&model.Treasurer{
		GuildId:       req.GuildId,
		VaultId:       req.VaultId,
		UserDiscordId: req.UserDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId: req.GuildId,
		VaultId: req.VaultId,
		Action:  consts.TreasurerAddType,
		Target:  req.UserDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTransaction.Create failed")
		return nil, err
	}

	return treasurer, nil
}

func (e *Entity) TransferVaultToken(req *request.TransferVaultTokenRequest) error {
	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.Vault.GetById failed")
		return err
	}

	token, err := e.svc.MochiPay.GetToken(req.Token, req.Chain)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.svc.MochiPay.GetToken failed")
		return err
	}

	treasurerRequest, err := e.repo.TreasurerRequest.GetById(req.RequestId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.TreasurerRequest.GetById failed")
		return err
	}

	profile, err := e.svc.MochiProfile.GetByDiscordID(treasurerRequest.Requester)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.Profile.GetByDiscordId failed")
		return err
	}

	balance, err := e.vaultwallet.Balance(token, vault.WalletAddress)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.vaultwallet.Balance failed")
		return err
	}

	// check and validate balances
	amountBigIntStr := util.FloatToString(req.Amount, token.Decimal)
	amountBigInt, err := util.StringToBigInt(amountBigIntStr)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - util.StringToBigInt failed")
		return err
	}

	cmp, err := util.CmpBigInt(balance, amountBigInt)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - util.CmpBigInt failed")
		return err
	}

	if cmp == -1 {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - balance not enough")
		return err
	}

	// address = "" aka destination addres = "", use mochi wallet instead

	if req.Address == "" {
		account, err := e.vaultwallet.GetAccountByWalletNumber(int(vault.WalletNumber))
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.vaultwallet.GetAccountByWalletNumber failed")
			return err
		}

		privateKey, err := e.vaultwallet.GetPrivateKeyByAccount(account)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.vaultwallet.GetPrivateKeyByAccount failed")
			return err
		}

		_, err = e.svc.MochiPay.TransferVaultMochiPay(request.MochiPayVaultRequest{
			ProfileId:  profile.ID,
			Amount:     amountBigIntStr,
			To:         e.cfg.CentralizedWalletAddress,
			PrivateKey: privateKey,
			Token:      token.Symbol,
			Chain:      token.Chain.ChainId,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.svc.MochiPay.TransferVaultMochiPay failed")
			return err
		}
	}

	// TODO(trkhoi): implement case has destination address

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId:   req.GuildId,
		VaultId:   req.VaultId,
		Action:    consts.TreasurerTransferType,
		ToAddress: req.Address,
		Amount:    req.Amount,
		Token:     req.Token,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTransaction.Create failed")
		return err
	}

	return nil
}

func (e *Entity) CreateTreasurerResult(req *request.CreateTreasurerResultRequest) error {
	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetById failed")
		return err
	}

	action, thumbnail := prepareParamNotifyTreasurerResult(req.Type)

	msg := prepareMessageNotifyTreasurerResult(req, action, vault.Name, thumbnail)

	err = sendNotifyTreasurerResult(msg, req.ChannelId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - sendNotifyTreasurerResult failed")
		return err
	}

	return nil
}

func prepareParamNotifyTreasurerResult(notifyType string) (action, thumbnail string) {
	action = consts.TreasurerAddedAction
	thumbnail = "https://cdn.discordapp.com/attachments/1090195482506174474/1092703907911847976/image.png"
	if notifyType == consts.TreasurerRemoveType {
		action = consts.TreasurerRemovedAction
		thumbnail = "https://cdn.discordapp.com/attachments/1090195482506174474/1092755046556516394/image.png"
	} else if notifyType == consts.TreasurerTransferType {
		action = consts.TreasurerTransferType
		thumbnail = "https://cdn.discordapp.com/attachments/1003381172178530494/1105400697836556368/vault_open.gif"
	}
	return action, thumbnail
}

func prepareMessageNotifyTreasurerResult(req *request.CreateTreasurerResultRequest, action, vaultName, thumbnail string) (msg discordgo.MessageSend) {
	if req.Address == "" {
		req.Address = "Mochi Wallet"
	}
	if req.Status == consts.TreasurerStatusSuccess {
		description := fmt.Sprintf("<@%s> has been %s to **%s vault**", req.UserDiscordID, action, vaultName)
		if action == consts.TreasurerTransferType {
			description = fmt.Sprintf("%s %s has been sent to `%s`", req.Amount, req.Token, req.Address)
		}

		msg = discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("<:approve_vault:1090242787435356271> Treasurer was successfullly %s", action),
					Description: description,
					Color:       0xFCD3C1,
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: thumbnail,
					},
					Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Type /feedback to report",
					},
				},
			},
		}
	} else {
		description := fmt.Sprintf("<@%s> has not been %s to **%s vault**", req.UserDiscordID, action, vaultName)
		if action == consts.TreasurerTransferType {
			description = fmt.Sprintf("%s %s has not been sent to `%s`", req.Amount, req.Token, req.Address)
		}
		msg = discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("<:revoke:967285238055174195> Treasurer was not %s", action),
					Description: description,
					Color:       0xFCD3C1,
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: thumbnail,
					},
					Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Type /feedback to report",
					},
				},
			},
		}
	}
	return msg
}

func sendNotifyTreasurerResult(msg discordgo.MessageSend, channelId string) error {
	err := e.svc.Discord.SendMessage(channelId, msg)
	if err != nil {
		e.log.Fields(logger.Fields{"msg": msg, "channelId": channelId}).Errorf(err, "[entity.AddTreasurerToVault] - e.svc.Discord.SendMessage failed")
		return err
	}
	return nil
}

func (e *Entity) CreateTreasurerRequest(req *request.CreateTreasurerRequest) (*response.CreateTreasurerRequestResponse, error) {
	// get vault from name and guild id
	vault, err := e.repo.Vault.GetByNameAndGuildId(req.VaultName, req.GuildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vault not exist")
		}
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetByNameAndGuildId failed")
		return nil, err
	}

	// create treasurer request
	treasurerReq, err := e.repo.TreasurerRequest.Create(&model.TreasurerRequest{
		GuildId:       req.GuildId,
		VaultId:       vault.Id,
		UserDiscordId: req.UserDiscordId,
		Message:       req.Message,
		Requester:     req.Requester,
		Type:          req.Type,
		Amount:        req.Amount,
		Chain:         req.Chain,
		Token:         req.Token,
		Address:       req.Address,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}

	// add submission with status pending for all treasurer in vaul
	treasurers, err := e.repo.Treasurer.GetByVaultId(vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.GetByVaultId failed")
		return nil, err
	}
	treasurerSubmission := make([]model.TreasurerSubmission, 0)
	for _, treasurer := range treasurers {
		treasurerSubmission = append(treasurerSubmission, model.TreasurerSubmission{
			VaultId:   vault.Id,
			GuildId:   req.GuildId,
			RequestId: treasurerReq.Id,
			Status:    consts.TreasurerSubmissionStatusPending,
			Submitter: treasurer.UserDiscordId,
		})
	}
	err = e.repo.TreasurerSubmission.Create(treasurerSubmission)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.TreasurerSubmission.Create failed")
		return nil, err
	}

	return &response.CreateTreasurerRequestResponse{
		Request:   *treasurerReq,
		Treasurer: treasurers,
	}, nil
}

func (e *Entity) CreateTreasurerSubmission(req *request.CreateTreasurerSubmission) (resp *response.CreateTreasurerSubmissionResponse, err error) {
	modelSubmission := model.TreasurerSubmission{
		VaultId:   req.VaultId,
		RequestId: req.RequestId,
		Submitter: req.Sumitter,
		Status:    req.Choice,
	}

	// get pending submission
	_, err = e.repo.TreasurerSubmission.GetPendingSubmission(&modelSubmission)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{"req": req}).Infof("[entity.CreateTreasurerSubmission] - submission already processed")
			return nil, fmt.Errorf("submission already processed")
		}
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerSubmission.GetPendingSubmission failed")
		return nil, err
	}

	// update pending submission
	submission, err := e.repo.TreasurerSubmission.UpdatePendingSubmission(&modelSubmission)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerSubmission.GetPendingSubmission failed")
		return nil, err
	}

	// check if total submission >= threshold
	// get all submission of this vault
	submissions, err := e.repo.TreasurerSubmission.GetByRequestId(req.RequestId, req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerSubmission.GetByRequestId failed")
		return nil, err
	}

	totalApprovedSubmission := 0
	totalRejectedSubmisison := 0
	for _, submission := range submissions {
		if submission.Status == consts.TreasurerSubmissionStatusApproved {
			totalApprovedSubmission++
		}
		if submission.Status == consts.TreasurerSubmissionStatusRejected {
			totalRejectedSubmisison++
		}
	}

	submission.GuildId = submissions[0].GuildId
	submission.Vault = submissions[0].Vault
	threshold, _ := strconv.ParseFloat(submissions[0].Vault.Threshold, 64)
	percentage := float64(totalApprovedSubmission) / float64(len(submissions)) * 100
	allowedRejectVote := int64(len(submissions)) - int64(math.Ceil(float64(len(submissions))*threshold/100))

	resp = &response.CreateTreasurerSubmissionResponse{
		Submission: *submission,
		VoteResult: response.VoteResult{
			IsApproved:                false,
			TotalApprovedSubmission:   int64(totalApprovedSubmission),
			TotalRejectedSubmisison:   int64(totalRejectedSubmisison),
			AllowedRejectedSubmisison: allowedRejectVote,
			TotalVote:                 int64(totalApprovedSubmission + totalRejectedSubmisison),
			TotalSubmission:           int64(len(submissions)),
			Percentage:                fmt.Sprintf("%.2f", percentage),
			Threshold:                 fmt.Sprintf("%.2f", threshold),
		},
	}

	if percentage >= threshold {
		resp.VoteResult.IsApproved = true
	}

	// notify treasurer about process voting
	treasurers, err := e.repo.Treasurer.GetByVaultId(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.Treasurer.GetByVaultId failed")
		return nil, err
	}
	for _, treasurer := range treasurers {
		msg := discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "<:bell:1087564962941124679> Mochi notifications",
					Description: fmt.Sprintf("<@%s> %s for request #%d", req.Sumitter, req.Choice, req.RequestId),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Approved",
							Value:  fmt.Sprintf("<:approve_vault:1090242787435356271> `%d/%d`", totalApprovedSubmission, len(submissions)),
							Inline: true,
						},
						{
							Name:   "Rejected",
							Value:  fmt.Sprintf("<:revoke:1077631119073230970> `%d`", totalRejectedSubmisison),
							Inline: true,
						},
						{
							Name:   "Waiting",
							Value:  fmt.Sprintf("<:clock:1080757110146605086> `%d`", len(submissions)-totalApprovedSubmission-totalRejectedSubmisison),
							Inline: true,
						},
					},
					Color: 0xFCD3C1,
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: "https://cdn.discordapp.com/attachments/1090195482506174474/1090905984299442246/image.png",
					},
					Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Type /feedback to report",
					},
				},
			},
		}
		err = e.svc.Discord.SendDM(treasurer.UserDiscordId, msg)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.svc.Discord.SendDM failed")
			continue
		}

		// DM result to user
		if resp.VoteResult.IsApproved {
			action, thumbnail := prepareParamNotifyTreasurerResult(req.Type)
			msg := prepareMessageNotifyTreasurerResult(&request.CreateTreasurerResultRequest{Status: consts.TreasurerStatusSuccess, UserDiscordID: submissions[0].TreasurerRequest.UserDiscordId, Token: submissions[0].TreasurerRequest.Token, Amount: submissions[0].TreasurerRequest.Amount, Address: submissions[0].TreasurerRequest.Address}, action, submissions[0].Vault.Name, thumbnail)
			err = e.svc.Discord.SendDM(treasurer.UserDiscordId, msg)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.svc.Discord.SendDM failed")
				continue
			}
		} else {
			if int64(totalRejectedSubmisison) > allowedRejectVote {
				action, thumbnail := prepareParamNotifyTreasurerResult(req.Type)
				msg := prepareMessageNotifyTreasurerResult(&request.CreateTreasurerResultRequest{Status: consts.TreasurerStatusFail, UserDiscordID: submissions[0].TreasurerRequest.UserDiscordId, Token: submissions[0].TreasurerRequest.Token, Amount: submissions[0].TreasurerRequest.Amount, Address: submissions[0].TreasurerRequest.Address}, action, submissions[0].Vault.Name, thumbnail)
				err = e.svc.Discord.SendDM(treasurer.UserDiscordId, msg)
				if err != nil {
					e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.svc.Discord.SendDM failed")
					continue
				}
			}
		}
	}

	// update request status
	if resp.VoteResult.IsApproved {
		err = e.repo.TreasurerRequest.UpdateStatus(submission.RequestId, consts.TreasurerRequestStatusApproved)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerRequest.UpdateStatus failed")
			return nil, err
		}
	}

	return resp, nil
}

func (e *Entity) RemoveTreasurerFromVault(req *request.AddTreasurerToVaultRequest) (*model.Treasurer, error) {
	treasurer, err := e.repo.Treasurer.Delete(&model.Treasurer{
		GuildId:       req.GuildId,
		VaultId:       req.VaultId,
		UserDiscordId: req.UserDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.RemoveTreasurerFromVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId: req.GuildId,
		VaultId: req.VaultId,
		Action:  consts.TreasurerRemoveType,
		Target:  req.UserDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.RemoveTreasurerFromVault] - e.repo.VaultTransaction.Create failed")
		return nil, err
	}

	return treasurer, nil
}

func (e *Entity) GetVaultDetail(vaultName, guildId string) (*response.VaultDetailResponse, error) {
	vault, err := e.repo.Vault.GetByNameAndGuildId(vaultName, guildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vault not exist")
		}
		e.log.Fields(logger.Fields{"vaultName": vaultName}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetByNameAndGuildId failed")
		return nil, err
	}

	// get treasurers
	treasurers, err := e.repo.Treasurer.GetByGuildIdAndVaultId(guildId, vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.Treasurer.GetByGuildIdAndVaultName failed")
		return nil, err
	}

	// get current request
	currentRequest, err := e.repo.TreasurerRequest.GetCurrentRequest(vault.Id, guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.TreasurerRequest.GetCurrentRequest failed")
		return nil, err
	}

	currentRequestResponse := make([]response.CurrentRequest, 0)
	for _, req := range currentRequest {
		totalApprovedSubmisison := 0
		for _, sub := range req.TreasurerSubmission {
			if sub.Status == consts.TreasurerSubmissionStatusApproved {
				totalApprovedSubmisison++
			}
		}
		currentRequestResponse = append(currentRequestResponse, response.CurrentRequest{
			Target:                  req.UserDiscordId,
			Action:                  util.Capitalize(req.Type),
			Token:                   req.Token,
			Amount:                  req.Amount,
			Address:                 req.Address,
			TotalSubmission:         int64(len(req.TreasurerSubmission)),
			TotalApprovedSubmission: int64(totalApprovedSubmisison),
		})
	}

	// get recent transaction
	recentTransactions, err := e.repo.VaultTransaction.GetRecentTx(vault.Id, guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.VaultTransaction.GetRecentTx failed")
		return nil, err
	}
	recentTxResponse := make([]response.VaultTransaction, 0)
	for _, tx := range recentTransactions {
		recentTxResponse = append(recentTxResponse, response.VaultTransaction{
			Action:    util.Capitalize(strings.Replace(tx.Action, "_", " ", -1)),
			Target:    tx.Target,
			Date:      tx.CreatedAt,
			Amount:    tx.Amount,
			Token:     tx.Token,
			ToAddress: tx.ToAddress,
			Threshold: tx.Threshold,
		})
	}

	return &response.VaultDetailResponse{
		WalletAddress:     vault.WalletAddress,
		EstimatedTotal:    "",
		Balance:           []response.Balance{},
		MyNft:             []response.MyNft{},
		Treasurer:         treasurers,
		RecentTransaction: recentTxResponse,
		CurrentRequest:    currentRequestResponse,
	}, nil
}
