package badapter

type ErrorCode int64

const (
	REJECTED_MBX_KEY ErrorCode = -2015
)

type BinanceFutureErrorResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
