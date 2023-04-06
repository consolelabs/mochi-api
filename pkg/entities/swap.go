package entities

import (
	"strconv"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetSwapRoutes(req *request.GetSwapRouteRequest) (*response.KyberSwapRoutes, error) {
	// get from token
	fromToken, err := e.repo.KyberswapSupportedToken.GetByTokenChain(req.From, int64(req.ChainId), req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get from token")
		return nil, err
	}
	// get to token
	toToken, err := e.repo.KyberswapSupportedToken.GetByTokenChain(req.To, int64(req.ChainId), req.ChainName)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[repo.GetByTokenChain] - cannot get to token")
		return nil, err
	}

	// convert string float to string big int
	amount, _ := strconv.ParseFloat(req.Amount, 64)
	stringAmount := util.FloatToBigInt(amount, int64(fromToken.Decimals)).String()

	swapRoutes, err := e.svc.Kyber.GetSwapRoutes(fromToken.ChainName, fromToken.Address, toToken.Address, stringAmount)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[kyber.GetSwapRoutes] - cannot get swap routes")
		return nil, err
	}

	swapRoutes.Data.TokenIn = *fromToken
	swapRoutes.Data.TokenOut = *toToken
	return swapRoutes, nil
}
