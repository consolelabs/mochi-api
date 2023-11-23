package sentrygo

import "time"

type Service interface {
	Flush(timeout time.Duration) bool
	CaptureErrorEvent(p SentryCapturePayload)
}
