package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type GetTrackingWalletsResponse struct {
	Data model.UserWalletWatchlist `json:"data"`
}

type GetOneWalletResponse struct {
	Data *model.UserWalletWatchlistItem `json:"data"`
}

type ListAsset struct {
	Pnl               string            `json:"pnl"`
	LatestSnapshotBal string            `json:"latest_snapshot_bal"`
	Balance           []WalletAssetData `json:"balance"`
}
type WalletAssetData struct {
	ChainID        int        `json:"chain_id"`
	ContractName   string     `json:"contract_name"`
	ContractSymbol string     `json:"contract_symbol"`
	AssetBalance   float64    `json:"asset_balance"`
	UsdBalance     float64    `json:"usd_balance"`
	Token          AssetToken `json:"token"`
	Amount         string     `json:"amount"`
}

type AssetToken struct {
	Name    string          `json:"name"`
	Symbol  string          `json:"symbol"`
	Decimal int64           `json:"decimal"`
	Price   float64         `json:"price"`
	Native  bool            `json:"native"`
	Chain   AssetTokenChain `json:"chain"`
}

type AssetTokenChain struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
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
	ChainID     int                       `json:"chain_id"`
	TxHash      string                    `json:"tx_hash"`
	ScanBaseUrl string                    `json:"scan_base_url"`
	SignedAt    time.Time                 `json:"signed_at"`
	Actions     []WalletTransactionAction `json:"actions"`
	HasTransfer bool                      `json:"has_transfer"`
	Successful  bool                      `json:"successful"`
}

type GenerateWalletVerificationResponseData struct {
	Code string `json:"code"`
}

type WalletBinanceResponse struct {
	TotalBtc float64 `json:"total_btc"`
	Price    float64 `json:"price"`
}
