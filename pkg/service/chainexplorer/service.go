package chainexplorer

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetGasTracker(listChain []model.Chain) ([]response.GasTrackerResponse, error)
}
