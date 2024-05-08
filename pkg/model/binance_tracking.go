package model

import "time"

type BinanceTracking struct {
	ID           int64     `json:"id"`
	ProfileId    string    `json:"profile_id"`
	SpotLastTime time.Time `json:"spot_last_time"`
}
