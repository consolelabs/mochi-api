package util

import "strings"

func isMissingPermissionsErr(msg string) bool {
	return strings.Contains(msg, "missing permissions") && strings.Contains(msg, "50013")
}

func isMissingAccessErr(msg string) bool {
	return strings.Contains(msg, "missing access") && strings.Contains(msg, "50001")
}

func isMemberNotFoundErr(msg string) bool {
	return strings.Contains(msg, "404 not found") && strings.Contains(msg, "10007")
}

func IsAcceptableErr(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return isMissingPermissionsErr(msg) || isMissingAccessErr(msg) || isMemberNotFoundErr(msg)
}
