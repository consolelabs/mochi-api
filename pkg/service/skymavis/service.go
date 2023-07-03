package skymavis

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetAddressFarming(address string) (*response.WalletFarmingResponse, error)
	GetOwnedNfts(address string) (*response.AxieMarketNftResponse, error)
	GetInternalTxnsByHash(hash string) (*response.SkymavisTransactionsResponse, error)
}
