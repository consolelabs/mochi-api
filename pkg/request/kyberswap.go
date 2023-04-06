package request

import "github.com/defipod/mochi/pkg/model"

type BuildRouteRequest struct {
	Recipient    string             `json:"recipient" binding:"required"`
	Sender       string             `json:"sender" binding:"required"`
	RouteSummary model.RouteSummary `json:"route_mummary" binding:"required"`
}
