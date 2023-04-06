package treasurer

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(treasurer *model.Treasurer) (*model.Treasurer, error)
	GetByVaultId(vaultId int64) (treasurers []model.Treasurer, err error)
	Delete(treasurer *model.Treasurer) (*model.Treasurer, error)
	GetByGuildIdAndVaultId(guildId string, vaultId int64) (treasurer []model.Treasurer, err error)
}
