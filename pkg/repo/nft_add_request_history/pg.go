package nftaddrequesthistory

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetOne(q GetOneQuery) (*model.NftAddRequestHistory, error) {
	result := &model.NftAddRequestHistory{}
	return result, pg.db.Where("address = ? AND chain_id = ?", q.Address, q.ChainID).First(result).Error
}

func (pg *pg) UpsertOne(model model.NftAddRequestHistory) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "chain_id"}},
		DoNothing: true,
	}).Create(&model).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
