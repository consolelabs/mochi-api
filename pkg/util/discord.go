package util

import (
	"strings"
	"time"
)

func isMissingPermissionsErr(msg string) bool {
	return strings.Contains(msg, "missing permissions") && strings.Contains(msg, "50013")
}

func isMissingAccessErr(msg string) bool {
	return strings.Contains(msg, "missing access") && strings.Contains(msg, "50001")
}

func isMemberNotFoundErr(msg string) bool {
	return strings.Contains(msg, "404 not found") && strings.Contains(msg, "10007")
}

func IsRoleNotFoundErr(msg string) bool {
	msg = strings.ToLower(msg)
	return strings.Contains(msg, "404 not found") && strings.Contains(msg, "10011")
}

func IsGuildNotFound(msg string) bool {
	return strings.Contains(msg, "404 not found") && strings.Contains(msg, "10004") // bot not in guild

}

func IsAcceptableErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return isMissingPermissionsErr(msg) ||
		isMissingAccessErr(msg) ||
		isMemberNotFoundErr(msg) ||
		IsGuildNotFound(msg)
}

// RetryRequest retry handler until it succeeds or is acceptable or reaches the limit of times
func RetryRequest(handler func() error, times int, interval time.Duration) error {
	err := handler()
	for i := 0; err != nil && !IsAcceptableErr(err) && i < times-1; i++ {
		time.Sleep(interval)
		err = handler()
	}
	return err
}
