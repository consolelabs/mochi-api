package util

import "strings"

func ConvertChainToChainId(chain string) string {
	switch chain {
	case "sol":
		return "999"
	case "eth":
		return "1"
	case "ftm":
		return "250"
	case "op":
		return "10"
	case "bsc":
		return "56"
	case "paintswap":
		return "250"
	case "opensea":
		return "1"
	case "quixotic":
		return "10"
	case "apt":
		return "9999"
	case "arb":
		return "42161"
	case "polygon":
		return "137"
	case "okc":
		return "66"
	case "onus":
		return "1975"
	case "sui":
		return "9997"
	default:
		return chain
	}
}

func ConvertChainIDToChain(chain string) string {
	switch chain {
	case "999":
		return "sol"
	case "1":
		return "eth"
	case "250":
		return "ftm"
	case "10":
		return "op"
	case "42161":
		return "arb"
	case "66":
		return "okc"
	case "1975":
		return "onus"
	case "sol":
		return "sol"
	case "137":
		return "polygon"
	case "9999":
		return "apt"
	case "9997":
		return "sui"
	default:
		return chain
	}
}

func ConvertMarkplaceToChainId(marketplace string) int64 {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return 1
	case "quixotic":
		return 10
	case "paintswap":
		return 250
	case "looksrare":
		return 1
	case "x2y2":
		return 1
	case "nftkey":
		return 250
	default:
		return 250
	}
}

func ConvertInputToChainId(input string) string {
	mapChainIdChain := map[string]string{
		"1":          "eth",
		"128":        "heco",
		"56":         "bsc",
		"137":        "matic",
		"10":         "op",
		"199":        "btt",
		"66":         "okc",
		"1285":       "movr",
		"42220":      "celo",
		"1088":       "metis",
		"25":         "cro",
		"0x64":       "xdai",
		"288":        "boba",
		"250":        "ftm",
		"1975":       "onus",
		"9999":       "apt",
		"9997":       "sui",
		"999":        "sol",
		"0xa86a":     "avax",
		"42161":      "arb",
		"1313161554": "aurora",
		"paintswap":  "ftm",
		"opensea":    "eth",
		"quixotic":   "op",
	}
	chainId := ""
	if _, exist := mapChainIdChain[strings.ToLower(input)]; exist {
		chainId = input
	}

	for k, v := range mapChainIdChain {
		if v == strings.ToLower(input) {
			chainId = k
		}
	}
	return ConvertChainToChainId(chainId)
}
