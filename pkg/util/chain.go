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
		return "9996"
	case "base":
		return "8453"
	case "zksync":
		return "324"
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
	case "9996":
		return "sui"
	case "8453":
		return "base"
	case "zksync":
		return "324"
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
		"9996":       "sui",
		"999":        "sol",
		"7979":       "dos",
		"0xa86a":     "avax",
		"42161":      "arb",
		"1313161554": "aurora",
		"paintswap":  "ftm",
		"opensea":    "eth",
		"quixotic":   "op",
		"8453":       "base",
		"324":        "zksync",
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

// TODO(trkhoi): enrich table chains in database
func ConvertChainIdToChainName(chainId int64) string {
	switch chainId {
	case 1:
		return "ethereum"
	case 250:
		return "fantom"
	case 56:
		return "bsc"
	case 137:
		return "polygon"
	case 43114:
		return "avalanche"
	case 42161:
		return "arbitrum"
	case 10:
		return "optimism"
	case 199:
		return "bttc"
	case 42262:
		return "oasis"
	case 25:
		return "cronos"
	case 106:
		return "velas"
	case 1313161554:
		return "aurora"
	case 999:
		return "solana"
	case 8453:
		return "base"
	case 324:
		return "zksync"
	default:
		return "ethereum"
	}
}

func ConvertChainNameToChainId(chainName string) int64 {
	switch chainName {
	case "ethereum":
		return 1
	case "fantom":
		return 250
	case "bsc":
		return 56
	case "polygon":
		return 137
	case "avalanche":
		return 43114
	case "arbitrum":
		return 42161
	case "optimism":
		return 10
	case "bttc":
		return 199
	case "oasis":
		return 42262
	case "cronos":
		return 25
	case "velas":
		return 106
	case "aurora":
		return 1313161554
	case "solana":
		return 999
	case "base":
		return 8453
	case "zksync":
		return 324
	default:
		return 1
	}
}

// TODO: not used anymore
func ConvertCoingeckoChain(chainName string) int64 {
	switch chainName {
	case "ethereum":
		return 1
	case "fantom":
		return 250
	case "binance-smart-chain":
		return 56
	case "polygon-pos":
		return 137
	case "avalanche":
		return 43114
	case "arbitrum-one":
		return 42161
	case "optimistic-ethereum":
		return 10
	case "bittorrent":
		return 199
	case "oasis":
		return 42262
	case "cronos":
		return 25
	case "velas":
		return 106
	case "aurora":
		return 1313161554
	case "solana":
		return 999
	case "base":
		return 8453
	case "zksync":
		return 324
	default:
		return 0
	}
}

func ConvertChainCoingecko(chain string) string {
	switch chain {
	case "binance-smart-chain":
		return "bsc"
	case "polygon-pos":
		return "polygon"
	default:
		return ""
	}
}
