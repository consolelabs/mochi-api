package salebotmarketplace

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List() ([]model.SaleBotMarketplace, error)
	GetOne(name string) (*model.SaleBotMarketplace, error)
}
