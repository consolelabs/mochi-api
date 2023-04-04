package entities

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateVault(req *request.CreateVaultRequest) (*model.Vault, error) {
	vault, err := e.repo.Vault.Create(&model.Vault{
		GuildId:   req.GuildId,
		Name:      req.Name,
		Threshold: req.Threshold,
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
	return e.repo.Vault.UpdateThreshold(&model.Vault{
		GuildId:   req.GuildId,
		Name:      req.Name,
		Threshold: req.Threshold,
	})
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

	// send msg to channel
	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetById failed")
		return nil, err
	}

	err = e.svc.Discord.SendMessage(req.ChannelId, discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "<:approve_vault:1090242787435356271> Treasurer was successfullly added",
				Description: fmt.Sprintf("<@%s> has been added to **%s vault**", req.UserDiscordID, vault.Name),
				Color:       0xFCD3C1,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: "https://cdn.discordapp.com/attachments/1090195482506174474/1092703907911847976/image.png",
				},
				Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
				Footer: &discordgo.MessageEmbedFooter{
					Text: "Type /feedback to report",
				},
			},
		},
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.svc.Discord.SendMessage failed")
		return nil, err
	}

	return treasurer, nil
}

func (e *Entity) CreateAddTreasurerRequest(req *request.CreateTreasurerRequest) (*response.CreateTreasurerRequestResponse, error) {
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
	for _, submission := range submissions {
		if submission.Status == consts.TreasurerSubmissionStatusApproved {
			totalApprovedSubmission++
		}
	}

	submission.GuildId = submissions[0].GuildId
	submission.Vault = submissions[0].Vault
	threshold, _ := strconv.ParseFloat(submissions[0].Vault.Threshold, 64)
	percentage := math.Ceil(float64(totalApprovedSubmission)/float64(len(submissions))) * 100

	resp = &response.CreateTreasurerSubmissionResponse{
		Submission: *submission,
		VoteResult: response.VoteResult{
			IsApproved:              false,
			TotalApprovedSubmission: int64(totalApprovedSubmission),
			TotalSubmission:         int64(len(submissions)),
			Percentage:              fmt.Sprintf("%.2f", percentage),
			Threshold:               fmt.Sprintf("%.2f", threshold),
		},
	}

	if percentage >= threshold {
		resp.VoteResult.IsApproved = true
	}

	return resp, nil
}
