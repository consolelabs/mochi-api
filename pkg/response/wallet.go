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
	Nfts              []WalletNftData     `json:"nfts"`
}

type WalletStakingData struct {
	TokenName string  `json:"token_name"`
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"amount"`
	Reward    float64 `json:"reward"`
	Price     float64 `json:"price"`
}

type WalletAssetData struct {
	ChainID        int                            `json:"chain_id"`
	ContractName   string                         `json:"contract_name"`
	ContractSymbol string                         `json:"contract_symbol"`
	AssetBalance   float64                        `json:"asset_balance"`
	UsdBalance     float64                        `json:"usd_balance"`
	Token          AssetToken                     `json:"token"`
	Amount         string                         `json:"amount"`
	DetailStaking  *BinanceStakingProductPosition `json:"detail_staking"`
	DetailLending  *BinancePositionAmountVos      `json:"detail_lending"`
}

type AssetToken struct {
	Id      string          `json:"id"`
	Address string          `json:"address"`
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

type GetBinanceAsset struct {
	Asset      []WalletAssetData                    `json:"asset"`
	Earn       []WalletAssetData                    `json:"earn"`
	SimpleEarn WalletBinanceAssetSimpleEarnResponse `json:"simple_earn"`
}

type WalletBinanceAssetSimpleEarnResponse struct {
	TotalAmountInBTC          string `json:"total_amount_in_btc"`
	TotalAmountInUSDT         string `json:"total_amount_in_usdt"`
	TotalFlexibleAmountInBTC  string `json:"total_flexible_amount_in_btc"`
	TotalFlexibleAmountInUSDT string `json:"total_flexible_amount_in_usdt"`
	TotalLockedInBTC          string `json:"total_locked_in_btc"`
	TotalLockedInUSDT         string `json:"total_locked_in_usdt"`
	BtcPrice                  string `json:"btc_price"`
}

// sky mavis
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

type WalletNftResponse struct {
	Data []WalletNftData `json:"data"`
}

type WalletNftData struct {
	Total          int64               `json:"total"`
	CollectionName string              `json:"collection_name"`
	Tokens         []WalletNftMetadata `json:"tokens"`
}

type WalletNftMetadata struct {
	Image          string `json:"image"`
	MarketplaceURL string `json:"marketplace_url"`
	TokenName      string `json:"token_name"`
}

// axie nft
type AxieMarketNftResponse struct {
	Data *AxieNftListData `json:"data"`
}

type AxieNftListData struct {
	Axies      *AxieNftResult      `json:"axies"`
	Equipments *EquipmentNftResult `json:"equipments"`
	Lands      *LandNftResult      `json:"lands"`
	Items      *LandItemNftResult  `json:"items"`
}

type AxieNftResult struct {
	Total   int64          `json:"total"`
	Results []AxieMetadata `json:"results"`
}

type AxieMetadata struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
	// Owner          string `json:"owner"`
	// Level          int    `json:"level"`
	// MinPrice       string `json:"minPrice"`
	// TokenAddress   string `json:"tokenAddress"`
	// MarketplaceURL string `json:"marketplace_url"`
	// TokenName      string `json:"token_name"`
}

type EquipmentNftResult struct {
	Total   int64               `json:"total"`
	Results []EquipmentMetadata `json:"results"`
}

type EquipmentMetadata struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
	// Total          int      `json:"total"`
	// MinPrice       string   `json:"minPrice"`
	// Collections    []string `json:"collections"`
	// Rarity         string   `json:"rarity"`
	// Image          string   `json:"image"`
	// MarketplaceURL string   `json:"marketplace_url"`
	// TokenName      string   `json:"token_name"`
}

type LandItemMetadata struct {
	FigureURL string `json:"figureURL"`
	Name      string `json:"name"`
	ItemID    int    `json:"itemId"`
	Alias     string `json:"itemAlias"`
	// TokenID        string `json:"tokenId"`
	// MinPrice       string `json:"minPrice"`
	// Rarity         string `json:"rarity"`
	// Image          string `json:"image"`
	// MarketplaceURL string `json:"marketplace_url"`
	// TokenName      string `json:"token_name"`
}

type LandItemNftResult struct {
	Total   int64              `json:"total"`
	Results []LandItemMetadata `json:"results"`
}

type LandMetadata struct {
	LandType string `json:"landType"`
	Col      int    `json:"col"`
	Row      int    `json:"row"`
	// TokenID        string `json:"tokenId"`
	// MinPrice       string `json:"minPrice"`
	// Image          string `json:"image"`
	// MarketplaceURL string `json:"marketplace_url"`
	// TokenName      string `json:"token_name"`
}

type LandNftResult struct {
	Total   int64          `json:"total"`
	Results []LandMetadata `json:"results"`
}

type SkymavisTransactionsResponse struct {
	Results []SkymavisTransactionsResultItem `json:"results"`
	Total   int64                            `json:"total"`
}

type SkymavisTransactionsResultItem struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Value  string `json:"value"`
	TxType string `json:"tx_type"`
	// TxHash      string `json:"tx_hash"`
	// Hash        string `json:"hash"`
	// Input       string `json:"input"`
	// Index       int    `json:"index"`
	// Success     bool   `json:"success"`
	// BlockNumber int    `json:"block_number"`
	// BlockHash   string `json:"block_hash"`
	// Timestamp   int    `json:"timestamp"`
}
