package handler

import (
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/logger"
	_ "github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"github.com/gin-gonic/gin"
)

type error interface {
	Error() string
}

// GetNFTDetail     godoc
// @Summary     Get NFT Detail
// @Description Get NFT Detail
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Param       id   path  string true  "Token ID"
// @Param       guild_id   query  string true  "Guild ID"
// @Success     200 {object} response.IndexerGetNFTTokenDetailResponseWithSuggestions
// @Router      /nfts/{symbol}/{id} [get]
func (h *Handler) GetNFTDetail(c *gin.Context) {
	symbol := c.Param("symbol")
	tokenID := c.Param("id")
	guildID := c.Query("guild_id")
	// to prevent error when query db
	if guildID == "" {
		guildID = "0"
	}

	res, err := h.entities.GetNFTDetail(symbol, tokenID, guildID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.log.Fields(logger.Fields{"symbol": symbol, "token_id": tokenID, "guild_id": guildID}).Info("[handler.GetNFTDetail] - record not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		h.log.Fields(logger.Fields{"symbol": symbol, "id": tokenID}).Error(err, "[handler.GetNFTDetail] - failed to get NFt detail")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res.Data.Image = util.StandardizeUri(res.Data.Image)

	c.JSON(http.StatusOK, res)
}

// TODO: add test for this api
// CreateNFTCollection     godoc
// @Summary     Create NFT Collection
// @Description Create NFT Collection
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateNFTCollectionRequest true "Create nft collection request"
// @Success     200 {object} response.CreateNFTCollectionResponse
// @Router      /nfts/collections [post]
func (h *Handler) CreateNFTCollection(c *gin.Context) {
	var req request.CreateNFTCollectionRequest

	if err := c.BindJSON(&req); err != nil {
		h.log.Fields(logger.Fields{"body": req}).Error(err, "[handler.CreateNFTCollection] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.entities.CreateNFTCollection(req)
	if err != nil {
		if strings.Contains(err.Error(), "Already added") || strings.Contains(err.Error(), "does not have") {
			h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Info("[handler.CreateNFTCollection] - duplicated record")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Error(err, "[handler.CreateNFTCollection] - failed to create NFT collection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateNFTCollectionResponse{Data: data})
}

// GetSupportedChains     godoc
// @Summary     Get supported chains
// @Description Get supported chains
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetSupportedChains
// @Router      /nfts/supported-chains [post]
func (h *Handler) GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, response.GetSupportedChains{Data: []string{"eth", "heco", "bsc", "matic", "op", "btt", "okt", "movr", "celo", "metis", "cro", "xdai", "boba", "ftm", "avax", "arb", "aurora"}})
}

// ListAllNFTCollections     godoc
// @Summary     List all nft collections
// @Description List all nft collections
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Success     200 {object} response.ListAllNFTCollectionsResponse
// @Router      /nfts [get]
func (h *Handler) ListAllNFTCollections(c *gin.Context) {
	nfts, err := h.entities.ListAllNFTCollections()
	if err != nil {
		h.log.Error(err, "[handler.ListAllNFTCollections] - failed to list all NFT collections")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ListAllNFTCollectionsResponse{Data: nfts})
}

// GetNFTCollectionTickers     godoc
// @Summary     Get NFT collection tickers
// @Description Get NFT collection tickers
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Success     200 {object} response.IndexerNFTCollectionTickersResponse
// @Router      /nfts/collections/{symbol}/tickers [get]
func (h *Handler) GetNFTCollectionTickers(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		h.log.Info("[handler.GetNFTCollectionTickers] - symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	res, err := h.entities.GetNFTCollectionTickers(symbol, c.Request.URL.RawQuery)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": symbol, "query": c.Request.URL.RawQuery}).Error(err, "[handler.GetNFTCollectionTickers] - failed to get NFT collection ticker")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetNFTTradingVolume     godoc
// @Summary     Get NFT trading volume
// @Description Get NFT trading volume
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Success     200 {object} response.NFTTradingVolumeResponse
// @Router      /nfts/trading-volume [get]
func (h *Handler) GetNFTTradingVolume(c *gin.Context) {
	nfts, err := h.entities.GetSvc().Indexer.GetNFTTradingVolume()
	if err != nil {
		h.log.Error(err, "[handler.GetNFTTradingVolume] - failed to get NFT trading volume from indexer")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nfts})
}

// GetNFTCollections     godoc
// @Summary     Get NFT trading volume
// @Description Get NFT trading volume
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       page   query  int false  "Page" default(0)
// @Param       size   query  int false  "Size" default(10)
// @Success     200 {object} response.NFTCollectionsResponse
// @Router      /nfts/collections [get]
func (h *Handler) GetNFTCollections(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	if page == "" {
		page = "0"
	}
	if size == "" {
		size = "10"
	}
	data, err := h.entities.GetNFTCollections(page, size)
	if err != nil {
		h.log.Fields(logger.Fields{"page": page, "size": size}).Error(err, "[handler.GetNFTCollections] - failed to get NFT collections")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetNFTTokens     godoc
// @Summary     Get NFT Tokens
// @Description Get NFT Tokens
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Success     200 {object} response.IndexerGetNFTTokensResponse
// @Router      /nfts/collections/{symbol} [get]
func (h *Handler) GetNFTTokens(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		h.log.Info("[handler.GetNFTTokens] - symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	data, err := h.entities.GetNFTTokens(symbol, c.Request.URL.RawQuery)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[handler.GetNFTTokens] - failed to get NFT tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// CreateNFTSalesTracker     godoc
// @Summary     Create NFT Sales tracker
// @Description Create NFT Sales tracker
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Success     200 {object} response.IndexerGetNFTTokensResponse
// @Router      /nfts/collections/{symbol} [get]
func (h *Handler) CreateNFTSalesTracker(c *gin.Context) {
	var req request.NFTSalesTrackerRequest
	if err := c.Bind(&req); err != nil {
		h.log.Fields(logger.Fields{"address": req.ContractAddress, "platform": req.Platform, "guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.CreateNFTSalesTracker] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.entities.CreateSalesTracker(req)
	if err != nil {
		h.log.Fields(logger.Fields{"address": req.ContractAddress, "platform": req.Platform, "guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[handler.CreateNFTSalesTracker] - failed to create sales tracker")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// GetDetailNftCollection     godoc
// @Summary     Get detail nft collection
// @Description Get detail nft collection
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Success     200 {object} response.GetDetailNftCollectionResponse
// @Router      /nfts/collections/{symbol}/detail [get]
func (h *Handler) GetDetailNftCollection(c *gin.Context) {
	collectionSymbol := c.Param("symbol")
	if collectionSymbol == "" {
		h.log.Info("[handler.GetDetailNftCollection] - symbol empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	collection, err := h.entities.GetDetailNftCollection(collectionSymbol)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": collectionSymbol}).Error(err, "[handler.GetDetailNftCollection] - failed to get detail NFT collection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.GetDetailNftCollectionResponse{Data: collection})
}

// GetAllNFTSalesTracker     godoc
// @Summary     Get all nft sales tracker
// @Description Get all nft sales tracker
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetAllNFTSalesTrackerResponse
// @Router      /nfts/sales-tracker [get]
func (h *Handler) GetAllNFTSalesTracker(c *gin.Context) {
	guildID := c.Query("guild_id")
	if guildID != "" {
		data, err := h.entities.GetNFTSaleSTrackerByGuildID(guildID)
		if err != nil {
			h.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[handler.GetNFTSaleSTrackerByGuildID] - failed to get nft sales tracker")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}
	data, err := h.entities.GetAllNFTSalesTracker()
	if err != nil {
		h.log.Error(err, "[handler.GetAllNFTSalesTracker] - failed to get all NFT sales tracker")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get info"})
		return
	}

	c.JSON(http.StatusOK, response.GetAllNFTSalesTrackerResponse{Data: data})
}

// DeleteNFTSalesTracker     godoc
// @Summary     Delete NFT sales tracker
// @Description Delete NFT sales tracker
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Param       contract_address   query  string true  "Contract Address"
// @Success     200 {object} response.ResponseMessage
// @Router      /nfts/sales-tracker [delete]
func (h *Handler) DeleteNFTSalesTracker(c *gin.Context) {
	guildID := c.Query("guild_id")
	contractAddress := c.Query("contract_address")
	if guildID == "" {
		h.log.Info("[handler.DeleteNFTSalesTracker] - guild id empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "guild id is required"})
		return
	}
	if contractAddress == "" {
		h.log.Info("[handler.DeleteNFTSalesTracker] - contract address empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "contract address is required"})
		return
	}
	err := h.entities.DeleteNFTSalesTracker(guildID, contractAddress)
	if err != nil {
		h.log.Fields(logger.Fields{"guildID": guildID, "contractAddress": contractAddress}).Error(err, "[handler.DeleteDefaultRoleByGuildID] - failed to delete default role config")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseMessage{Message: "OK"})

}

// GetCollectionCount     godoc
// @Summary     Get collection count
// @Description Get collection count
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetCollectionCountResponse
// @Router      /nfts/collections/stats [get]
func (h *Handler) GetCollectionCount(c *gin.Context) {
	data, err := h.entities.GetCollectionCount()
	if err != nil {
		h.log.Error(err, "[handler.GetCollectionCount] - failed to get collections count")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get collections count"})
		return
	}

	c.JSON(http.StatusOK, response.GetCollectionCountResponse{Data: data})
}

// GetNewListedNFTCollection     godoc
// @Summary     Get new listed nft collection
// @Description Get new listed nft collection
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       page   query  int false  "Page" default(0)
// @Param       size   query  int false  "Size" default(10)
// @Param       interval   query  int false  "Interval" default(7)
// @Success     200 {object} response.NFTNewListedResponse
// @Router      /nfts/new-listed [get]
func (h *Handler) GetNewListedNFTCollection(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	interval := c.Query("interval")
	if interval == "" {
		interval = "7"
	}
	if page == "" {
		page = "0"
	}
	if size == "" {
		size = "10"
	}

	data, err := h.entities.GetNewListedNFTCollection(interval, page, size)
	if err != nil {
		h.log.Fields(logger.Fields{"page": page, "size": size, "interval": interval}).Error(err, "[handler.GetNewListedNFTCollection] - failed to get new listed NFT collection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetNftMetadataAttrIcon     godoc
// @Summary     Get NFT metadata attribute icon
// @Description Get NFT metadata attribute icon
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Success     200 {object} response.NftMetadataAttrIconResponse
// @Router      /nfts/icons [get]
func (h *Handler) GetNftMetadataAttrIcon(c *gin.Context) {
	data, err := h.entities.GetNftMetadataAttrIcon()
	if err != nil {
		h.log.Error(err, "[handler.GetNftMetadataAttrIcon] - failed to get NFT metadata icons")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get icons"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetNFTCollectionByAddressChain     godoc
// @Summary     Get NFT collection by address chain
// @Description Get NFT collection by address chain
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       address   path  string true  "Collection Address"
// @Param       chain   query  string true  "Chain"
// @Success     200 {object} model.NFTCollection
// @Router      /nfts/collections/address/{address} [get]
func (h *Handler) GetNFTCollectionByAddressChain(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		h.log.Info("[handler.GetNFTCollectionByAddressChain] - address empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "address is required"})
		return
	}
	inputChain := c.Query("chain")
	if inputChain == "" {
		h.log.Info("[handler.GetNFTCollectionByAddressChain] - input chain empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "chain is required"})
		return
	}
	chainId := util.ConvertInputToChainId(inputChain)

	data, err := h.entities.GetNFTCollectionByAddressChain(address, chainId)
	if err != nil {
		h.log.Fields(logger.Fields{"address": address}).Error(err, "[handler.GetNFTCollectionByAddressChain] - failed to get NFT Collection by Address and chain")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// UpdateNFTCollection     godoc
// @Summary     Update NFT Collection
// @Description Update NFT Collection
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       address   path  string true  "Collection Address"
// @Success     200 {object} response.ResponseMessage
// @Router      /nfts/collections/{address} [patch]
func (h *Handler) UpdateNFTCollection(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		h.log.Info("[handler.GetNFTCollectionByAddressChain] - address empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "address is required"})
		return
	}
	err := h.entities.UpdateNFTCollection(address)
	if err != nil {
		h.log.Fields(logger.Fields{"address": address}).Error(err, "[handler.UpdateNFTCollection] - failed to update collection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}
