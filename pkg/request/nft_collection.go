package request

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateNFTCollectionRequest struct {
	Address   string `json:"address"`
	Chain     string `json:"chain"`
	ChainID   string `json:"chain_id"`
	Author    string `json:"author"`
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
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
	if input.ChainID == "" {
		err = fmt.Errorf("chain is required")
		return err
	}
	mapChainChainId := map[string]string{
		"1":          "eth",
		"128":        "heco",
		"56":         "bsc",
		"137":        "matic",
		"10":         "op",
		"199":        "btt",
		"66":         "okt",
		"1285":       "movr",
		"42220":      "celo",
		"1088":       "metis",
		"25":         "cro",
		"0x64":       "xdai",
		"288":        "boba",
		"250":        "ftm",
		"0xa86a":     "avax",
		"42161":      "arb",
		"1313161554": "aurora",
		"paintswap":  "ftm",
		"opensea":    "eth",
		"quixotic":   "op",
	}
	if c, exist := mapChainChainId[strings.ToLower(input.ChainID)]; exist {
		input.Chain = c
	}
	// handle usecase req chainId = "eth"
	for _, v := range mapChainChainId {
		if v == input.ChainID {
			input.Chain = v
		}
	}
	if input.Chain == "" {
		err = fmt.Errorf("chain is not supported/invalid")
		return err
	}
	return err
}
