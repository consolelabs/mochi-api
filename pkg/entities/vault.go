package entities

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateVault(req *request.CreateVaultRequest) (*model.Vault, error) {
	return e.repo.Vault.Create(&model.Vault{
		GuildId:   req.GuildId,
		Name:      req.Name,
		Threshold: req.Threshold,
	})
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
	addTreasurerReq, err := e.repo.TreasurerRequest.GetById(req.RequestId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("request not exist")
		}
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.TreasurerRequest.GetById failed")
		return nil, err
	}

	// get vault from name and guild id
	vault, err := e.repo.Vault.GetById(addTreasurerReq.VaultId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vault not exist")
		}
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetByNameAndGuildId failed")
		return nil, err
	}

	// create treasurer
	treasurer, err := e.repo.Treasurer.Create(&model.Treasurer{
		GuildId:       addTreasurerReq.GuildId,
		VaultId:       vault.Id,
		UserDiscordId: addTreasurerReq.UserDiscordId,
		Message:       addTreasurerReq.Message,
		RequestId:     addTreasurerReq.Id,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}

	// soft delete treasurer request
	err = e.repo.TreasurerRequest.Delete(&model.TreasurerRequest{
		GuildId:       addTreasurerReq.GuildId,
		VaultId:       addTreasurerReq.VaultId,
		UserDiscordId: addTreasurerReq.UserDiscordId,
	})
	if err != nil {
		// not return here, just log
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.TreasurerRequest.Delete failed")
	}
	return treasurer, nil
}

func (e *Entity) CreateAddTreasurerRequest(req *request.CreateAddTreasurerRequest) (*model.TreasurerRequest, error) {
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
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}
	return treasurerReq, nil
}
