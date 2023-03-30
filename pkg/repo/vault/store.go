package vault

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(vault *model.Vault) (*model.Vault, error)
	GetByGuildId(guildId string) ([]model.Vault, error)
	UpdateThreshold(vault *model.Vault) (*model.Vault, error)
}
