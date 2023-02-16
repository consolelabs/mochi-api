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

type WalletTransactionAction struct {
	Name           string            `json:"name"`
	Signature      string            `json:"signature"`
	From           string            `json:"from"`
	To             string            `json:"to"`
	Amount         float64           `json:"amount"`
	Unit           string            `json:"unit"`
	NativeTransfer bool              `json:"native_transfer"`
	Contract       *ContractMetadata `json:"contract"`
}

type WalletTransactionData struct {
	ChainID     int       `json:"chain_id"`
	TxHash      string    `json:"tx_hash"`
	ScanBaseUrl string    `json:"scan_base_url"`
	SignedAt    time.Time `json:"signed_at"`
	// From        string                    `json:"from"`
	// To          string                    `json:"to"`
	Actions     []WalletTransactionAction `json:"actions"`
	HasTransfer bool                      `json:"has_transfer"`
	Successful  bool                      `json:"successful"`
}
