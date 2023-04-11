package entities

import (
	"strconv"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetSwapRoutes(req *request.GetSwapRouteRequest) (*response.SwapRouteResponse, error) {
	// get from token
	fromToken, err := e.repo.KyberswapSupportedToken.GetByTokenChain(req.From, int64(req.ChainId), req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get from token")
		return nil, err
	}
	// get to token
	toTokenOverview, err := e.repo.KyberswapSupportedToken.GetByTokenChain(req.To, int64(req.ChainId), req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get to token")
		return nil, err
	}

	// convert string float to string big int
	amount, _ := strconv.ParseFloat(req.Amount, 64)
	stringAmount := util.FloatToBigInt(amount, int64(fromToken.Decimals)).String()

	var swapRoutes *response.KyberSwapRoutes
	if req.ChainId == 101 || req.ChainName == "solana" {
		swapRoutes, err = e.svc.Kyber.GetSwapRoutesSolana("solana", fromToken.Address, toTokenOverview.Address, stringAmount)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.GetSwapRoutes] - cannot get swap routes")
			return nil, err
		}
	} else {
		swapRoutes, err = e.svc.Kyber.GetSwapRoutesEVM(fromToken.ChainName, fromToken.Address, toTokenOverview.Address, stringAmount)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.GetSwapRoutes] - cannot get swap routes")
			return nil, err
		}
	}

	swapRoutes.Data.TokenIn = *fromToken
	swapRoutes.Data.TokenOut = *toTokenOverview

	// mapping route
	newRoute := make([][]response.RouteElement, 0)
	for _, route := range swapRoutes.Data.RouteSummary.Route {
		newRouteElement := make([]response.RouteElement, 0)
		for _, routeEle := range route {
			toToken, err := e.repo.KyberswapSupportedToken.GetByAddressChain(routeEle.TokenOut, int64(req.ChainId), req.ChainName)
			if err != nil && err != gorm.ErrRecordNotFound {
				e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByAddressChain] - cannot get to token")
				return nil, err
			}
			if err == gorm.ErrRecordNotFound {
				toToken = toTokenOverview
			}
			newRouteElement = append(newRouteElement, response.RouteElement{
				Pool:              routeEle.Pool,
				TokenIn:           routeEle.TokenIn,
				TokenOut:          routeEle.TokenOut,
				LimitReturnAmount: routeEle.LimitReturnAmount,
				SwapAmount:        routeEle.SwapAmount,
				AmountOut:         routeEle.AmountOut,
				Exchange:          routeEle.Exchange,
				PoolLength:        routeEle.PoolLength,
				PoolType:          routeEle.PoolType,
				PoolExtra:         routeEle.PoolExtra,
				Extra:             routeEle.Extra,
				TokenOutSymbol:    toToken.Symbol,
			})
		}
		newRoute = append(newRoute, newRouteElement)

	}
	return &response.SwapRouteResponse{
		Code:    swapRoutes.Code,
		Message: swapRoutes.Message,
		Data: response.SwapRoute{
			TokenIn:       swapRoutes.Data.TokenIn,
			TokenOut:      swapRoutes.Data.TokenOut,
			RouterAddress: swapRoutes.Data.RouterAddress,
			RouteSummary: response.RouteSummary{
				TokenIn:                      swapRoutes.Data.RouteSummary.TokenIn,
				AmountIn:                     swapRoutes.Data.RouteSummary.AmountIn,
				AmountInUsd:                  swapRoutes.Data.RouteSummary.AmountInUsd,
				TokenInMarketPriceAvailable:  swapRoutes.Data.RouteSummary.TokenInMarketPriceAvailable,
				TokenOut:                     swapRoutes.Data.RouteSummary.TokenOut,
				AmountOut:                    swapRoutes.Data.RouteSummary.AmountOut,
				AmountOutUsd:                 swapRoutes.Data.RouteSummary.AmountOutUsd,
				TokenOutMarketPriceAvailable: swapRoutes.Data.RouteSummary.TokenOutMarketPriceAvailable,
				Gas:                          swapRoutes.Data.RouteSummary.Gas,
				GasPrice:                     swapRoutes.Data.RouteSummary.GasPrice,
				GasUsd:                       swapRoutes.Data.RouteSummary.GasUsd,
				ExtraFee:                     swapRoutes.Data.RouteSummary.ExtraFee,
				Route:                        newRoute,
			},
		},
	}, nil
}

func (e *Entity) BuildSwapRoutes(req request.BuildRouteRequest) (*response.BuildRoute, error) {
	buildRouteResp, err := e.svc.Kyber.BuildSwapRoutes(req.ChainName, &request.KyberBuildSwapRouteRequest{
		Recipient:         req.Recipient,
		Sender:            req.Sender,
		Source:            "kyberswap",
		SkipSimulateTx:    false,
		SlippageTolerance: 50,
		RouteSummary:      req.RouteSummary,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.BuildSwapRoutes] - cannot build swap routes")
		return nil, err
	}

	return buildRouteResp, nil
}

func (e *Entity) Swap(req request.SwapRequest) (interface{}, error) {
	// get from token
	fromToken, err := e.repo.KyberswapSupportedToken.GetByTokenChain(req.From, 0, req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get from token")
		return nil, err
	}
	// get to token
	toToken, err := e.repo.KyberswapSupportedToken.GetByTokenChain(req.To, 0, req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get to token")
		return nil, err
	}
	// convert string float to string big int
	amount, _ := strconv.ParseFloat(req.Amount, 64)
	bigIntAmount := util.FloatToBigInt(amount, int64(fromToken.Decimals))

	e.svc.Abi.SwapTokenOnKyber(request.KyberSwapRequest{
		FromTokenAddress:   fromToken.Address,
		ToTokenAddress:     toToken.Address,
		Amount:             bigIntAmount,
		ChainName:          req.ChainName,
		CentralizedAddress: "0x140dd183e18ba39bd9BE82286ea2d96fdC48117A",
		RouterAddress:      req.RouterAddress,
		EncodedData:        req.EncodedData,
	})
	return nil, nil
}
