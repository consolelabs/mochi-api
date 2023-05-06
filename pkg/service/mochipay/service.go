package mochipay

import (
	"github.com/defipod/mochi/pkg/request"
)

type Service interface {
	SwapMochiPay(req request.KyberSwapRequest) error
	GetBalance(profileId, token, chainId string) (*GetBalanceDataResponse, error)
	Transfer(req request.MochiPayTransferRequest) error
	CreateToken(req CreateTokenRequest) error
	ListTokens(symbol string) ([]Token, error)
	GetToken(symbol, chainId string) (*Token, error)
}