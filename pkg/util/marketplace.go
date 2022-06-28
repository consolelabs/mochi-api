package util

import "strings"

func convertPaintswapToFtmAddress(paintswapMarketplace string) string {
	splittedPaintswap := strings.Split(paintswapMarketplace, "/")
	return splittedPaintswap[len(splittedPaintswap)-1]
}

func convertOpenseaToEthAddress(openseaMarketplace string) string {
	return ""
}

func HandleMarketplaceLink(contractAddress, chain string) string {
	switch strings.Contains(contractAddress, "/") {
	case false:
		return contractAddress
	case true:
		switch chain {
		case "paintswap":
			return convertPaintswapToFtmAddress(contractAddress)
		case "opensea":
			return convertOpenseaToEthAddress(contractAddress)
		default:
			return convertPaintswapToFtmAddress(contractAddress)
		}
	default:
		return contractAddress
	}
}
