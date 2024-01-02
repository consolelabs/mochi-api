package butils

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"
)

func signature(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	signingKey := fmt.Sprintf("%x", mac.Sum(nil))
	return signingKey

}

func QueryString(q map[string]string, apiSecret string) string {
	var s []string
	for k, v := range q {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	queryString := strings.Join(s, "&")
	signature := signature(queryString, apiSecret)
	return queryString + "&signature=" + signature
}

func FirstTimestamp() int64 {
	return 1498880255000
}
