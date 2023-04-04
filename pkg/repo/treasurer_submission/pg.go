package treasurersubmission

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(submissions []model.TreasurerSubmission) error {
	return pg.db.Create(&submissions).Error
}

func (pg *pg) GetPendingSubmission(model *model.TreasurerSubmission) (submission *model.TreasurerSubmission, err error) {
	return model, pg.db.Where("vault_id = ? and request_id = ? and submitter = ? and status = ?", model.VaultId, model.RequestId, model.Submitter, consts.TreasurerSubmissionStatusPending).First(&submission).Error
}

func (pg *pg) UpdatePendingSubmission(model *model.TreasurerSubmission) (*model.TreasurerSubmission, error) {
	return model, pg.db.Model(&model).Where("vault_id = ? and request_id = ? and submitter = ? and status = ?", model.VaultId, model.RequestId, model.Submitter, consts.TreasurerSubmissionStatusPending).Update("status", model.Status).Error
}
func (pg *pg) GetByRequestId(requestId, vaultId int64) (submissions []model.TreasurerSubmission, err error) {
	return submissions, pg.db.Model(model.TreasurerSubmission{}).Where("request_id = ? and vault_id = ?", requestId, vaultId).Preload(clause.Associations).Find(&submissions).Error
}
