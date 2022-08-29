package util

import "strings"

func ConvertChainToChainId(chain string) string {
	switch chain {
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
	default:
		return chain
	}
}

func ConvertChainIDToChain(chain string) string {
	switch chain {
	case "1":
		return "eth"
	case "250":
		return "ftm"
	case "10":
		return "op"
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
		"66":         "okt",
		"1285":       "movr",
		"42220":      "celo",
		"1088":       "metis",
		"25":         "cro",
		"0x64":       "xdai",
		"288":        "boba",
		"250":        "ftm",
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
		if v == input {
			chainId = k
		}
	}
	return ConvertChainToChainId(chainId)
}
