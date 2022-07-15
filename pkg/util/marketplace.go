package util

import (
	"math/big"
	"strings"
)

func GetURLMarketPlace(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://opensea.io/assets/"
	case "paintswap":
		return "https://paintswap.finance/marketplace/collections/"
	case "quixotic":
		return "https://quixotic.io/asset/"
	default:
		return ""
	}
}

func GetGainEmoji(gain *big.Float) string {
	cmp := gain.Cmp(big.NewFloat(0))
	if cmp == 1 {
		return "<a:increase:997355555431665736>"
	} else {
		return "<a:decrease:997355537551327363>"
	}
}

func GetChangePnl(pnl *big.Float) string {
	cmp := pnl.Cmp(big.NewFloat(0))
	if cmp == 1 {
		return "+"
	} else {
		return "-"
	}
}
