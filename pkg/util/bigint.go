package util

import (
	"fmt"
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

func BigIntToFloat(bigInt *big.Int, decimal float64) (intVal float64, fracVal big.Accuracy) {
	floatVal := new(big.Float).SetInt(bigInt)
	intVal, fracVal = new(big.Float).Quo(floatVal, big.NewFloat(decimal)).Float64()
	return
}
