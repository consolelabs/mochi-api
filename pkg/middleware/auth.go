package middleware

import (
	"strings"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/gin-gonic/gin"
)

func getTokenStringFromContext(c *gin.Context) string {
	authorization := c.GetHeader("Authorization")
	tokenStr := strings.Split(authorization, " ")

	if authorization != "" && len(tokenStr) == 2 {
		return tokenStr[1]
	}

	cookieStr, err := c.Cookie(consts.TokenCookieKey)
	if err != nil || cookieStr == "" {
		return ""
	}

	return cookieStr

}
