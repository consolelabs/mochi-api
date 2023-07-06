package entities

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	query "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	mochipayrequest "github.com/defipod/mochi/pkg/service/mochipay"
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

func getAllChainDetailPlatform(platforms []CoingeckoDetailPlatform) (chains []int64) {
	for _, platform := range platforms {
		chains = append(chains, platform.ChainId)
	}

	return chains
}
func (e *Entity) UpsertDetailPlatforms(coins []model.CoingeckoSupportedTokens) (newCoins []model.CoingeckoSupportedTokens, err error) {
	for _, coin := range coins {
		if coin.DetailPlatforms != nil {
			var platforms []CoingeckoDetailPlatform
			err := json.Unmarshal(coin.DetailPlatforms, &platforms)
			if err != nil {
				return coins, err
			}

			chainsOfCoin := getAllChainDetailPlatform(platforms)
			// TODO(trkhoi): remove when coingecko table has enough solana address
			// if detailPlatform have solana -> continue
			if slices.Contains(chainsOfCoin, int64(999)) {
				newCoins = append(newCoins, coin)
				continue
			}

			// if not, get data from coingecko
			coinDetail, err, _ := e.svc.CoinGecko.GetCoin(coin.ID)
			if err != nil {
				e.log.Fields(logger.Fields{"coinGeckoId": coin.ID}).Error(err, "[entity.UpsertAllChainTokenData] e.svc.CoinGecko.GetCoin failed")
				continue
			}

			for k, v := range coinDetail.DetailPlatforms {
				if k == "solana" {
					platforms = append(platforms, CoingeckoDetailPlatform{
						ChainId: util.ConvertCoingeckoChain(k),
						Address: v.ContractAddress,
						Decimal: int64(v.DecimalPlace),
					})
				}
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
			continue
		}

		platforms := make([]CoingeckoDetailPlatform, 0)

		coinDetail, err, _ := e.svc.CoinGecko.GetCoin(coin.ID)
		if err != nil {
			e.log.Fields(logger.Fields{"coinGeckoId": coin.ID}).Error(err, "[entity.UpsertAllChainTokenData] e.svc.CoinGecko.GetCoin failed")

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

func (e *Entity) formatRouteSwap(req *request.GetSwapRouteRequest, swapRoutes *response.ProviderSwapRoutes) *response.SwapRouteResponse {
	if swapRoutes == nil {
		return &response.SwapRouteResponse{}
	}

	var routeSummary interface{}
	if swapRoutes.Aggregator == "kyber-swap" {
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
			routeSummary = response.RouteSummary{
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
			}
		}
	}

	if swapRoutes.Aggregator == "jupyter" {
		var swapData response.JupiterSwapRoutesSol
		swapByte, _ := json.Marshal(swapRoutes.SwapData)
		err := json.Unmarshal(swapByte, &swapData)
		if err != nil {
			return nil
		}

		newRoute := make([][]response.RouteElement, 0)
		for _, route := range swapData.RoutePlan {
			newRouteElement := make([]response.RouteElement, 0)
			newRouteElement = append(newRouteElement, response.RouteElement{
				Pool:              route.SwapInfo.AmmKey,
				TokenIn:           route.SwapInfo.InputMint,
				TokenOut:          route.SwapInfo.OutputMint,
				LimitReturnAmount: "0",
				SwapAmount:        route.SwapInfo.InAmount,
				AmountOut:         route.SwapInfo.OutAmount,
				// Exchange:       routeEle.Exchange,
				// PoolLength:     routeEle.PoolLength,
				PoolType:       route.SwapInfo.Label,
				TokenOutSymbol: req.To,
			})

			newRoute = append(newRoute, newRouteElement)
		}

		routeSummary = response.RouteSummary{
			TokenIn:                      swapData.InputMint,
			AmountIn:                     swapData.InAmount,
			AmountInUsd:                  "",
			TokenInMarketPriceAvailable:  true,
			TokenOut:                     swapData.OutputMint,
			AmountOut:                    swapData.OutAmount,
			AmountOutUsd:                 "",
			TokenOutMarketPriceAvailable: true,
			Gas:                          "",
			GasPrice:                     "",
			GasUsd:                       "",
			Route:                        newRoute,
		}
	}

	return &response.SwapRouteResponse{
		Code:      swapRoutes.Code,
		Message:   swapRoutes.Message,
		ChainName: swapRoutes.Data.TokenIn.ChainName,
		Data: response.SwapRoute{
			TokenIn:       swapRoutes.Data.TokenIn,
			TokenOut:      swapRoutes.Data.TokenOut,
			RouterAddress: swapRoutes.Data.RouterAddress,
			RouteSummary:  routeSummary,
			Aggregator:    swapRoutes.Aggregator,
			SwapData:      swapRoutes.SwapData,
		},
		Provider: swapRoutes.Aggregator,
	}
}

func (e *Entity) EnrichTokenMochiPay(route *response.SwapRouteResponse) error {
	// code != 1 means fail -> no need enrich mochi pay token
	if route.Code != 1 {
		return nil
	}

	var fromTokenAddress, toTokenAddress string
	if route.Provider != "jupyter" {
		routeSummary := &model.RouteSummary{}
		routeByte, _ := json.Marshal(route.Data.RouteSummary)
		err := json.Unmarshal(routeByte, routeSummary)
		if err != nil {
			return err
		}

		fromTokenAddress, err = util.ConvertToChecksumAddr(routeSummary.TokenIn)
		if err != nil {
			return err
		}

		toTokenAddress, err = util.ConvertToChecksumAddr(routeSummary.TokenOut)
		if err != nil {
			return err
		}
	}

	if route.Provider == "jupyter" {
		quoteResp := &response.JupyterQuoteResponse{}
		quoteByte, _ := json.Marshal(route.Data.RouteSummary)
		err := json.Unmarshal(quoteByte, quoteResp)
		if err != nil {
			return err
		}

		fromTokenAddress = quoteResp.InputMint
		toTokenAddress = quoteResp.OutputMint
	}

	err := e.svc.MochiPay.CreateBatchToken(mochipayrequest.CreateBatchTokenRequest{
		Tokens: []mochipayrequest.CreateTokenRequest{
			{
				Id:          uuid.New().String(),
				Name:        route.Data.TokenIn.Name,
				Symbol:      strings.ToUpper(route.Data.TokenIn.Symbol),
				Decimal:     route.Data.TokenIn.Decimals,
				ChainId:     strconv.Itoa(int(util.ConvertChainNameToChainId(route.ChainName))),
				Address:     fromTokenAddress,
				CoinGeckoId: route.Data.TokenIn.CoingeckoId,
			},
			{
				Id:          uuid.New().String(),
				Name:        route.Data.TokenOut.Name,
				Symbol:      strings.ToUpper(route.Data.TokenOut.Symbol),
				Decimal:     route.Data.TokenOut.Decimals,
				ChainId:     strconv.Itoa(int(util.ConvertChainNameToChainId(route.ChainName))),
				Address:     toTokenAddress,
				CoinGeckoId: route.Data.TokenOut.CoingeckoId,
			},
		},
	})
	if err != nil {
		e.log.Fields(logger.Fields{"route": route}).Error(err, "[mochi-pay.CreateBatchToken] - cannot create batch token")
		return err
	}

	return nil
}

// rule1: from token must be in the list of user balances, for now consider chain too
// rule2: to token must be in the list of our supported chains
func (e *Entity) FilterPossibleToken(profileId string, fromTokens, toTokens []model.Token) (possibleFromTokens, possibleToTokens []model.Token, err error) {
	// get balance from mochi pay
	balances, err := e.svc.MochiPay.GetListBalances(profileId)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId}).Error(err, "[mochi-pay.GetListBalances] - cannot get list balances")
		return nil, nil, err
	}

	// filter from token with user balances
	for _, fromToken := range fromTokens {
		for _, balance := range balances.Data {
			if strings.EqualFold(fromToken.Address, balance.Token.Address) && strconv.Itoa(fromToken.ChainID) == balance.Token.ChainId && strings.EqualFold(fromToken.Symbol, balance.Token.Symbol) {
				possibleFromTokens = append(possibleFromTokens, fromToken)
			}
		}
	}

	// get chain from mochi pay
	chains, err := e.svc.MochiPay.GetListChains()
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId}).Error(err, "[mochi-pay.GetListChains] - cannot get list chains")
		return nil, nil, err
	}

	// filter chain to token
	for _, toToken := range toTokens {
		for _, chain := range chains.Data {
			if strings.EqualFold(strconv.Itoa(toToken.ChainID), chain.ChainId) {
				possibleToTokens = append(possibleToTokens, toToken)
			}
		}
	}

	// if from token not in user balance, return original fromToken
	if len(possibleFromTokens) == 0 {
		possibleFromTokens = fromTokens
	}

	// if to token not in our supported chains, return original toToken
	if len(possibleToTokens) == 0 {
		possibleToTokens = toTokens
	}

	return possibleFromTokens, possibleToTokens, nil
}
