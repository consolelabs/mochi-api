package entities

import (
	"encoding/json"
	"fmt"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	query "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) getAllChainToken(symbol string) (tokens []model.Token, err error) {
	// step 1: look for internal db
	// assuming coingecko where we have most data
	coingeckoTokens, err := e.repo.CoingeckoSupportedTokens.List(query.ListQuery{Symbol: symbol})
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[repo.CoingeckoSupportedTokens.List] - cannot get data from coingecko")
		return nil, err
	}

	newCoingeckoTokens, err := e.UpsertDetailPlatforms(coingeckoTokens)
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[entity.UpsertDetailPlatforms] - cannot upsert detail platform")
		return nil, err
	}

	if len(newCoingeckoTokens) == 0 {
		return nil, nil
	}

	for _, token := range newCoingeckoTokens {
		var platforms []CoingeckoDetailPlatform
		err = json.Unmarshal(token.DetailPlatforms, &platforms)
		if err != nil {
			e.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[json.Unmarshal] - cannot unmarshal detail platform")
			return nil, err
		}

		for _, platform := range platforms {
			tokens = append(tokens, model.Token{
				Name:        token.Name,
				Symbol:      token.Symbol,
				ChainID:     int(platform.ChainId),
				Address:     platform.Address,
				Decimals:    int(platform.Decimal),
				CoinGeckoID: token.ID,
			})
		}
	}

	return tokens, nil
}

func (e *Entity) UpsertDetailPlatforms(coins []model.CoingeckoSupportedTokens) (newCoins []model.CoingeckoSupportedTokens, err error) {
	for _, coin := range coins {
		if coin.DetailPlatforms != nil {
			newCoins = append(newCoins, coin)
			continue
		}

		platforms := make([]CoingeckoDetailPlatform, 0)

		coinDetail, err, _ := e.svc.CoinGecko.GetCoin(coin.ID)
		if err != nil {
			e.log.Fields(logger.Fields{"coinGeckoId": coin.ID}).Error(err, "[entity.UpsertAllChainTokenData] e.svc.CoinGecko.GetCoin failed")
			continue
		}

		for platform := range coinDetail.DetailPlatforms {
			chainId := util.ConvertCoingeckoChain(platform)
			// case chain not supported yet
			if int(chainId) == 0 {
				continue
			}

			// case chain supported
			platforms = append(platforms, CoingeckoDetailPlatform{
				ChainId: chainId,
				Address: coinDetail.DetailPlatforms[platform].ContractAddress,
				Decimal: int64(coinDetail.DetailPlatforms[platform].DecimalPlace),
			})
		}

		bytedetailPlatforms, err := json.Marshal(platforms)
		if err != nil {
			return coins, err
		}

		coin.DetailPlatforms = bytedetailPlatforms
		_, err = e.repo.CoingeckoSupportedTokens.Upsert(&coin)
		if err != nil {
			e.log.Fields(logger.Fields{"coinGeckoId": coin.ID}).Error(err, "[entity.UpsertAllChainTokenData] e.repo.CoingeckoSupportedTokens.Upsert failed")
			return coins, err
		}
		newCoins = append(newCoins, coin)
	}

	return newCoins, nil
}

func (e *Entity) GetSwapRoutes(req *request.GetSwapRouteRequest) (*response.SwapRouteResponse, error) {
	// step 1: we need to identify which token user want to swap
	// e.g: user input "usdc" -> we need to parse it into any {token_address, chain_id} possibles

	fromTokens, err := e.getAllChainToken(req.From)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.getAllChainToken] - cannot get all chain token")
		return nil, err
	}

	toTokens, err := e.getAllChainToken(req.To)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.getAllChainToken] - cannot get all chain token")
		return nil, err
	}

	amount := util.FloatToString(req.Amount, 18)
	// step 2: now we have 2 set of tokens, we need to find the route
	routes, err := e.svc.Swap.GetAllRoutes(fromTokens, toTokens, amount)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[swap.GetAllRoutes] - cannot get all route")
		return nil, err
	}

	// // step 3: we identiy which route is best for user
	r, err := e.svc.Swap.GetBestRoute(routes)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.getBestRoute] - cannot get best route")
		return nil, err
	}

	toRoutes := e.formatRouteSwap(req, r)

	return toRoutes, nil

}

func (e *Entity) formatRouteSwap(req *request.GetSwapRouteRequest, swapRoutes *response.ProviderSwapRoutes) *response.SwapRouteResponse {
	if swapRoutes == nil {
		return &response.SwapRouteResponse{}
	}

	newRoute := make([][]response.RouteElement, 0)
	for _, route := range swapRoutes.Data.RouteSummary.Route {
		newRouteElement := make([]response.RouteElement, 0)
		for _, routeEle := range route {
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
				TokenOutSymbol:    req.To,
			})
		}
		newRoute = append(newRoute, newRouteElement)

	}
	return &response.SwapRouteResponse{
		Code:    swapRoutes.Code,
		Message: swapRoutes.Message,
		// ChainName: fromToken.ChainName,
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
	}
}

func (e *Entity) Swap(req request.SwapRequest) (interface{}, error) {
	// get profile
	profile, err := e.svc.MochiProfile.GetByDiscordID(req.UserDiscordId, true)
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

	chainId := int64(0)

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
	buildRouteResp, err := e.svc.Kyber.BuildSwapRoutes(req.ChainName, &request.KyberBuildSwapRouteRequest{
		Recipient:         e.cfg.CentralizedWalletAddress,
		Sender:            e.cfg.CentralizedWalletAddress,
		Source:            consts.ClientID,
		SkipSimulateTx:    false,
		SlippageTolerance: 500,
		RouteSummary:      req.RouteSummary,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.BuildSwapRoutes] - cannot build swap routes")
		return nil, err
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
		Address:       e.cfg.CentralizedWalletAddress,
		RouterAddress: buildRouteResp.Data.RouterAddress,
		EncodedData:   buildRouteResp.Data.Data,
		Gas:           buildRouteResp.Data.Gas,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.SwapMochiPay] - cannot swap mochi pay")
		return nil, err
	}
	return nil, nil
}
