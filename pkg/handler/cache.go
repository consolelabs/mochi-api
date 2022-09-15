package handler

import (
	"net/http"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gin-gonic/gin"
)

// SetUpvoteMessageCache     godoc
// @Summary     Set or overwrite an upvote message cache
// @Description Set or overwrite an upvote message cache
// @Tags        Cache
// @Accept      json
// @Produce     json
// @Param       Request  body request.SetUpvoteMessageCacheRequest true "Set upvote message cache request"
// @Success     200 {object} response.ResponseMessage
// @Router      /cache/upvote [post]
func (h *Handler) SetUpvoteMessageCache(c *gin.Context) {
	req := request.SetUpvoteMessageCacheRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		h.log.
			Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID, "msgID": req.MessageID, "userID": req.UserID}).
			Error(err, "[handler.SetUpvoteCache] - failed to read JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ChannelID == "" || req.GuildID == "" || req.MessageID == "" || req.UserID == "" {
		h.log.
			Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID, "msgID": req.MessageID, "userID": req.UserID}).
			Info("[handler.SetUpvoteCache] - missing request fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing request fields"})
		return
	}

	err = h.entities.SetUpvoteMessageCache(&req)
	if err != nil {
		h.log.
			Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID, "msgID": req.MessageID, "userID": req.UserID}).
			Error(err, "[handler.SetUpvoteCache] - failed to set cache")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.ResponseMessage{Message: "ok"})
}
