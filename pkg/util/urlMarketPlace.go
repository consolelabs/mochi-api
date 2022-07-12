package util

func GetURLMarketPlace(marketplace string) string {
	var urlMarketPlace = ""
	if marketplace == "Opensea" {
		urlMarketPlace = "[" + marketplace + "]" + "(https://opensea.io/assets/ethereum/"
	} else if marketplace == "Paintswap" {
		urlMarketPlace = "[" + marketplace + "]" + "(https://paintswap.finance/marketplace/assets/"
	} else { // quixotic
		urlMarketPlace = "[" + marketplace + "]" + "(https://quixotic.io/asset/"
	}
	return urlMarketPlace
}
