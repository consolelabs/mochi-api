package vaultconfig

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(vaultConfig *model.VaultConfig) error
	GetByGuildId(guildId string) (vaultConfig *model.VaultConfig, err error)
}
