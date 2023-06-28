package model

import (
	"gorm.io/datatypes"
)

type CoingeckoInfo struct {
	ID   string         `json:"id"`
	Info datatypes.JSON `json:"info"`
}
