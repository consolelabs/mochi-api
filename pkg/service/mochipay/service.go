package mochipay

import (
	"github.com/defipod/mochi/pkg/request"
)

type Service interface {
	SwapMochiPay(req request.MochiPaySwapRequest) error
	GetBalance(profileId, token, chainId string) (*GetBalanceDataResponse, error)
	Transfer(req request.MochiPayTransferRequest) (*TipResponse, error)
	CreateToken(req CreateTokenRequest) error
	ListTokens(symbol string) ([]Token, error)
	GetToken(symbol, chainId string) (*Token, error)
	TransferVaultMochiPay(req request.MochiPayVaultRequest) (*VaultResponse, error)
	CreateBatchToken(req CreateBatchTokenRequest) error
	GetTokenByProperties(req TokenProperties) (*Token, error)
	GetListBalances(profileId string) (*GetBalanceDataResponse, error)
	GetListChains() (*GetChainDataResponse, error)

	// TransferV2
	TransferV2(req TransferV2Request) (*TransferV2Response, error)
	ApplicationTransfer(req ApplicationTransferRequest) (*ApplicationTransferResponse, error)
}
