package chainexplorer

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetGasTracker() ([]response.GasTrackerResponse, error)
}
