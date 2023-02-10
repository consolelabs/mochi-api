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

// List     			godoc
// @Summary     Get user's trackng wallets
// @Description Get user's trackng wallets
// @Tags        Wallet
// @Accept      json
// @Produce     json
// @Param       req   path  request.GetTrackingWalletsRequest true  "req"
// @Success     200 {object} response.GetTrackingWalletsResponse
// @Router      /users/:id/wallets [get]
func (h *Handler) List(c *gin.Context) {
	var req request.GetTrackingWalletsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.List] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	items, err := h.entities.GetTrackingWallets(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.List] entity.GetTrackingWallets() failed")
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
// @Param       req   path  request.GetOneWalletRequest true  "req"
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

// Track     	godoc
// @Summary     Track new wallet
// @Description Track new wallet
// @Tags        Wallet
// @Accept      json
// @Produce     json
// @Param       id   			path  string true  "user ID"
// @Param       request  	body 	request.TrackWalletRequest true "req"
// @Success     200 {object} 		response.ResponseMessage
// @Router      /users/:id/wallets [post]
func (h *Handler) Track(c *gin.Context) {
	var req request.TrackWalletRequest
	if err := c.BindJSON(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.Track] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	err := h.entities.TrackWallet(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.Track] entity.TrackWallet() failed")
		code := http.StatusInternalServerError
		if err == baseerr.ErrConflict {
			code = http.StatusConflict
		}
		c.JSON(code, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(response.ResponseMessage{Message: "OK"}, nil, nil, nil))
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
	items, err := h.entities.ListWalletAssets(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListWalletAssets() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(items, nil, nil, nil))
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
