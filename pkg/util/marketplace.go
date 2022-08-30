package util

import (
	"math/big"
	"strings"
)

func GetURLMarketPlace(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://opensea.io/collection/"
	case "paintswap":
		return "https://paintswap.finance/marketplace/collections/"
	case "quixotic":
		return "https://quixotic.io/asset/"
	default:
		return ""
	}
}

func GetTransactionUrl(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://etherscan.io/tx/"
	case "paintswap":
		return "https://ftmscan.com/tx/"
	case "quixotic":
		return "https://optimistic.etherscan.io/tx/"
	case "looksrare":
		return "https://etherscan.io/tx/"
	case "nftkey":
		return "https://ftmscan.com/tx/"
	case "x2y2":
		return "https://etherscan.io/tx/"
	default:
		return ""
	}
}

func GetWalletUrl(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://etherscan.io/address/"
	case "paintswap":
		return "https://ftmscan.com/address/"
	case "quixotic":
		return "https://optimistic.etherscan.io/address/"
	case "looksrare":
		return "https://etherscan.io/address/"
	case "nftkey":
		return "https://ftmscan.com/address/"
	case "x2y2":
		return "https://etherscan.io/address/"
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
