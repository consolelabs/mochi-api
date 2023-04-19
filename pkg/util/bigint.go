package util

import (
	"fmt"
	"math"
	"math/big"
)

func StringToBigInt(s string) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert string to big int")
	}
	return n, nil
}

func BigIntToFloat(v *big.Int, decimal int) (f float64) {
	x := new(big.Float).SetInt(v)
	y := new(big.Float).SetFloat64(math.Pow10(decimal))
	f, _ = new(big.Float).Quo(x, y).Float64()
	return
}
