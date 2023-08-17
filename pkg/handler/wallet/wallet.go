package wallet

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

	if !util.ValidateNumberSeries(req.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.ListOwnedWallets] validate profile id failed")
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

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.ListOwnedWallets] validate profile id failed")
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

func (h *Handler) ListAssets(c *gin.Context) {
	var req request.ListWalletAssetsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.log.Error(err, "[handler.Wallet.ListAssets] ShouldBindUri() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.ListAssets] validate profile id failed")
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
	farmingData, err := h.entities.ListWalletFarmings(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListWalletFarmings() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// staking data
	stakingData, err := h.entities.ListWalletStakings(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListWalletStakings() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	// nft data
	nftData, err := h.entities.ListWalletNfts(req)
	if err != nil {
		h.log.Fields(logger.Fields{"req": req}).Error(err, "[handler.Wallet.ListAssets] entity.ListWalletNfts() failed")
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

	if !util.ValidateNumberSeries(req.WatchlistBaseRequest.ProfileID) {
		err := errors.New("profile Id is invalid")
		h.log.Error(err, "[handler.ListTransactions] validate profile id failed")
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
