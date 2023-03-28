package entities

import (
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
