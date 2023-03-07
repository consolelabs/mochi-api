package coingeckotokenalias

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

func (pg *pg) SetAlias(alias string) error {
	aliasModel := model.CoingeckoTokenAlias{
		Alias: alias,
	}
	return pg.db.Create(aliasModel).Error
}

func (pg *pg) GetOne(alias string) (*model.CoingeckoTokenAlias, error) {
	var model model.CoingeckoTokenAlias
	return &model, pg.db.Where("alias = ?", alias).First(&model).Error
}
