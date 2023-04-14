package mochipay

import (
	"github.com/defipod/mochi/pkg/request"
)

type Service interface {
	SwapMochiPay(req request.KyberSwapRequest) error
	GetBalance(profileId, token string) (*GetBalanceDataResponse, error)
	Transfer(req request.MochiPayTransferRequest) error
	CreateToken(req CreateTokenRequest) error
}
