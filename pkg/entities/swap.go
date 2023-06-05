package entities

import (
	"fmt"
	"strconv"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	mochipayrequest "github.com/defipod/mochi/pkg/service/mochipay"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetSwapRoutes(req *request.GetSwapRouteRequest) (*response.SwapRouteResponse, error) {
	// step 1: we need to identify which token user want to swap
	// e.g: user input "usdc" -> we need to parse it into any {token_address, chain_id} possibles

	fromTokens, err := e.getAllChainToken(req.From)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[GetSwapRoutes.getAllChainToken] - cannot get all chain token")
		return nil, err
	}

	toTokens, err := e.getAllChainToken(req.To)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[GetSwapRoutes.getAllChainToken] - cannot get all chain token")
		return nil, err
	}

	amount := util.FloatToString(req.Amount, 18)
	// step 2.0: filter possible route first.
	// rule1: from token must be in the list of user balances, for now consider chain too
	// rule2: to token must be in the list of our supported chains
	fromTokens, toTokens, err = e.FilterPossibleToken(req.ProfileId, fromTokens, toTokens)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[GetSwapRoutes.FilterPossibleToken] - cannot filter possible token")
		return nil, err
	}

	// step 2.1: now we have 2 set of tokens, we need to find the route
	routes, err := e.svc.Swap.GetAllRoutes(fromTokens, toTokens, amount)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[swap.GetAllRoutes] - cannot get all route")
		return nil, err
	}

	// // step 3: we identiy which route is best for user
	r, err := e.svc.Swap.GetBestRoute(routes)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[GetSwapRoutes.getBestRoute] - cannot get best route")
		return nil, err
	}

	toRoute := e.formatRouteSwap(req, r)

	// step 4: enrich data token in mochi pay
	err = e.EnrichTokenMochiPay(toRoute)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.EnrichTokenMochiPay] - cannot enrich token mochi pay")
		return toRoute, nil
	}

	return toRoute, nil

}

func (e *Entity) Swap(req request.SwapRequest) (interface{}, error) {
	// get profile
	profile, err := e.svc.MochiProfile.GetByDiscordID(req.UserDiscordId, true)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-profile.GetByDiscordID] - cannot get profile")
		return nil, err
	}
	chainId := util.ConvertChainNameToChainId(req.ChainName)

	// hash swap address to compare with db
	fromTokenAddress, err := util.ConvertToChecksumAddr(req.RouteSummary.TokenIn)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[util.ConvertToChecksumAddr] - cannot convert to checksum address")
		return nil, err
	}

	toTokenAddress, err := util.ConvertToChecksumAddr(req.RouteSummary.TokenOut)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[util.ConvertToChecksumAddr] - cannot convert to checksum address")
		return nil, err
	}

	// get token from mochi pay
	fromToken, err := e.svc.MochiPay.GetTokenByProperties(mochipayrequest.TokenProperties{
		ChainId: strconv.Itoa(int(chainId)),
		Address: fromTokenAddress,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.svc.MochiPay.GetToken] - cannot get from token")
		return nil, err
	}

	toToken, err := e.svc.MochiPay.GetTokenByProperties(mochipayrequest.TokenProperties{
		ChainId: strconv.Itoa(int(chainId)),
		Address: toTokenAddress,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.svc.MochiPay.GetToken] - cannot get to token")
		return nil, err
	}

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

	// compare balance and amountIn from swap route
	amountSwap, err := util.StringToBigInt(req.RouteSummary.AmountIn)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[util.StringToBigInt] - cannot convert string to big int")
		return nil, err
	}

	bal, err := util.StringToBigInt(balance.Data[0].Amount)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[util.StringToBigInt] - cannot convert string to big int")
		return nil, err
	}

	if amountSwap.Cmp(bal) == 1 {
		return nil, fmt.Errorf("insufficient balance")
	}

	userPublicKey := e.cfg.CentralizedWalletAddress
	if chainId == 999 {
		userPublicKey = e.solana.GetCentralizedWalletAddress()
	}
	// build route
	buildRouteResp, err := e.svc.Swap.BuildSwapRoutes(req.ChainName, &request.BuildSwapRouteRequest{
		Recipient:         userPublicKey,
		Sender:            userPublicKey,
		Source:            consts.ClientID,
		SkipSimulateTx:    false,
		SlippageTolerance: 500,
		RouteSummary:      req.RouteSummary,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[GetSwapRoutes.BuildSwapRoutes] - cannot build swap routes")
		return nil, err
	}

	// send payload to mochi-pay
	err = e.svc.MochiPay.SwapMochiPay(request.MochiPaySwapRequest{
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
		Aggregator:    req.Aggregator,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[mochi-pay.SwapMochiPay] - cannot swap mochi pay")
		return nil, err
	}
	return nil, nil
}
