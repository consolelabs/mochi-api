package emojis

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
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
	codesQuery := c.Query("codes")
	var codes []string

	if codesQuery != "" {
		codes = strings.Split(codesQuery, ",")

		for index, c := range codes {
			codes[index] = strings.ToUpper(c)
		}
	}

	emojis, err := h.entities.GetListEmojis(codes)
	if err != nil {
		h.log.Error(err, "[handler.ListEmojis] - failed to get list emojis")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CreateResponse(emojis, nil, nil, nil))
}
