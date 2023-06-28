package wallet

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	baseerr "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
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

// ListOwnedWallets     			godoc
// @Summary     Get user's wallets
// @Description Get user's wallets
// @Tags        Wallet
// @Accept      json
// @Produce     json
// @Param       guild_id   query  string true  "guild ID"
// @Param       req   body  request.GetTrackingWalletsRequest true  "req"
// @Success     200 {object} response.GetTrackingWalletsResponse
// @Router      /users/:id/wallets [get]
func (h *Handler) ListOwnedWallets(c *gin.Context) {
	var req request.GetTrackingWalletsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.ListOwned] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.GuildID = c.Query("guild_id")
	if req.GuildID == "" {
		err := errors.New("guild_id is required")
		h.log.Error(err, "[handler.Wallet.ListOwned] not enough query params")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.IsOwner = true
	items, err := h.entities.GetTrackingWallets(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListOwned] entity.GetTrackingWallets() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(items, nil, nil, nil))
}

// ListTrackingWallets godoc
// @Summary     Get user's tracking wallets
// @Description Get user's tracking wallets
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       id   			path  string true  "user ID"
// @Success     200 {object} response.GetTrackingWalletsResponse
// @Router      /users/{id}/watchlists/wallets [get]
func (h *Handler) ListTrackingWallets(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		err := errors.New("user ID is required")
		h.log.Error(err, "[handler.Wallet.ListTracking] not enough path params")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
	}

	req := request.GetTrackingWalletsRequest{
		WalletBaseRequest: request.WalletBaseRequest{
			UserID: userID,
		},
	}

	items, err := h.entities.GetTrackingWallets(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListTracking] entity.GetTrackingWallets() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(items, nil, nil, nil))
}

// GetOne     			godoc
// @Summary     Find one user's trackng wallet
// @Description Find one user's trackng wallet
// @Tags        Wallet
// @Accept      json
// @Produce     json
// @Param       id   path  string true  "user Id"
// @Param       query   path  string true  "alias or address query"
// @Success     200 {object} response.GetOneWalletResponse
// @Router      /users/:id/wallets/:query [get]
func (h *Handler) GetOne(c *gin.Context) {
	var req request.GetOneWalletRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.GetOne] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	items, err := h.entities.GetOneWallet(req)
	if err != nil {
		code := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			code = http.StatusNotFound
		}
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.GetOne] entity.GetOneWallet() failed")
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(items, nil, nil, nil))
}

// UpdateTrackingInfo		godoc
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
func (h *Handler) UpdateTrackingInfo(c *gin.Context) {
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

	wallet, err := h.entities.UpdateTrackingInfo(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.UpdateTrackingInfo] entity.UpdateTrackingInfo() failed")
		c.JSON(baseerr.GetStatusCode(err), response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(wallet, nil, nil, nil))
}

// Track     	godoc
// @Summary     Track new wallet
// @Description Track new wallet
// @Tags        WatchList
// @Accept      json
// @Produce     json
// @Param       id   			path  string true  "user ID"
// @Param       request  	body 	request.TrackWalletRequest true "req"
// @Success     200 {object} 		response.ResponseMessage
// @Router      /users/{id}/watchlists/wallets/track [post]
func (h *Handler) Track(c *gin.Context) {
	req, err := extractTrackRequestFromCtx(c)
	if err != nil {
		h.log.Error(err, "[handler.Wallet.Track] extractTrackRequestFromCtx() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	mod, err := req.RequestToUserWalletWatchlistItemModel()
	if err != nil {
		h.log.Error(err, "[handler.Wallet.Track] RequestToUserWalletWatchlistItemModel() failed")
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

	req.UserID = userID

	return req, nil
}

// Untrack     	godoc
// @Summary     Untrack a wallet
// @Description Untrack a wallet
// @Tags        Wallet
// @Accept      json
// @Produce     json
// @Param       id   path  string true  "user ID"
// @Param       req   query request.UntrackWalletRequest true  "req"
// @Success     200 {object} response.ResponseMessage
// @Router      /users/:id/wallets [post]
func (h *Handler) Untrack(c *gin.Context) {
	var req request.UntrackWalletRequest
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
	err := h.entities.UntrackWallet(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.Untrack] entity.UntrackWallet() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
}

func (h *Handler) ListAssets(c *gin.Context) {
	var req request.ListWalletAssetsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.ListAssets] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	items, pnl, latestSnapshotBal, err := h.entities.ListWalletAssets(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListWalletAssets() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// farming data
	farmingData, err := h.entities.ListEthWalletFarming(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListEthWalletFarming() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// staking data
	stakingData, err := h.entities.ListEthWalletStaking(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListEthWalletStaking() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// nft data
	nftData, err := h.entities.ListEthWalletNfts(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListEthWalletNfts() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	data := response.ListAsset{
		Balance:           items,
		Pnl:               pnl,
		LatestSnapshotBal: latestSnapshotBal,
		Farming:           farmingData,
		Staking:           stakingData,
		Nfts:              nftData,
	}

	c.JSON(http.StatusOK, response.CreateResponse(data, nil, nil, nil))
}

func (h *Handler) ListTransactions(c *gin.Context) {
	var req request.ListWalletTransactionsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.ListTransactions] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	items, err := h.entities.ListWalletTxns(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListTransactions] entity.ListWalletTxns() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(items, nil, nil, nil))
}

func (h *Handler) GenerateWalletVerification(c *gin.Context) {
	var uriReq request.WalletBaseRequest
	if err := c.ShouldBindUri(&uriReq); err != nil {
		h.log.Error(err, "[handler.Wallet.GenerateWalletVerification] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var req request.GenerateWalletVerificationRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.GenerateWalletVerification] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	req.UserID = uriReq.UserID

	code, err := h.entities.GenerateWalletVerification(req)
	if err != nil {
		h.log.Error(err, "[handler.Wallet.GenerateWalletVerification] entity.GenerateWalletVerification() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.GenerateVerificationResponse{Code: code}, nil, nil, nil))
}
