package model

import "github.com/golang-jwt/jwt/v4"

type JWTData struct {
	DiscordAccessToken string `json:"discord_access_token,omitempty"`
	UserDiscordID      string `json:"user_discord_id,omitempty"`
	jwt.RegisteredClaims
}
