package util

func ConvertChainToChainId(chain string) string {
	switch chain {
	case "eth":
		return "1"
	case "ftm":
		return "250"
	case "op":
		return "10"
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
