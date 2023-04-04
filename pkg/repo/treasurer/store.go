package treasurer

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(treasurer *model.Treasurer) (*model.Treasurer, error)
	GetByVaultId(vaultId int64) (treasurers []model.Treasurer, err error)
}
