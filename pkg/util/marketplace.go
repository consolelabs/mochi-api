package util

import "strings"

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
