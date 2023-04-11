package kyberswapsupportedtokens

import (
	"strings"

	"gorm.io/gorm"

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

func (pg *pg) GetByTokenChain(symbol string, chainId int64, chainName string) (model *model.KyberswapSupportedToken, err error) {
	return model, pg.db.Where("lower(symbol) = ? AND (chain_id = ? OR chain_name = ?)", strings.ToLower(symbol), chainId, chainName).First(&model).Error
}

func (pg *pg) GetByAddressChain(address string, chainId int64, chainName string) (model *model.KyberswapSupportedToken, err error) {
	return model, pg.db.Where("lower(address) = lower(?) AND (chain_id = ? OR chain_name = ?)", address, chainId, chainName).First(&model).Error
}
