package kyberswapsupportedtokens

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) CreateBatchs(models []model.KyberswapSupportedToken) error {
	return pg.db.Create(&models).Error
}

func (pg *pg) Create(model *model.KyberswapSupportedToken) (*model.KyberswapSupportedToken, error) {
	return model, pg.db.Save(&model).Error
}

func (pg *pg) GetByTokenChain(symbol string, chainId int64, chainName string) (model *model.KyberswapSupportedToken, err error) {
	return model, pg.db.Where("lower(symbol) = ? AND (chain_id = ? OR chain_name = ?)", strings.ToLower(symbol), chainId, chainName).First(&model).Error
}

func (pg *pg) GetByAddressChain(address string, chainId int64, chainName string) (model *model.KyberswapSupportedToken, err error) {
	return model, pg.db.Where("lower(address) = lower(?) AND (chain_id = ? OR chain_name = ?)", address, chainId, chainName).First(&model).Error
}

func (pg *pg) GetByToken(symbol string) (tokens []model.KyberswapSupportedToken, err error) {
	return tokens, pg.db.Where("lower(symbol) = ?", strings.ToLower(symbol)).Find(&tokens).Error
}

func (pg *pg) Upsert(token *model.KyberswapSupportedToken) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chain_name"}, {Name: "address"}, {Name: "symbol"}},
		UpdateAll: true,
	}).Create(token).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
