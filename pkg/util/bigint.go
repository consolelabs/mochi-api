package util

import (
	"errors"
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

// FloatToString convert float to big int string with given decimal
// Ignore negative float
// example: FloatToString("0.000000000000000001", 18) => "1"
func FloatToString(s string, decimal int64) string {
	bigval, ok := new(big.Float).SetString(s)
	zero := new(big.Float).SetFloat64(0)
	if !ok || bigval.Cmp(zero) == -1 {
		return "0"
	}

	d := new(big.Float).SetFloat64(math.Pow10(int(decimal)))
	bigval.Mul(bigval, d)

	r := new(big.Int)
	bigval.Int(r) // store converted number in r
	return r.String()
}

// Cmp compare x and y and returns:
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
func Cmp(x, y string) (int, error) {
	n1, ok1 := new(big.Int).SetString(x, 10)
	n2, ok2 := new(big.Int).SetString(y, 10)
	if !ok1 || !ok2 {
		return 0, errors.New("invalid x or y")
	}
	return n1.Cmp(n2), nil
}

func CmpBigInt(x, y *big.Int) (int, error) {
	return x.Cmp(y), nil
}

func CalculateTokenBalance(bigNumber string, decimal int) (bal float64) {
	balance, ok := new(big.Float).SetString(bigNumber)
	if !ok {
		return
	}
	parsedBal, _ := balance.Float64()
	bal = parsedBal / math.Pow10(decimal)
	return
}
