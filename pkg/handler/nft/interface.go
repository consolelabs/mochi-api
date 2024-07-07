package nft

import (
	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type IHandler interface {
	GetNFTDetail(c *gin.Context)
	GetNFTActivity(c *gin.Context)
	CreateNFTCollection(c *gin.Context)
	handleCreateSolanaCollection(c *gin.Context, req request.CreateNFTCollectionRequest) *model.NFTCollection
	handleCreateEVMCollection(c *gin.Context, req request.CreateNFTCollectionRequest) *model.NFTCollection
	GetSupportedChains(c *gin.Context)
	ListAllNFTCollections(c *gin.Context)
	GetNFTCollectionTickers(c *gin.Context)
	GetNFTTradingVolume(c *gin.Context)
	GetNFTCollections(c *gin.Context)
	GetNFTTokens(c *gin.Context)
	GetDetailNftCollection(c *gin.Context)
	GetCollectionCount(c *gin.Context)
	GetNewListedNFTCollection(c *gin.Context)
	GetNftMetadataAttrIcon(c *gin.Context)
	GetNFTCollectionByAddressChain(c *gin.Context)
	UpdateNFTCollection(c *gin.Context)
	GetGuildDefaultNftTicker(c *gin.Context)
	SetGuildDefaultNftTicker(c *gin.Context)
	GetSuggestionNFTCollections(c *gin.Context)
	GetNftTokenTickers(c *gin.Context)
	GetNftSalesHandler(c *gin.Context)
	GetProfileNFTBalances(c *gin.Context)
}
