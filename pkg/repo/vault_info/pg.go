package vaultinfo

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

func (pg *pg) Get() (vaultInfo *model.VaultInfo, err error) {
	return vaultInfo, pg.db.First(&vaultInfo).Error
}
