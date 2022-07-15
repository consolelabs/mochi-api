package util

import (
	"fmt"
	"math/big"
	"strings"
)

func FormatCryptoPrice(price big.Float) string {
	priceFloat64, _ := price.Float64()
	cmp := price.Cmp(big.NewFloat(1))
	if cmp == 1 || cmp == 0 {
		formatPrice := fmt.Sprintf("%.2f", priceFloat64)
		if string(strings.Split(formatPrice, ".")[1]) == "00" {
			return strings.Split(formatPrice, ".00")[0]
		}
		return formatPrice
	} else {
		return fmt.Sprintf("%.4f", priceFloat64)
	}
}
