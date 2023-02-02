package entities

import (
	"github.com/defipod/mochi/pkg/model"
	salebottwitterconfig "github.com/defipod/mochi/pkg/repo/sale_bot_twitter_config"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetSaleBotTwitterConfigs(marketplace string) ([]model.SaleBotTwitterConfig, error) {
	return e.repo.SaleBotTwitterConfig.List(salebottwitterconfig.ListQuery{MarketplaceName: marketplace})
}

func (e *Entity) CreateSaleBotTwitterConfig(req request.CreateTwitterSaleConfigRequest) (*model.SaleBotTwitterConfig, error) {
	marketplace, err := e.repo.SaleBotMarketplace.GetOne(req.Marketplace)
	if err != nil {
		e.log.Errorf(err, "[entity.CreateSaleBotTwitterConfigs] repo.SaleBotMarketplace.GetOne() failed")
		return nil, err
	}
	cfg := &model.SaleBotTwitterConfig{
		Address:       req.Address,
		ChainID:       req.ChainID,
		MarketplaceID: marketplace.ID,
	}
	err = e.repo.SaleBotTwitterConfig.Create(cfg)
	if err != nil {
		e.log.Errorf(err, "[entity.CreateSaleBotTwitterConfigs] repo.SaleBotTwitterConfig.Create() failed")
	}
	return cfg, err
}
