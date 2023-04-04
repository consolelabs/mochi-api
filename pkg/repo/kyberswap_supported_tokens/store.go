package kyberswapsupportedtokens

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateBatchs(models []model.KyberswapSupportedToken) error
	GetByTokenChain(symbol string, chainId int64, chainName string) (model *model.KyberswapSupportedToken, err error)
}
