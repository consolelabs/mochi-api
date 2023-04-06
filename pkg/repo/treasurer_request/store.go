package treasurerrequest

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(treasurerReq *model.TreasurerRequest) (*model.TreasurerRequest, error)
	GetById(id int64) (treasurerReq *model.TreasurerRequest, err error)
	Delete(model *model.TreasurerRequest) error
	UpdateStatus(requestId int64, status bool) error
	GetCurrentRequest(vaultId int64, guildId string) (treasurerReq []model.TreasurerRequest, err error)
}
