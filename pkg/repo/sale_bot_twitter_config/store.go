package salebottwitterconfig

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(q ListQuery) ([]model.SaleBotTwitterConfig, error)
	Create(cfg *model.SaleBotTwitterConfig) error
}
