package productdata

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
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

// ProductBotCommand     godoc
// @Summary     Get product bot commands
// @Description Get product bot commands
// @Tags        ProductMetadata
// @Accept      json
// @Produce     json
// @Param       req   query  request.ProductBotCommandRequest true  "request"
// @Success     200 {object} response.ProductBotCommand
// @Router      /product-metadata/commands [get]
func (h *Handler) ProductBotCommand(c *gin.Context) {
	req := request.ProductBotCommandRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.ProductBotCommand] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	commands, err := h.entities.ProductBotCommand(req)
	if err != nil {
		h.log.Error(err, "[handler.ProductBotCommand] entities.ProductBotCommand() failed")
		c.JSON(http.StatusInternalServerError, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse[any](commands, nil, nil, nil))
}

// ProductChangelogs     godoc
// @Summary     Get product changelogs
// @Description Get product changelogs
// @Tags        ProductMetadata
// @Accept      json
// @Produce     json
// @Param       req   query  request.ProductChangelogsRequest false  "request"
// @Success     200 {object} response.ProductChangelogs
// @Router      /product-metadata/changelogs [get]
func (h *Handler) ProductChangelogs(c *gin.Context) {
	req := request.ProductChangelogsRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.ProductChangelogs] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	productChangelogs, total, err := h.entities.ProductChangelogs(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagination := response.PaginationResponse{
		Pagination: model.Pagination{
			Page: req.Page,
			Size: req.Size,
		},
		Total: total,
	}

	resp := response.ProductChangelogs{
		Data:       productChangelogs,
		Pagination: pagination,
	}

	c.JSON(http.StatusOK, resp)
}

// CreateProductChangelogsView   godoc
// @Summary     Created product changelogs viewed
// @Description Created product changelogs viewed
// @Tags        ProductMetadata
// @Accept      json
// @Produce     json
// @Param       req   body  request.CreateProductChangelogsViewRequest true  "create product changelogs viewed request"
// @Success     200 {object} response.CreateProductChangelogsViewed
// @Router      /product-metadata/changelogs/view [post]
func (h *Handler) CreateProductChangelogsView(c *gin.Context) {
	req := request.CreateProductChangelogsViewRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err, "[handler.CreateProductChangelogsView] BindJSON() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	productChangelogsView, err := h.entities.CreateProductChangelogsView(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(productChangelogsView, nil, nil, nil))
}

// GetProductChangelogsView     godoc
// @Summary     Get product changelogs viewed
// @Description Get product changelogs viewed
// @Tags        ProductMetadata
// @Accept      json
// @Produce     json
// @Param       req   query  request.GetProductChangelogsViewRequest  false  "get product changelogs viewed request"
// @Success     200 {object} response.GetProductChangelogsViewed
// @Router      /product-metadata/changelogs/view [get]
func (h *Handler) GetProductChangelogsView(c *gin.Context) {
	req := request.GetProductChangelogsViewRequest{}
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetProductChangelogsView] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	productChangelogsViews, err := h.entities.GetProductChangelogsView(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(productChangelogsViews, nil, nil, nil))
}

func (h *Handler) CrawlChangelogs(c *gin.Context) {
	go h.entities.CrawlChangelogs()
	c.JSON(http.StatusOK, response.CreateResponse(map[string]string{"message": "ok"}, nil, nil, nil))
}

func (h *Handler) GetProductHashtag(c *gin.Context) {
	req := request.GetProductHashtagRequest{}
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetProductHashtag] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	productHashtag, err := h.entities.GetProductHashtag(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(productHashtag, nil, nil, nil))
}

func (h *Handler) GetProductTheme(c *gin.Context) {
	req := request.GetProductThemeRequest{}
	if err := c.BindQuery(&req); err != nil {
		h.log.Error(err, "[handler.GetProductTheme] BindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	productTheme, err := h.entities.GetProductTheme(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CreateResponse(productTheme, nil, nil, nil))
}
