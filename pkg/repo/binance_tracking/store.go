package binancetracking

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	FirstOrCreate(binanceTracking *model.BinanceTracking) (*model.BinanceTracking, error)
	Update(binanceTracking *model.BinanceTracking) error
}
