package util

import (
	"math/big"
	"strings"
)

func GetURLMarketPlace(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://opensea.io/assets/ethereum/"
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
		return "<a:increase:997330373560250379>"
	} else {
		return "<a:decrease:997330345089302588>"
	}
}
