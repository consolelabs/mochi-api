package entities

import (
	"encoding/json"
	"fmt"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	query "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) getAllChainToken(symbol string, coingeckoId string) (tokens []model.Token, err error) {
	// step 1: look for internal db
	// assuming coingecko where we have most data
	coingeckoTokens, err := e.repo.CoingeckoSupportedTokens.List(query.ListQuery{Symbol: symbol, ID: coingeckoId})
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": symbol, "coingeckoId": coingeckoId}).Error(err, "[repo.CoingeckoSupportedTokens.List] - cannot get data from coingecko")
		return nil, err
	}

	if len(coingeckoTokens) == 0 {
		return nil, nil
	}

	var platforms []model.CoingeckoDetailPlatform
	err = json.Unmarshal(coingeckoTokens[0].DetailPlatforms, &platforms)
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": symbol, "coingeckoId": coingeckoId}).Error(err, "[json.Unmarshal] - cannot unmarshal detail platform")
		return nil, err
	}

	for _, platform := range platforms {
		tokens = append(tokens, model.Token{
			Symbol:  coingeckoTokens[0].Symbol,
			ChainID: int(platform.ChainId),
			Address: platform.Address,
		})
	}

	return tokens, nil
}

func (e *Entity) GetSwapRoutes(req *request.GetSwapRouteRequest) (*response.SwapRouteResponse, error) {
	// step 1: we need to identify which token user want to swap
	// e.g: user input "usdc" -> we need to parse it into any {token_address, chain_id} possibles

	fromTokens, err := e.getAllChainToken(req.From, req.FromTokenId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.getAllChainToken] - cannot get all chain token")
		return nil, err
	}

	toTokens, err := e.getAllChainToken(req.To, req.ToTokenId)
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

	fmt.Println(routes)

	// // step 3: we identiy which route is best for user
	// r, err := e.getBestRoute(routes)
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.getBestRoute] - cannot get best route")
	// 	return nil, err
	// }

	// return to user
	return nil, nil

	// get from token
	// fromTokens, err := e.repo.KyberswapSupportedToken.GetByToken(req.From)
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get data for from token")
	// 	return nil, err
	// }
	// // get to token
	// toTokenOverviews, err := e.repo.KyberswapSupportedToken.GetByToken(req.To)
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get data for to token")
	// 	return nil, err
	// }

	// var fromToken, toToken *model.KyberswapSupportedToken
	// // fromToken and toToken is 2 list of token. Ex: fromToken = [{"eth", "spell"}, {"ftm", "spell"}], toToken = [{"ftm", "usdt"}, {"polygon", "usdt"}]
	// // case 1: found overlap chain -> ex: spell -> usdt has same chain ftm -> return first match
	// endLooking := false
	// for _, from := range fromTokens {
	// 	for _, to := range toTokenOverviews {
	// 		if from.ChainName == to.ChainName {
	// 			fromToken = &from
	// 			toToken = &to
	// 			endLooking = true
	// 			break
	// 		}
	// 	}
	// 	if endLooking {
	// 		break
	// 	}
	// }

	// // case 2: not found overlap -> which means maybe token has not exist in database yet -> query with coingecko
	// if fromToken == nil || toToken == nil {
	// 	// coingecko
	// 	fromToken, toToken, err = e.ImportNonExistToken(req)
	// 	if err != nil {
	// 		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.ImportNonExistToken] - cannot import token")
	// 		return nil, err
	// 	}
	// }

	// // convert string float to string big int
	// amount, _ := strconv.ParseFloat(req.Amount, 64)
	// stringAmount := util.FloatToBigInt(amount, int64(fromToken.Decimals)).String()

	// var swapRoutes *response.KyberSwapRoutes
	// if fromToken.ChainId == 101 || fromToken.ChainName == "solana" {
	// 	swapRoutes, err = e.svc.Kyber.GetSwapRoutesSolana("solana", fromToken.Address, toToken.Address, stringAmount)
	// 	if err != nil {
	// 		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.GetSwapRoutes] - cannot get swap routes")
	// 		return nil, err
	// 	}
	// } else {
	// 	swapRoutes, err = e.svc.Kyber.GetSwapRoutesEVM(fromToken.ChainName, fromToken.Address, toToken.Address, stringAmount)
	// 	if err != nil {
	// 		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.GetSwapRoutes] - cannot get swap routes")
	// 		return nil, err
	// 	}
	// }

	// // case kyber return route not found
	// if swapRoutes.Message == "route not found" || swapRoutes.Code != 0 {
	// 	return nil, errors.ErrKyberRouteNotFound
	// }

	// swapRoutes.Data.TokenIn = *fromToken
	// swapRoutes.Data.TokenOut = *toToken

	// // mapping route
	// newRoute := make([][]response.RouteElement, 0)
	// for _, route := range swapRoutes.Data.RouteSummary.Route {
	// 	newRouteElement := make([]response.RouteElement, 0)
	// 	for _, routeEle := range route {
	// 		toTokenDetail, err := e.repo.KyberswapSupportedToken.GetByAddressChain(routeEle.TokenOut, int64(fromToken.ChainId), fromToken.ChainName)
	// 		if err != nil && err != gorm.ErrRecordNotFound {
	// 			e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByAddressChain] - cannot get to token")
	// 			return nil, err
	// 		}
	// 		if err == gorm.ErrRecordNotFound {
	// 			toTokenDetail = toToken
	// 		}
	// 		newRouteElement = append(newRouteElement, response.RouteElement{
	// 			Pool:              routeEle.Pool,
	// 			TokenIn:           routeEle.TokenIn,
	// 			TokenOut:          routeEle.TokenOut,
	// 			LimitReturnAmount: routeEle.LimitReturnAmount,
	// 			SwapAmount:        routeEle.SwapAmount,
	// 			AmountOut:         routeEle.AmountOut,
	// 			Exchange:          routeEle.Exchange,
	// 			PoolLength:        routeEle.PoolLength,
	// 			PoolType:          routeEle.PoolType,
	// 			PoolExtra:         routeEle.PoolExtra,
	// 			Extra:             routeEle.Extra,
	// 			TokenOutSymbol:    toTokenDetail.Symbol,
	// 		})
	// 	}
	// 	newRoute = append(newRoute, newRouteElement)

	// }
	// return &response.SwapRouteResponse{
	// 	Code:      swapRoutes.Code,
	// 	Message:   swapRoutes.Message,
	// 	ChainName: fromToken.ChainName,
	// 	Data: response.SwapRoute{
	// 		TokenIn:       swapRoutes.Data.TokenIn,
	// 		TokenOut:      swapRoutes.Data.TokenOut,
	// 		RouterAddress: swapRoutes.Data.RouterAddress,
	// 		RouteSummary: response.RouteSummary{
	// 			TokenIn:                      swapRoutes.Data.RouteSummary.TokenIn,
	// 			AmountIn:                     swapRoutes.Data.RouteSummary.AmountIn,
	// 			AmountInUsd:                  swapRoutes.Data.RouteSummary.AmountInUsd,
	// 			TokenInMarketPriceAvailable:  swapRoutes.Data.RouteSummary.TokenInMarketPriceAvailable,
	// 			TokenOut:                     swapRoutes.Data.RouteSummary.TokenOut,
	// 			AmountOut:                    swapRoutes.Data.RouteSummary.AmountOut,
	// 			AmountOutUsd:                 swapRoutes.Data.RouteSummary.AmountOutUsd,
	// 			TokenOutMarketPriceAvailable: swapRoutes.Data.RouteSummary.TokenOutMarketPriceAvailable,
	// 			Gas:                          swapRoutes.Data.RouteSummary.Gas,
	// 			GasPrice:                     swapRoutes.Data.RouteSummary.GasPrice,
	// 			GasUsd:                       swapRoutes.Data.RouteSummary.GasUsd,
	// 			ExtraFee:                     swapRoutes.Data.RouteSummary.ExtraFee,
	// 			Route:                        newRoute,
	// 		},
	// 	},
	// }, nil
}

