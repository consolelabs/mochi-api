package util

import (
	"fmt"
	"math/big"
	"strings"
)

func GetURLMarketPlace(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://opensea.io/collection/"
	case "paintswap":
		return "https://paintswap.finance/marketplace/collections/"
	case "quixotic":
		return "https://quixotic.io/asset/"
	default:
		return ""
	}
}

func GetTransactionUrl(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://etherscan.io/tx/"
	case "paintswap":
		return "https://ftmscan.com/tx/"
	case "quixotic":
		return "https://optimistic.etherscan.io/tx/"
	case "looksrare":
		return "https://etherscan.io/tx/"
	case "nftkey":
		return "https://ftmscan.com/tx/"
	case "x2y2":
		return "https://etherscan.io/tx/"
	default:
		return ""
	}
}

func GetWalletUrl(marketplace string) (urlMarketPlace string) {
	switch strings.ToLower(marketplace) {
	case "opensea":
		return "https://etherscan.io/address/"
	case "paintswap":
		return "https://ftmscan.com/address/"
	case "quixotic":
		return "https://optimistic.etherscan.io/address/"
	case "looksrare":
		return "https://etherscan.io/address/"
	case "nftkey":
		return "https://ftmscan.com/address/"
	case "x2y2":
		return "https://etherscan.io/address/"
	default:
		return ""
	}
}

func GetCollectionExplorerUrl(address string, chainID string) string {
	switch chainID {
	case "1":
		return "https://etherscan.io/address/" + address
	case "250":
		return "https://ftmscan.com/address/" + address
	case "10":
		return "https://optimistic.etherscan.io/address/" + address
	case "56":
		return "https://bscscan.com/address/" + address
	default:
		return ""
	}
}

func GetGainEmoji(gain *big.Float) string {
	cmp := gain.Cmp(big.NewFloat(0))
	if cmp == 1 {
		return "<a:increase:997355555431665736>"
	} else {
		return "<a:decrease:997355537551327363>"
	}
}

func GetChangePnl(pnl *big.Float) string {
	cmp := pnl.Cmp(big.NewFloat(0))
	if cmp == 1 {
		return "+"
	} else {
		return "-"
	}
}

func GetTokenMarketplaceUrl(collectionAddress, collectionSymbol, tokenID, platformName string) (tokenMarketplaceUrl string) {
	switch strings.ToLower(platformName) {
	case "opensea":
		return fmt.Sprintf("https://opensea.io/assets/ethereum/%s/%s", collectionAddress, tokenID)
	case "paintswap":
		return fmt.Sprintf("https://paintswap.finance/marketplace/assets/%s/%s", collectionAddress, tokenID)
	case "quixotic":
		return fmt.Sprintf("https://qx.app/asset/%s/%s", collectionAddress, tokenID)
	case "nftkey":
		return fmt.Sprintf("https://nftkey.app/collections/bitkillasavax/token-details/?tokenId=%s", tokenID)
	default:
		return ""
	}
}
