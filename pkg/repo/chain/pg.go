package chain

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

func (pg *pg) GetAll() ([]model.Chain, error) {
	var chains []model.Chain
	return chains, pg.db.Find(&chains).Error
}

func (pg *pg) GetByID(id int) (model.Chain, error) {
	var chain model.Chain
	return chain, pg.db.Where("id = ?", id).First(&chain).Error
}

func (pg *pg) GetByShortName(shortName string) (*model.Chain, error) {
	chain := &model.Chain{}
	if err := pg.db.First(chain, "upper(short_name) = upper(?)", shortName).Error; err != nil {
		return nil, err
	}
	return chain, nil
}

func (pg *pg) GetOne(q GetOneQuery) (chain *model.Chain, err error) {
	db := pg.db
	if q.CoingeckoId != "" {
		db = db.Where("coin_gecko_id = ?", q.CoingeckoId)
	}
	if q.Type != "" {
		db = db.Where("\"type\" = ?", q.Type)
	}
	return chain, db.First(&chain).Error
}
