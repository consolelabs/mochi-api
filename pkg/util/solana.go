package util

import (
	"context"
	"math"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/program/token"
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

// If mint authority is diabled, then the contract is considered as "ownership renounced"
func IsProgramContractRenounced(addr string) (bool, error) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	accountInfo, err := c.GetAccountInfo(context.Background(), addr)
	if err != nil {
		return false, err
	}

	mintAcc, err := token.MintAccountFromData(accountInfo.Data)
	if err != nil {
		return false, err
	}

	if mintAcc.MintAuthority == nil {
		return true, nil
	}
	return false, nil
}
