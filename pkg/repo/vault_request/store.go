package treasurerrequest

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(treasurerReq *model.VaultRequest) (*model.VaultRequest, error)
	GetById(id int64) (treasurerReq *model.VaultRequest, err error)
	Delete(model *model.VaultRequest) error
	UpdateStatus(requestId int64, status bool) error
	GetCurrentRequest(vaultId int64, guildId string) (treasurerReq []model.VaultRequest, err error)
}
