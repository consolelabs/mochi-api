package entities

import (
	"fmt"
	"strconv"

	"github.com/k0kubun/pp"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetSwapRoutes(req *request.GetSwapRouteRequest) (*response.KyberSwapRoutes, error) {
	amount, _ := strconv.ParseFloat(req.Amount, 64)
	stringAmount := fmt.Sprintf("%.0f", amount*1000000000000000000)

	swapRoutes, err := e.svc.Kyber.GetSwapRoutes(stringAmount)
	if err != nil {
		return nil, err
	}
	pp.Println("check data")
	pp.Println(swapRoutes)

	return swapRoutes, nil
}
