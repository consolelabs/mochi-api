package entities

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
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
		solanaCentralizedAddress := e.solana.GetCentralizedWalletAddress()

		swapRoutes, err = e.svc.Kyber.GetSwapRoutesSolana("solana", fromToken.Address, toTokenOverview.Address, stringAmount, solanaCentralizedAddress)
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

func (e *Entity) Swap(req request.SwapRequest) (interface{}, error) {
	// get profile
	profile, err := e.svc.MochiProfile.GetByDiscordID(req.UserDiscordId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-profile.GetByDiscordID] - cannot get profile")
		return nil, err
	}

	// get token
	fromToken, err := e.repo.KyberswapSupportedToken.GetByAddressChain(req.RouteSummary.TokenIn, 0, req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByAddressChain] - cannot get from token")
		return nil, err
	}

	toToken, err := e.repo.KyberswapSupportedToken.GetByAddressChain(req.RouteSummary.TokenOut, 0, req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByAddressChain] - cannot get to token")
		return nil, err
	}

	chainId := util.ConvertChainNameToChainId(req.ChainName)

	// get balance
	balance, err := e.svc.MochiPay.GetBalance(profile.ID, fromToken.Symbol, fmt.Sprintf("%d", chainId))
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.GetBalance] - cannot get balance")
		return nil, err
	}
	if len(balance.Data) == 0 {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.GetBalance] - balance not found")
		return nil, fmt.Errorf("insufficient balance")
	}

	amountSwap, _ := util.StringToBigInt(req.RouteSummary.AmountIn)
	bal, _ := util.StringToBigInt(balance.Data[0].Amount)
	if amountSwap.Cmp(bal) == 1 {
		return nil, fmt.Errorf("insufficient balance")
	}

	// build route kyber
	buildRouteResp, err := e.buildSwapRoute(req)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[swap.buildSwapRoute] - cannot build swap routes")
		return nil, err
	}

	centralizedWalletAddress := e.cfg.CentralizedWalletAddress
	if req.ChainName == "solana" {
		centralizedWalletAddress = e.solana.GetCentralizedWalletAddress()
	}

	// send payload to mochi-pay
	err = e.svc.MochiPay.SwapMochiPay(request.KyberSwapRequest{
		ProfileId:     profile.ID,
		OriginId:      req.UserDiscordId,
		Platform:      consts.PlatformDiscord,
		FromToken:     fromToken.Symbol,
		ToToken:       toToken.Symbol,
		ChainId:       chainId,
		AmountIn:      buildRouteResp.Data.AmountIn,
		AmountOut:     buildRouteResp.Data.AmountOut,
		ChainName:     req.ChainName,
		Address:       centralizedWalletAddress,
		RouterAddress: buildRouteResp.Data.RouterAddress,
		EncodedData:   buildRouteResp.Data.Data,
		Gas:           buildRouteResp.Data.Gas,
		AccountPK:     buildRouteResp.Data.AccountPK,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.SwapMochiPay] - cannot swap mochi pay")
		return nil, err
	}
	return nil, nil
}

func (e *Entity) buildSwapRoute(req request.SwapRequest) (buildRouteResp *response.BuildRoute, err error) {
	solanaCentralizedAddress := e.solana.GetCentralizedWalletAddress()

	if req.ChainName == "solana" {
		buildRouteResp, err = e.svc.Kyber.BuildSwapRoutesSol(req.ChainName, solanaCentralizedAddress, &request.KyberBuildSwapRouteRequest{
			RouteSummary: req.RouteSummary,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.BuildSwapRoutesSol] - cannot build swap routes solana")
			return nil, err
		}

	} else {
		buildRouteResp, err = e.svc.Kyber.BuildSwapRoutes(req.ChainName, &request.KyberBuildSwapRouteRequest{
			Recipient:         e.cfg.CentralizedWalletAddress,
			Sender:            e.cfg.CentralizedWalletAddress,
			Source:            consts.ClientID,
			SkipSimulateTx:    false,
			SlippageTolerance: 50,
			RouteSummary:      req.RouteSummary,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.BuildSwapRoutes] - cannot build swap routes evm")
			return nil, err
		}
	}

	return buildRouteResp, nil
}
