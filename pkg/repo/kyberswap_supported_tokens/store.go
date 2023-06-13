package kyberswapsupportedtokens

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateBatchs(models []model.KyberswapSupportedToken) error
	Create(model *model.KyberswapSupportedToken) (*model.KyberswapSupportedToken, error)
	GetByTokenChain(symbol string, chainId int64, chainName string) (model *model.KyberswapSupportedToken, err error)
	GetByAddressChain(address string, chainId int64, chainName string) (model *model.KyberswapSupportedToken, err error)
	GetByToken(symbol string) (tokens []model.KyberswapSupportedToken, err error)
	Upsert(token *model.KyberswapSupportedToken) error
}
