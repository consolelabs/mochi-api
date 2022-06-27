package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNFTDetail(c *gin.Context) {
	symbol := c.Param("symbol")
	tokenID := c.Param("id")

	data, err := h.entities.GetNFTDetail(symbol, tokenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) CreateNFTCollection(c *gin.Context) {
	var req request.CreateNFTCollectionRequest

	if err := req.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checksumAddress, _ := util.ConvertToChecksumAddr(req.Address)

	checkExitsNFT, err := h.entities.CheckExistNftCollection(checksumAddress)
	if err != nil {
		if checkExitsNFT {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if checkExitsNFT {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already added. Nft is in sync progress"})
		return
	}

	req.Address = checksumAddress

	data, err := h.entities.CreateNFTCollection(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) GetSupportedChains(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": []string{"eth", "heco", "bsc", "matic", "op", "btt", "okt", "movr", "celo", "metis", "cro", "xdai", "boba", "ftm", "avax", "arb", "aurora"}})
}

func (h *Handler) ListAllNFTCollections(c *gin.Context) {
	nfts, err := h.entities.ListAllNFTCollections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nfts})
}

func (h *Handler) GetNFTCollectionTickers(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	data, err := h.entities.GetNFTCollection(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
func (h *Handler) GetNFTTradingVolume(c *gin.Context) {
	nfts, err := h.entities.GetSvc().Indexer.GetNFTTradingVolume()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nfts})
}

func (h *Handler) GetNFTCollections(c *gin.Context) {
	data, err := h.entities.GetNFTCollections(c.Request.URL.RawQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) GetNFTTokens(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	data, err := h.entities.GetNFTTokens(symbol, c.Request.URL.RawQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *Handler) CreateNFTSalesTracker(c *gin.Context) {
	var req request.NFTSalesTrackerRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.entities.CreateNFTSalesTracker(req.ContractAddress, req.Platform, req.GuildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})

}

func (h *Handler) GetDetailNftCollection(c *gin.Context) {
	collectionSymbol := c.Query("collectionSymbol")
	collection, err := h.entities.GetDetailNftCollection(collectionSymbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": collection})
}
func (h *Handler) GetAllNFTSalesTracker(c *gin.Context) {
	data, err := h.entities.GetAllNFTSalesTracker()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
