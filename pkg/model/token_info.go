package model

import "gorm.io/datatypes"

type TokenInfo struct {
	Source string         `json:"source"`
	Token  string         `json:"token"`
	Data   datatypes.JSON `json:"data"`
}
