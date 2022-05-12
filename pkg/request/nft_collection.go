package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type CreateNFTCollectionRequest struct {
	Address string `json:"address"`
	Chain   string `json:"chain"`
	ChainID string `json:"-"`
}

func (input *CreateNFTCollectionRequest) Bind(c *gin.Context) error {
	err := c.BindJSON(input)
	if err != nil {
		return err
	}
	if input.Address == "" {
		err = fmt.Errorf("address is required")
		return err
	}
	if input.Chain == "" {
		err = fmt.Errorf("chain is required")
		return err
	}
	mapChainChainId := map[string]string{
		"eth":    "1",
		"heco":   "128",
		"bsc":    "56",
		"matic":  "137",
		"op":     "10",
		"btt":    "199",
		"okt":    "66",
		"movr":   "1285",
		"celo":   "42220",
		"metis":  "1088",
		"cro":    "25",
		"xdai":   "0x64",
		"boba":   "288",
		"ftm":    "250",
		"avax":   "0xa86a",
		"arb":    "42161",
		"aurora": "1313161554",
	}
	if c, exist := mapChainChainId[strings.ToLower(input.Chain)]; exist {
		input.ChainID = c
	}
	if input.ChainID == "" {
		err = fmt.Errorf("chain is not supported/invalid")
		return err
	}
	return err
}
