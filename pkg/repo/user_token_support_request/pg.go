package user_token_support_request

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) ([]model.UserTokenSupportRequest, error) {
	var requests []model.UserTokenSupportRequest
	db := pg.db
	if q.TokenAddress != "" {
		db = db.Where("token_address = ?", q.TokenAddress)
	}
	if q.TokenChainID != nil {
		db = db.Where("token_chain_id = ?", *q.TokenChainID)
	}
	return requests, db.Find(&requests).Error
}

func (pg *pg) Create(req *model.UserTokenSupportRequest) error {
	return pg.db.Create(&req).Error
}
func (pg *pg) CreateWithHook(req *model.UserTokenSupportRequest, afterCreateFn func(id int) error) error {
	tx := pg.db.Begin()
	tx = tx.Create(&req)
	if err := afterCreateFn(req.ID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) Get(id int) (req *model.UserTokenSupportRequest, err error) {
	return req, pg.db.First(&req, id).Error
}

func (pg *pg) Update(req *model.UserTokenSupportRequest) error {
	return pg.db.Save(&req).Error
}

func (pg *pg) UpdateWithHook(req *model.UserTokenSupportRequest, afterUpdateFn func(id int) error) error {
	tx := pg.db.Begin()
	tx = tx.Save(&req)
	if err := afterUpdateFn(req.ID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) Delete(id int) error {
	return pg.db.Delete(&model.UserTokenSupportRequest{}, id).Error
}
