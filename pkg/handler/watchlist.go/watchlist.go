package watchlist

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	baseerr "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

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

// ListUserTrackingWallets godoc
// @Summary     Get user's tracking wallets
// @Description Get user's tracking wallets
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       id   			path  string true  "profile ID"
// @Success     200 {object} response.GetTrackingWalletsResponse
// @Router      /users/{id}/watchlists/wallets [get]
func (h *Handler) ListUserTrackingWallets(c *gin.Context) {
	var base request.WatchlistBaseRequest
	if err := c.BindUri(&base); err != nil {
		h.log.Error(err, "[handler.ListTrackingWallets] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	req := request.GetTrackingWalletsRequest{
		ProfileID:   base.ProfileID,
		WithBalance: true,
	}

	if !util.ValidateNumberSeries(req.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.ListUserTrackingWallets] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	items, err := h.entities.GetTrackingWallets(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListTracking] entity.GetTrackingWallets() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(items, nil, nil, nil))
}

// UpdateTrackingWalletInfo		godoc
// @Summary     				Update tracked wallet's info
// @Description 				Update tracked wallet's info
// @Tags        				WatchList
// @Accept      				json
// @Produce     				json
// @Param       				id   path  string true  "user Id"
// @Param       				address   path  string true  "address or current alias of tracked wallet"
// @Param       				request body request.UpdateTrackingInfoRequest true "req"
// @Success     				200 {object} response.GetOneWalletResponse
// @Router      				/users/{id}/watchlists/wallets/{address} [put]
func (h *Handler) UpdateTrackingWalletInfo(c *gin.Context) {
	var req request.UpdateTrackingInfoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.UpdateTrackingInfo] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.UpdateTrackingInfo] ShouldBindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.UpdateTrackingWalletInfo] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	wallet, err := h.entities.UpdateTrackingInfo(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.UpdateTrackingInfo] entity.UpdateTrackingInfo() failed")
		c.JSON(baseerr.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(wallet, nil, nil, nil))
}

