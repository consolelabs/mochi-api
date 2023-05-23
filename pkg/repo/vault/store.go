package vault

import (
	"database/sql"

	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	Create(vault *model.Vault) (*model.Vault, error)
	GetByGuildId(guildId string) ([]model.Vault, error)
	UpdateThreshold(vault *model.Vault) (*model.Vault, error)
	GetById(id int64) (vault *model.Vault, err error)
	GetByNameAndGuildId(name string, guildId string) (vault *model.Vault, err error)
	GetLatestWalletNumber() (walletNumber sql.NullInt64, err error)
	List(ListQuery) (vaults []model.Vault, err error)
}
