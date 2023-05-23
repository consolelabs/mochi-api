package nft

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type error interface {
	Error() string
}

type Handler struct {
	entities *entities.Entity
	log      logger.Logger
}

func New(entities *entities.Entity, logger logger.Logger) IHandler {
	return &Handler{
		entities: entities,
		log:      logger,
	}
}

// GetNFTDetail     godoc
// @Summary     Get NFT Detail
// @Description Get NFT Detail
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Param       id   path  string true  "Token ID"
// @Param       guild_id   query  string false  "Guild ID"
// @Param       query_address   query  bool true  "Query by using Collection Address"
// @Success     200 {object} response.IndexerGetNFTTokenDetailResponseWithSuggestions
// @Router      /nfts/{symbol}/{id} [get]
func (h *Handler) GetNFTDetail(c *gin.Context) {
	symbol := c.Param("symbol")
	tokenID := c.Param("id")
	query := struct {
		GuildID        string `json:"guild_id" form:"guild_id"`
		QueryByAddress *bool  `json:"query_address" form:"query_address" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&query); err != nil {
		h.log.Fields(logger.Fields{"query": query}).Error(err, "[handler.GetNFTDetail] - failed to bind query")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.GetNFTDetail(symbol, tokenID, query.GuildID, *query.QueryByAddress)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": symbol, "id": tokenID}).Error(err, "[handler.GetNFTDetail] - failed to get NFt detail")
		c.JSON(baseerrs.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetNFTActivity     godoc
// @Summary     Get NFT Activity
// @Description Get NFT Activity
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Collection address | Symbol"
// @Param       id   path  string true  "Token ID"
// @Param       page   query  string false  "Page"
// @Param       size   query  string false  "Size"
// @Success     200 {object} response.GetNFTActivityResponse
// @Router      /nfts/{symbol}/{id}/activity [get]
func (h *Handler) GetNFTActivity(c *gin.Context) {
	collectionAddress := c.Param("symbol")
	tokenID := c.Param("id")
	res, err := h.entities.GetNFTActivity(collectionAddress, tokenID, c.Request.URL.RawQuery)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.log.Fields(logger.Fields{"collection_address": collectionAddress, "token_id": tokenID}).Info("[handler.GetNFTActivity] - record not found")
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Fields(logger.Fields{"collection_address": collectionAddress, "id": tokenID}).Error(err, "[handler.GetNFTActivity] - failed to get NFT activity")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var data *model.NFTCollection
	switch req.ChainID {
	case "sol":
		data = h.handleCreateSolanaCollection(c, req)
	case "apt":
		data = h.handleCreateBluemoveCollection(c, req)
	case "sui":
		data = h.handleCreateBluemoveCollection(c, req)
	default:
		data = h.handleCreateEVMCollection(c, req)
	}

	if data == nil {
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

func (h *Handler) handleCreateSolanaCollection(c *gin.Context, req request.CreateNFTCollectionRequest) *model.NFTCollection {
	data, err := h.entities.CreateSolanaNFTCollection(req)
	if err != nil {
		if strings.Contains(err.Error(), "Already added") || strings.Contains(err.Error(), "does not have") {
			h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Info("[handler.CreateNFTCollection] - duplicated record")
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return nil
		}
		h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Error(err, "[handler.CreateNFTCollection] - failed to create NFT collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return nil
	}
	return data
}

func (h *Handler) handleCreateEVMCollection(c *gin.Context, req request.CreateNFTCollectionRequest) *model.NFTCollection {
	data, err := h.entities.CreateEVMNFTCollection(req)
	if err != nil {
		if strings.Contains(err.Error(), "Already added") || strings.Contains(err.Error(), "does not have") {
			h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Info("[handler.CreateNFTCollection] - duplicated record")
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return nil
		}
		h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Error(err, "[handler.CreateNFTCollection] - failed to create NFT collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return nil
	}
	return data
}

func (h *Handler) handleCreateBluemoveCollection(c *gin.Context, req request.CreateNFTCollectionRequest) *model.NFTCollection {
	data, err := h.entities.CreateBluemoveNFTCollection(req)
	if err != nil {
		if strings.Contains(err.Error(), "Already added") || strings.Contains(err.Error(), "does not have") {
			h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Info("[handler.CreateNFTCollection] - duplicated record")
			c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
			return nil
		}
		h.log.Fields(logger.Fields{"address": req.Address, "chain": req.Chain, "chainID": req.ChainID, "author": req.Author}).Error(err, "[handler.CreateNFTCollection] - failed to create NFT collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return nil
	}
	return data
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
	c.JSON(http.StatusOK, response.CreateResponse([]string{"eth", "heco", "bsc", "matic", "op", "btt", "okt", "movr", "celo", "metis", "cro", "xdai", "boba", "ftm", "avax", "arb", "aurora"}, nil, nil, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(nfts, nil, nil, nil))

}

// GetNFTCollectionTickers     godoc
// @Summary     Get NFT collection tickers
// @Description Get NFT collection tickers
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       collection_address   query  string true  "CollectionAddress"
// @Param       from   query  string true  "from"
// @Param       to   query  string true  "to"
// @Success     200 {object} response.IndexerNFTCollectionTickersResponse
// @Router      /nfts/collections/tickers [get]
func (h *Handler) GetNFTCollectionTickers(c *gin.Context) {
	var req request.GetNFTCollectionTickersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetNFTCollectionTickers] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.GetNFTCollectionTickers(req, c.Request.URL.RawQuery)
	if err != nil {
		if err.Error() == "record not found" {
			h.log.Infof("[indexer.GetNFTCollectionTickers] Indexer does not have ticker for this collection. Req: %s", req)
			c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Fields(logger.Fields{"req": req, "query": c.Request.URL.RawQuery}).Error(err, "[handler.GetNFTCollectionTickers] - failed to get NFT collection ticker")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(nfts, nil, nil, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("symbol is required"), nil))
		return
	}

	data, err := h.entities.GetNFTTokens(symbol, c.Request.URL.RawQuery)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": symbol}).Error(err, "[handler.GetNFTTokens] - failed to get NFT tokens")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetDetailNftCollection     godoc
// @Summary     Get detail nft collection
// @Description Get detail nft collection
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   path  string true  "Symbol"
// @Param       query_address   query  bool true  "Query by collection address"
// @Success     200 {object} response.GetDetailNftCollectionResponse
// @Router      /nfts/collections/{symbol}/detail [get]
func (h *Handler) GetDetailNftCollection(c *gin.Context) {
	collectionSymbol := c.Param("symbol")
	if collectionSymbol == "" {
		h.log.Info("[handler.GetDetailNftCollection] - symbol empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("symbol is required"), nil))
		return
	}
	query := struct {
		QueryByAddress *bool `json:"query_address" form:"query_address" binding:"required"`
	}{}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	collection, err := h.entities.GetDetailNftCollection(collectionSymbol, *query.QueryByAddress)
	if err != nil {
		h.log.Fields(logger.Fields{"symbol": collectionSymbol}).Error(err, "[handler.GetDetailNftCollection] - failed to get detail NFT collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(collection, nil, nil, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, errors.New("cannot get collections count"), nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
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
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, errors.New("cannot get icons"), nil))
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
// @Success     200 {object} response.GetNFTCollectionByAddressChainResponse
// @Router      /nfts/collections/address/{address} [get]
func (h *Handler) GetNFTCollectionByAddressChain(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		h.log.Info("[handler.GetNFTCollectionByAddressChain] - address empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("address is required"), nil))
		return
	}
	inputChain := c.Query("chain")
	if inputChain == "" {
		h.log.Info("[handler.GetNFTCollectionByAddressChain] - input chain empty")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("chain is required"), nil))
		return
	}
	chainId := util.ConvertInputToChainId(inputChain)

	data, err := h.entities.GetNFTCollectionByAddressChain(address, chainId)
	if err != nil {
		h.log.Fields(logger.Fields{"address": address}).Error(err, "[handler.GetNFTCollectionByAddressChain] - failed to get NFT Collection by Address and chain")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
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
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("address is required"), nil))
		return
	}
	err := h.entities.UpdateNFTCollection(address)
	if err != nil {
		h.log.Fields(logger.Fields{"address": address}).Error(err, "[handler.UpdateNFTCollection] - failed to update collection")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}

// AddNftWatchlist     godoc
// @Summary     Add to user's nft watchlist
// @Description Add to user's nft watchlist
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       req body request.AddNftWatchlistRequest true "request"
// @Success     200 {object} response.NftWatchlistSuggestResponse
// @Router      /nfts/watchlist [post]
func (h *Handler) AddNftWatchlist(c *gin.Context) {
	var req request.AddNftWatchlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.AddNftWatchlist] - failed to bind request")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("invalid request"), nil))
		return
	}

	res, err := h.entities.AddNftWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.AddNftWatchlist] - failed to add watchlist")
		c.JSON(baseerrs.GetStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetNftWatchlist     godoc
// @Summary     Get user's nft watchlist
// @Description Get user's nft watchlist
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       user_id   query  string true  "user_id"
// @Param       page   query  string true  "page"
// @Param       size   query  string true  "size"
// @Success     200 {object} response.GetNftWatchlistResponse
// @Router      /nfts/watchlist [get]
func (h *Handler) GetNftWatchlist(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	req := request.GetNftWatchlistRequest{
		UserID: c.Query("user_id"),
		Page:   page,
		Size:   size,
	}

	data, err := h.entities.GetNftWatchlist(&req)
	if err != nil {
		h.log.Error(err, "[handler.GetNftWatchlist] - failed to get watchlist")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, data)
}

// DeleteNftWatchlist     godoc
// @Summary     Remove from user's nft watchlist
// @Description Remove from user's nft watchlist
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       symbol   query  string true  "symbol"
// @Param       user_id   query  string true  "user_id"
// @Success     200 {object} object
// @Router      /nfts/watchlist [delete]
func (h *Handler) DeleteNftWatchlist(c *gin.Context) {
	req := request.DeleteNftWatchlistRequest{
		Symbol: c.Query("symbol"),
		UserID: c.Query("user_id"),
	}

	err := h.entities.DeleteNftWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.DeleteNftWatchlist] - failed to delete watchlist")
		code := http.StatusInternalServerError
		if err == baseerrs.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nil})
}

// GetGuildDefaultNftTicker     godoc
// @Summary     Get guild default nft ticker
// @Description Get guild default nft ticker
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "Guild ID"
// @Param       query   query  string true  "Guild ticker query"
// @Success     200 {object} response.GetGuildDefaultNftTickerResponse
// @Router      /nfts/default-nft-ticker [get]
func (h *Handler) GetGuildDefaultNftTicker(c *gin.Context) {
	var req request.GetGuildDefaultNftTickerRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetGuildDefaultNftTicker] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.GetGuildDefaultNftTicker(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetGuildDefaultNftTicker] entity.GetGuildDefaultNftTicker() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// SetGuildDefaultNftTicker     godoc
// @Summary     Set guild default nft ticker
// @Description Set guild default nft ticker
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       Request  body request.GuildConfigDefaultNftTickerRequest true "Set guild default ticker request"
// @Success     200 {object} response.ResponseDataMessage
// @Router      /nfts/default-nft-ticker [post]
func (h *Handler) SetGuildDefaultNftTicker(c *gin.Context) {
	req := request.GuildConfigDefaultNftTickerRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.SetGuildDefaultNftTicker] c.ShouldBindJSON failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := h.entities.SetGuildDefaultNftTicker(req); err != nil {
		h.log.Error(err, "[handler.SetGuildDefaultNFtTicker] entity.SetGuildDefaultNftTicker failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetSuggestionNFTCollections     godoc
// @Summary     Get guild suggest nft collections
// @Description Get guild suggest nft collections
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       query   query  string true  "symbol collection query"
// @Success     200 {object} response.GetSuggestionNFTCollectionsResponse
// @Router      /nfts/collections/suggestion [get]
func (h *Handler) GetSuggestionNFTCollections(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		h.log.Info("[handler.GetSuggestionNFTCollections] query is required")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("missing query"), nil))
		return
	}

	h.log.Infof("DEBUG PROD: [handler.GetSuggestionNFTCollections] Checking req: ", query)

	collections, err := h.entities.GetSuggestionNftCollections(query)
	if err != nil {
		h.log.Fields(logger.Fields{"query": query}).Error(err, "[handler.GetSuggestionNFTCollections] entities.GetSuggestionNftCollections() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(collections, nil, nil, nil))
}

// GetNftTokenTickers     godoc
// @Summary     Get NFT token tickers
// @Description Get NFT token tickers
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       collection_address   query  string true  "CollectionAddress"
// @Param       token_id   query  string true  "Token ID"
// @Param       from   query  string true  "from"
// @Param       to   query  string true  "to"
// @Success     200 {object} response.IndexerGetNFTTokenTickersResponse
// @Router      /nfts/tickers [get]
func (h *Handler) GetNftTokenTickers(c *gin.Context) {
	var req request.GetNFTTokenTickersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.GetNftTokenTickers] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.GetNFTTokenTickers(req, c.Request.URL.RawQuery)
	if err != nil {
		if err.Error() == "record not found" {
			h.log.Infof("[entities.GetNFTTokenTickers] Indexer does not have ticker for this token. Req: %s", req)
			c.JSON(http.StatusNotFound, response.CreateResponse[any](nil, nil, err, nil))
			return
		}
		h.log.Fields(logger.Fields{"req": req, "query": c.Request.URL.RawQuery}).Error(err, "[entities.GetNFTTokenTickers] - failed to get NFT token ticker")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// GetNftsalesHandler     godoc
// @Summary     Get NFT sales
// @Description Get NFT sales
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       collection-address   query  string true  "Collection address"
// @Param       platform   query  string true  "Platform"
// @Success     200 {object} response.NftSalesResponse
// @Router      /nfts/sales [get]
func (h *Handler) GetNftSalesHandler(c *gin.Context) {
	addr := c.Query("collection-address")
	platform := c.Query("platform")
	data, err := h.entities.GetNftSales(addr, platform)
	if err != nil || data == nil {
		h.log.Fields(logger.Fields{"address": addr, "platform": platform}).Error(err, "[handler.GetNftSalesHandler] - failed to get NFT sales")
		c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, errors.New("collection not found"), nil))
		return
	}
	c.JSON(http.StatusOK, data)
}

// CreateTradeOffer     godoc
// @Summary     Create Trade Offer
// @Description Create Trade Offer
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       Request  body request.CreateTradeOfferRequest true "Create Trade Offer Request"
// @Success     200 {object} response.CreateTradeOfferResponse
// @Router      /nfts/trades [post]
func (h *Handler) CreateTradeOffer(c *gin.Context) {
	req := request.CreateTradeOfferRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateTradeOffer] c.ShouldBindJSON failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	data, err := h.entities.CreateTradeOffer(req)
	if err != nil {
		h.log.Error(err, "[handler.CreateTradeOffer] entities.CreateTradeOffer failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse(data, nil, nil, nil))
}

// GetTradeOffer     godoc
// @Summary     Get Trade Offer
// @Description Get Trade Offer
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       id path  string true  "Trade Offer ID"
// @Success     200 {object} response.GetTradeOfferResponse
// @Router      /nfts/trades/{id} [get]
func (h *Handler) GetTradeOffer(c *gin.Context) {
	tradeId := c.Param("id")
	if tradeId == "" {
		h.log.Info("[handler.GetTradeOffer] - trade id missing")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, errors.New("id is required"), nil))
		return
	}
	data, err := h.entities.GetTradeOffer(tradeId)
	if err != nil {
		h.log.Error(err, "[handler.GetTradeOffer] entities.GetTradeOffer failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// this used to get data from indexer and put in new table
// TODO(trkhoi): put this in background job instead of api
func (h *Handler) EnrichSoulboundNFT(c *gin.Context) {
	collectionAddress := c.Query("collection_address")
	err := h.entities.EnrichSoulboundNFT(collectionAddress)
	if err != nil {
		h.log.Error(err, "[handler.EnrichSoulboundNFT] entities.EnrichSoulboundNFT failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// GetSoulboundNFT     godoc
// @Summary     Get Nft Soulbound
// @Description Get Nft Soulbound
// @Tags        NFT
// @Accept      json
// @Produce     json
// @Param       collection_address query  string true  "collection address"
// @Success     200 {object} response.GetSoulBoundNFTResponse
// @Router      /nfts/soulbound [get]
func (h *Handler) GetSoulboundNFT(c *gin.Context) {
	collectionAddress := c.Query("collection_address")
	res, err := h.entities.GetSoulboundNFT(collectionAddress)
	if err != nil {
		h.log.Error(err, "[handler.GetSoulboundNFT] entities.GetSoulboundNFT failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse(res, nil, nil, nil))
}
