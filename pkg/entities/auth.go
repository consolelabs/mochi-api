package entities

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (e *Entity) Login(accessToken string) (*LoginResponse, error) {

	s, err := discordgo.New("Bearer " + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to open discord session: %v", err.Error())
	}

	du, err := s.User("@me")
	if err != nil {
		return nil, fmt.Errorf("failed to get current discord user: %v", err.Error())
	}

	u, err := e.repo.Users.GetOne(du.ID)
	switch err {
	case nil:
		if u.Username != du.Username {
			u.Username = du.Username
			err = e.repo.Users.Update(u)
			if err != nil {
				return nil, fmt.Errorf("failed to update user: %v", err.Error())
			}
		}
	case gorm.ErrRecordNotFound:
		if err := e.generateInDiscordWallet(&model.User{
			ID:       du.ID,
			Username: du.Username,
		}); err != nil {
			return nil, fmt.Errorf("failed to generate in-discord wallet: %v", err.Error())
		}
	default:
		return nil, fmt.Errorf("failed to get user: %v", err.Error())
	}

	expirationTime := time.Now().Add(e.cfg.JWTAccessTokenLifeSpan)

	issuer := jwt.NewWithClaims(jwt.SigningMethodHS512, model.JWTData{
		DiscordAccessToken: accessToken,
		UserDiscordID:      du.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	})

	token, err := issuer.SignedString(e.cfg.JWTSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to issue new token for discord_id %s with err %s", du.ID, err.Error())
	}

	return &LoginResponse{
		AccessToken: token,
		ExpiresAt:   expirationTime.Unix(),
	}, nil
}

func (e *Entity) GetMyDiscordInfo(accessToken string) (*discordgo.User, error) {
	s, err := discordgo.New("Bearer " + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to open discord session: %v", err.Error())
	}

	du, err := s.User("@me")
	if err != nil {
		return nil, fmt.Errorf("failed to get current discord user: %v", err.Error())
	}

	return du, nil
}
