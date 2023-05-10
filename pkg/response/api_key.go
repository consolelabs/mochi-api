package response

import (
	"time"
)

type ProfileApiKeyResponse struct {
	ProfileId string    `json:"profile_id"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
