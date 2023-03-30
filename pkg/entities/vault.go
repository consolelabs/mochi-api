package entities

import (
	"gorm.io/gorm"

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
