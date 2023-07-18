package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetInvestList(req *request.GetInvestListRequest) (*response.GetInvestListResponse, error) {
	resp, err := e.svc.Krystal.GetEarningOptions(req.Platforms, req.ChainIds, req.Types, req.Status, req.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"request": req}).Error(err, "[entities.GetInvestList] - svc.Krystal.GetEarningOptions failed")
		return nil, err
	}

	ivestItems := make([]response.InvestItem, 0)
	for _, item := range resp.Result {
		chain := response.InvestChain{
			ID:   item.Chain.ID,
			Logo: item.Chain.Logo,
			Name: item.Chain.Name,
		}

		platforms := make([]response.InvestPlatforms, 0)
		for _, platform := range item.Platforms {
			platforms = append(platforms, response.InvestPlatforms{
				Apy:       platform.Apy,
				Desc:      platform.Desc,
				Logo:      platform.Logo,
				Name:      platform.Name,
				RewardAPY: platform.RewardAPY,
				Status: response.InvestStatus{
					Detail: platform.Status.Detail,
					Value:  platform.Status.Value,
				},
				Tvl:  platform.Tvl,
				Type: platform.Type,
			})
		}

		token := response.InvestToken{
			Address:  item.Token.Address,
			Symbol:   item.Token.Symbol,
			Name:     item.Token.Name,
			Decimals: item.Token.Decimals,
		}

		ivestItems = append(ivestItems, response.InvestItem{
			Apy:       item.Apy,
			Tvl:       item.Tvl,
			Chain:     chain,
			Platforms: platforms,
			Token:     token,
		})
	}

	return &response.GetInvestListResponse{
		Data: ivestItems,
	}, nil
}
