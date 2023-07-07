package treasurer

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(treasurer *model.VaultTreasurer) (*model.VaultTreasurer, error)
	GetByVaultId(vaultId int64) (treasurers []model.VaultTreasurer, err error)
	Delete(treasurer *model.VaultTreasurer) (*model.VaultTreasurer, error)
	GetByGuildIdAndVaultId(guildId string, vaultId int64) (treasurer []model.VaultTreasurer, err error)
}
