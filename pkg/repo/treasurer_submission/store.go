package treasurersubmission

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(submissions []model.TreasurerSubmission) error
	UpdatePendingSubmission(model *model.TreasurerSubmission) (*model.TreasurerSubmission, error)
	GetByRequestId(requestId, vaultId int64) (submissions []model.TreasurerSubmission, err error)
	GetPendingSubmission(model *model.TreasurerSubmission) (submission *model.TreasurerSubmission, err error)
}
