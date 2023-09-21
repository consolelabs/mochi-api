package emojis

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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

// ListEmojis   godoc
// @Summary     list emojis
// @Description list emojis
// @Tags        Emojis
// @Accept      json
// @Produce     json
// @Param       codes query    string true "codes"
// @Success     200   {object} response.ListEmojisResponse
// @Router      /emojis [get]
func (h *Handler) ListEmojis(c *gin.Context) {
	req := request.GetListEmojiRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		h.log.Error(err, "[handler.ListEmojis] ShouldBindQuery() failed")
		c.JSON(http.StatusBadRequest, response.CreateResponse[any](nil, nil, err, nil))
		return
	}

	var codes []string

	if req.Codes != "" {
		codes = strings.Split(req.Codes, ",")

		for index, c := range codes {
			codes[index] = strings.ToUpper(c)
		}
	}
	req.ListCode = codes

	if req.Codes == "" {
		req.IsQueryAll = true
	}

	if req.Size <= 0 {
		req.Size = 10
	}

	emojis, total, err := h.entities.GetListEmojis(req)
	if err != nil {
		h.log.Error(err, "[handler.ListEmojis] - failed to get list emojis")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": emojis,
		"pagination": gin.H{
			"total": total,
			"page":  req.Page,
			"size":  req.Size,
		},
	})
}
