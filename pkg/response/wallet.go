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
	Pnl               string              `json:"pnl"`
	LatestSnapshotBal string              `json:"latest_snapshot_bal"`
	Balance           []WalletAssetData   `json:"balance"`
	Farming           []LiquidityPosition `json:"farming"`
	Staking           []WalletStakingData `json:"staking"`
}

type WalletStakingData struct {
	TokenName string  `json:"token_name"`
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"amount"`
	Reward    float64 `json:"reward"`
	Price     float64 `json:"price"`
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

type WalletBinanceAssetResponse struct {
	AssetBalance   string `json:"asset_balance"`
	TotalAmountUsd string `json:"total_amount_usd"`
	Asset          string `json:"asset"`
}

// skymavis
type WalletFarmingResponse struct {
	Data *WalletFarmingData `json:"data"`
}

type WalletFarmingData struct {
	LiquidityPositions []LiquidityPosition `json:"liquidityPositions"`
}

type LiquidityPosition struct {
	ID                    string              `json:"id"`
	LiquidityTokenBalance string              `json:"liquidityTokenBalance"`
	Pair                  PairData            `json:"pair"`
	Reward                WalletFarmingReward `json:"reward"`
}

type WalletFarmingReward struct {
	Amount float64       `json:"amount"`
	Token  PairTokenData `json:"token"`
}

type PairData struct {
	ID          string        `json:"id"`
	TotalSupply string        `json:"totalSupply"`
	ReserveUSD  string        `json:"reserveUSD"`
	Token0Price string        `json:"token0Price"`
	Token1Price string        `json:"token1Price"`
	Token0      PairTokenData `json:"token0"`
	Token1      PairTokenData `json:"token1"`
}

type PairTokenData struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Symbol       string         `json:"symbol"`
	TokenDayData []TokenDayData `json:"tokenDayData"`
	Balance      float64        `json:"balance"`
}

type TokenDayData struct {
	PriceUSD string `json:"priceUSD"`
}