// TrackWallet  godoc
// @Summary     Track new wallet
// @Description Track new wallet
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       id   			path  string true  "user ID"
// @Param       request  	body 	request.TrackWalletRequest true "req"
// @Success     200 {object} 		response.ResponseMessage
// @Router      /users/{id}/watchlists/wallets/track [post]
func (h *Handler) TrackWallet(c *gin.Context) {
	req, err := extractTrackRequestFromCtx(c)
	if err != nil {
		h.log.Error(err, "[handler.TrackWallet] extractTrackRequestFromCtx() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.TrackWallet] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	mod, err := req.RequestToUserWalletWatchlistItemModel()
	if err != nil {
		h.log.Error(err, "[handler.TrackWallet] RequestToUserWalletWatchlistItemModel() failed")
		c.JSON(baseerr.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err = h.entities.TrackWallet(mod, req.ChannelID, req.MessageID)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.Track] entity.TrackWallet() failed")
		c.JSON(baseerr.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func extractTrackRequestFromCtx(c *gin.Context) (request.TrackWalletRequest, error) {
	var req request.TrackWalletRequest
	if err := c.BindJSON(&req); err != nil {
		return req, err
	}

	userID := c.Param("id")
	if userID == "" {
		return req, errors.New("user ID is required")
	}

	req.ProfileID = userID

	return req, nil
}

// UntrackWallet godoc
// @Summary      Untrack a wallet
// @Description  Untrack a wallet
// @Tags         WatchList
// @Accept       json
// @Produce      json
// @Param        id   path  string true  "user ID"
// @Param        req   query request.UntrackWalletRequest true  "req"
// @Success      200 {object} response.ResponseMessage
// @Router       /users/{id}/watchlists/wallets/untrack [post]
func (h *Handler) UntrackWallet(c *gin.Context) {
	var req request.UntrackWalletRequest
	if err := c.BindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.Untrack] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.Untrack] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if req.Address == "" && req.Alias == "" {
		err := errors.New("either address or alias is required")
		h.log.Fields(logger.Fields{"req": req}).Errorf(err, "[handler.Wallet.Untrack] %s", err.Error())
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.UntrackWallet] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.UntrackWallet(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.Untrack] entity.UntrackWallet() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

// ListTrackingTokens     godoc
// @Summary     Get user's watchlist
// @Description Get user's watchlist
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       query query request.ListTrackingTokensRequest true "query"
// @Param       id    path  string true  "profile ID"
// @Success     200 {object} response.GetWatchlistResponse
// @Router      /users/{id}/watchlists/tokens [get]
func (h *Handler) ListTrackingTokens(c *gin.Context) {
	var base request.WatchlistBaseRequest
	if err := c.BindUri(&base); err != nil {
		h.log.Error(err, "[handler.ListTrackingTokens] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.ListTrackingTokensRequest
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.ListTrackingTokens] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.ProfileID = base.ProfileID

	if !util.ValidateNumberSeries(req.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.ListTrackingTokens] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.GetUserWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.ListTrackingTokens] entity.GetUserWatchlist() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(res, nil, nil, nil))
}

// TrackToken   godoc
// @Summary     Add to user's watchlist
// @Description Add to user's watchlist
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       req body request.AddToWatchlistRequest true "request"
// @Success     200 {object} response.AddToWatchlistResponse
// @Router      /users/{id}/watchlists/tokens/track [post]
func (h *Handler) TrackToken(c *gin.Context) {
	var req request.AddToWatchlistRequest
	if err := c.BindUri(&req); err != nil {
		h.log.Error(err, "[handler.TrackToken] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.TrackToken] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.TrackToken] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.AddToWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.TrackToken] entity.AddToWatchlist() failed")
		c.JSON(baseerr.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, res)
}

// UntrackToken godoc
// @Summary     Remove from user's watchlist
// @Description Remove from user's watchlist
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       req query request.RemoveFromWatchlistRequest true "request"
// @Success     200 {object} object
// @Router      /users/{id}/watchlists/tokens/untrack [post]
func (h *Handler) UntrackToken(c *gin.Context) {
	var req request.RemoveFromWatchlistRequest
	if err := c.BindUri(&req); err != nil {
		h.log.Error(err, "[handler.UntrackToken] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UntrackToken] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.UntrackToken] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.RemoveFromWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.UntrackToken] entity.RemoveFromWatchlist() failed")
		code := http.StatusInternalServerError
		if err == baseerr.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](nil, nil, nil, nil))
}

// TrackNft     godoc
// @Summary     Add to user's nft watchlist
// @Description Add to user's nft watchlist
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       req body request.AddNftWatchlistRequest true "request"
// @Success     200 {object} response.NftWatchlistSuggestResponse
// @Router      /users/{id}/watchlists/nfts/track [post]
func (h *Handler) TrackNft(c *gin.Context) {
	var req request.AddNftWatchlistRequest
	if err := c.BindUri(&req); err != nil {
		h.log.Error(err, "[handler.TrackNft] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.TrackNft] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.TrackNft] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	res, err := h.entities.AddNftWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.TrackNft] - failed to add watchlist")
		c.JSON(baseerr.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, res)
}

// ListTrackingNfts     godoc
// @Summary     Get user's nft watchlist
// @Description Get user's nft watchlist
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       query   query  request.ListTrackingNftsRequest true  "query"
// @Param       id   		path  string true  "profile ID"
// @Success     200 {object} response.GetNftWatchlistResponse
// @Router      /users/{id}/watchlists/nfts [get]
func (h *Handler) ListTrackingNfts(c *gin.Context) {
	var base request.WatchlistBaseRequest
	if err := c.BindUri(&base); err != nil {
		h.log.Error(err, "[handler.ListTrackingNfts] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.ListTrackingNftsRequest
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.ListTrackingNfts] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.ProfileID = base.ProfileID

	if !util.ValidateNumberSeries(req.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.ListTrackingNfts] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data, err := h.entities.GetNftWatchlist(&req)
	if err != nil {
		h.log.Error(err, "[handler.ListTrackingNfts] - failed to get watchlist")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

// UntrackNft     godoc
// @Summary     Remove from user's nft watchlist
// @Description Remove from user's nft watchlist
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       query   query  request.DeleteNftWatchlistRequest true  "symbol"
// @Param       id   path  string true  "profile ID"
// @Success     200 {object} object
// @Router      /users/{id}/watchlists/nfts/untrack [post]
func (h *Handler) UntrackNft(c *gin.Context) {
	var base request.WatchlistBaseRequest
	if err := c.BindUri(&base); err != nil {
		h.log.Error(err, "[handler.UntrackNft] BindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.DeleteNftWatchlistRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.UntrackNft] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.ProfileID = base.ProfileID

	if !util.ValidateNumberSeries(req.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.UntrackNft] validate profile id failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	err := h.entities.DeleteNftWatchlist(req)
	if err != nil {
		h.log.Error(err, "[handler.UntrackNft] - failed to delete watchlist")
		code := http.StatusInternalServerError
		if err == baseerr.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nil})
}

// ListTrackingWallets go doc
// @Summary     Get tracking wallets
// @Description Get tracking wallets
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param 			address 	query  string false  "address"
// @Param       with_balance	 query  bool false  "with balance"
// @Success     200 {object} response.GetTrackingWalletsResponse
// @Router      /watchlists/wallets [get]
func (h *Handler) ListTrackingWallets(c *gin.Context) {
	var (
		req request.GetTrackingWalletsRequest
		err error
	)

	req.WithBalance, err = strconv.ParseBool(c.Query("with_balance"))
	if err != nil {
		// Default false
		req.WithBalance = false
	}

	req.Address = c.Query("address")

	items, err := h.entities.GetTrackingWallets(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListTracking] entity.GetTrackingWallets() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(items, nil, nil, nil))
}
