package sentrygo

import "time"

type Service interface {
	Flush(timeout time.Duration) bool
	CaptureErrorEvent(msg string, data map[string]interface{})
}
