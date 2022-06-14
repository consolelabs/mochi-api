package userdefiactivitylog

import (
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	CreateTransferLogs(req request.TransferRequest, tokenID int, amountEach, totalAmount float64) error
}
