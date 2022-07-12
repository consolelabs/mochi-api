package util

import "strings"

func GetURLMarketPlace(marketplace string) (urlMarketPlace string) {
	if strings.ToLower(marketplace) == "opensea" {
		urlMarketPlace = "(https://opensea.io/assets/ethereum/"
	} else if strings.ToLower(marketplace) == "Paintswap" {
		urlMarketPlace = "(https://paintswap.finance/marketplace/assets/"
	} else { // quixotic
		urlMarketPlace = "(https://quixotic.io/asset/"
	}
	return urlMarketPlace
}
