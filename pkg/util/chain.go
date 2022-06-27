package util

func ConvertChainToChainId(chain string) string {
	switch chain {
	case "eth":
		return "1"
	case "ftm":
		return "250"
	case "op":
		return "10"
	default:
		return chain
	}
}
