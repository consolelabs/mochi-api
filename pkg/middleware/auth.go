package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/model"
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

func AuthGuard(cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := getTokenStringFromContext(ctx)

		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "please login to continue",
			})
			ctx.Abort()
			return
		}

		// always allows mochi bot
		if tokenStr == cfg.MochiBotSecret {
			ctx.Next()
			return
		}

		var claims jwt.MapClaims
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return cfg.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token or token expired",
				"error":   err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("profile_id", claims["profile_id"])
		ctx.Next()
	}
}

func WithAuthContext(cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := getTokenStringFromContext(ctx)
		ctx.Set("isMochiBot", tokenStr == cfg.MochiBotSecret)

		// always allows mochi bot
		if tokenStr == cfg.MochiBotSecret || tokenStr == "" {
			ctx.Next()
			return
		}

		var claims model.JWTData
		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return cfg.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.Next()
			return
		}

		ctx.Set("discord_access_token", claims.DiscordAccessToken)
		ctx.Set("user_discord_id", claims.UserDiscordID)

		ctx.Next()
	}
}

func ProfileAuthGuard(cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		profileAccessToken := ctx.GetHeader("Authorization")

		if profileAccessToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "please login to continue",
			})
			ctx.Abort()
			return
		}

		ctx.Set("profile_access_token", profileAccessToken)
		ctx.Next()
	}
}
