package treasurersubmission

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(submissions []model.VaultSubmission) error
	UpdatePendingSubmission(model *model.VaultSubmission) (*model.VaultSubmission, error)
	GetByRequestId(requestId, vaultId int64) (submissions []model.VaultSubmission, err error)
	GetPendingSubmission(model *model.VaultSubmission) (submission *model.VaultSubmission, err error)
}
