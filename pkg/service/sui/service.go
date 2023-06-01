package sui

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetBalance(address string) (*response.SuiAllBalance, error)
	GetCoinMetadata(coinType string) (*response.SuiCoinMetadata, error)
	GetAddressAssets(address string) ([]response.WalletAssetData, error)
	GetTransactionBlock(address string) (*response.SuiTransactionBlock, error)
	GetAddressTxn(address string) ([]response.WalletTransactionData, error)
}