func (e *Entity) Swap(req request.SwapRequest) (interface{}, error) {
	// get profile
	// profile, err := e.svc.MochiProfile.GetByDiscordID(req.UserDiscordId, true)
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-profile.GetByDiscordID] - cannot get profile")
	// 	return nil, err
	// }

	// // get token
	// fromToken, err := e.repo.KyberswapSupportedToken.GetByAddressChain(req.RouteSummary.TokenIn, 0, req.ChainName)
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByAddressChain] - cannot get from token")
	// 	return nil, err
	// }

	// toToken, err := e.repo.KyberswapSupportedToken.GetByAddressChain(req.RouteSummary.TokenOut, 0, req.ChainName)
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByAddressChain] - cannot get to token")
	// 	return nil, err
	// }

	// chainId := util.ConvertChainNameToChainId(req.ChainName)

	// get balance
	// balance, err := e.svc.MochiPay.GetBalance(profile.ID, fromToken.Symbol, fmt.Sprintf("%d", chainId))
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.GetBalance] - cannot get balance")
	// 	return nil, err
	// }
	// if len(balance.Data) == 0 {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.GetBalance] - balance not found")
	// 	return nil, fmt.Errorf("insufficient balance")
	// }

	// amountSwap, _ := util.StringToBigInt(req.RouteSummary.AmountIn)
	// bal, _ := util.StringToBigInt(balance.Data[0].Amount)
	// if amountSwap.Cmp(bal) == 1 {
	// 	return nil, fmt.Errorf("insufficient balance")
	// }

	// // build route kyber
	// buildRouteResp, err := e.svc.Kyber.BuildSwapRoutes(req.ChainName, &request.KyberBuildSwapRouteRequest{
	// 	Recipient:         e.cfg.CentralizedWalletAddress,
	// 	Sender:            e.cfg.CentralizedWalletAddress,
	// 	Source:            consts.ClientID,
	// 	SkipSimulateTx:    false,
	// 	SlippageTolerance: 50,
	// 	RouteSummary:      req.RouteSummary,
	// })
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.BuildSwapRoutes] - cannot build swap routes")
	// 	return nil, err
	// }

	// // send payload to mochi-pay
	// err = e.svc.MochiPay.SwapMochiPay(request.KyberSwapRequest{
	// 	ProfileId:     profile.ID,
	// 	OriginId:      req.UserDiscordId,
	// 	Platform:      consts.PlatformDiscord,
	// 	FromToken:     fromToken.Symbol,
	// 	ToToken:       toToken.Symbol,
	// 	ChainId:       chainId,
	// 	AmountIn:      buildRouteResp.Data.AmountIn,
	// 	AmountOut:     buildRouteResp.Data.AmountOut,
	// 	ChainName:     req.ChainName,
	// 	Address:       e.cfg.CentralizedWalletAddress,
	// 	RouterAddress: buildRouteResp.Data.RouterAddress,
	// 	EncodedData:   buildRouteResp.Data.Data,
	// 	Gas:           buildRouteResp.Data.Gas,
	// })
	// if err != nil {
	// 	e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.SwapMochiPay] - cannot swap mochi pay")
	// 	return nil, err
	// }
	return nil, nil
}
