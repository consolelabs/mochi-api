package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type GetTrackingWalletsResponse struct {
	Data []model.UserWalletWatchlistItem `json:"data"`
}

type GetOneWalletResponse struct {
	Data *model.UserWalletWatchlistItem `json:"data"`
}

type WalletAssetData struct {
	ChainID        int     `json:"chain_id"`
	ContractName   string  `json:"contract_name"`
	ContractSymbol string  `json:"contract_symbol"`
	AssetBalance   float64 `json:"asset_balance"`
	UsdBalance     float64 `json:"usd_balance"`
}

type ContractMetadata struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type WalletTransactionData struct {
	ChainID      int               `json:"chain_id"`
	TxHash       string            `json:"tx_hash"`
	TxBaseUrl    string            `json:"tx_base_url"`
	SignedAt     time.Time         `json:"signed_at"`
	From         string            `json:"from"`
	To           string            `json:"to"`
	NativeSymbol string            `json:"native_symbol"`
	Type         string            `json:"type"`
	Amount       float64           `json:"amount"`
	NftIDs       []string          `json:"nfts"`
	Contract     *ContractMetadata `json:"contract"`
}
