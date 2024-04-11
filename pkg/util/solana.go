package util

import (
	"context"
	"math"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
)

func GetSplTokenSupply(addr string) (float64, error) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	supply, dec, err := c.GetTokenSupply(context.Background(), addr)
	if err != nil {
		return 0, err
	}

	result := float64(supply) / math.Pow10(int(dec))
	return result, nil
}
