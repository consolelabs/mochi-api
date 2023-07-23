package model

import "gorm.io/datatypes"

type TokenInfo struct {
	Token string         `json:"token"`
	Data  datatypes.JSON `json:"data"`
}
