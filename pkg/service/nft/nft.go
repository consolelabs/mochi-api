package nft

import (
	"strconv"

	"github.com/defipod/mochi/pkg/response"
)

type NFT struct {
	getNFTURL string
}

func NewService() Service {
	return &NFT{
		getNFTURL: "3rd party API url",
	}
}

// Mock data, not yet implemented
func (n *NFT) GetNFTIndexer(limit int) ([]response.IndexerNFTResponse, error) {
	nftList := []response.IndexerNFTResponse{}
	for i := 1; i <= limit; i++ {
		volume := response.TradeVolume{
			BuyNumber:  11 * i,
			SellNumber: 11*i + 2,
		}
		resp := response.IndexerNFTResponse{
			CollectionAddress: "0xea5da2215038169ccf60dc5be188c43d9982f39" + strconv.Itoa(i),
			CollectionName:    "Cyber Neko " + strconv.Itoa(i),
			Symbol:            "Neko",
			ChainID:           "250",
			TradingVolume:     volume,
		}
		nftList = append(nftList, resp)
	}
	return nftList, nil
}
