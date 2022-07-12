package util

import "strings"

func GetURLMarketPlace(marketplace string) string {
	var urlMarketPlace = ""
	if strings.ToLower(marketplace) == "opensea" {
		urlMarketPlace = "[" + marketplace + "]" + "(https://opensea.io/assets/ethereum/"
	} else if strings.ToLower(marketplace) == "Paintswap" {
		urlMarketPlace = "[" + marketplace + "]" + "(https://paintswap.finance/marketplace/assets/"
	} else { // quixotic
		urlMarketPlace = "[" + marketplace + "]" + "(https://quixotic.io/asset/"
	}
	return urlMarketPlace
}
