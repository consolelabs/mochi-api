package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/krystal"
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

func (e *Entity) OnchainInvestData(req *request.OnchainInvestDataRequest) (*response.OnchainInvestDataResponse, error) {
	resp, err := e.svc.Krystal.BuildStakeTx(krystal.BuildStakeTxReq{
		Platform:     req.Platform,
		ChainID:      req.ChainID,
		EarningType:  req.Type,
		TokenAddress: req.TokenAddress,
		TokenAmount:  req.TokenAmount,
		UserAddress:  req.UserAddress,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"request": req}).Error(err, "[entities.OnchainInvestData] - svc.Krystal.BuildStakeTx failed")
		return nil, err
	}

	return &response.OnchainInvestDataResponse{
		Data: response.OnchainInvestData{
			TxObject: response.TxObject{
				From:     resp.TxObject.From,
				To:       resp.TxObject.To,
				Value:    resp.TxObject.Value,
				Data:     resp.TxObject.Data,
				GasPrice: resp.TxObject.GasPrice,
				Nonce:    resp.TxObject.Nonce,
				GasLimit: resp.TxObject.GasLimit,
			},
		},
	}, nil
}
