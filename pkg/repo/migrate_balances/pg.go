package migratebalances

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) StoreMigrateBalances(model *model.MigrateBalance) error {
	return pg.db.Save(model).Error
}
